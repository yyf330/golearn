package main

import (
	"os"
	"github.com/shima-park/agollo"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"fmt"
	"time"
)

type ContextClient struct {
	ApolloClient *Client
	exitCh       chan bool
	errorCh      <-chan *agollo.LongPollerError
	watchCh      <-chan *agollo.ApolloResponse
	ClientCache  *ClientCache
}

type ContextClientCache struct {
	data map[Client]*ContextClient
}

func GetContextClientCacheInitHandle() *ContextClientCache {
	if contextClient == nil {
		contextClient = &ContextClientCache{
			data: make(map[Client]*ContextClient),
		}
	}

	return contextClient
}

func (cc *ContextClientCache) Set(key Client, data *ContextClient) {
	cc.data[key] = data
}

func (cc *ContextClientCache) Get(key Client) *ContextClient {

	if v, ok := cc.data[key]; ok {
		return v
	} else {
		logs.Debug("client [%s] not in Cache.", key)
		return nil
	}
}

func (cc *ContextClientCache) Del(key Client) {
	delete(cc.data, key)
}

func (cc *ContextClientCache) Clean() {
	for key := range cc.List() {
		delete(cc.data, key)
	}
}

func (cc *ContextClientCache) List() map[Client]*ContextClient {
	return cc.data
}

func main() {

	contextClient := GetContextClientCacheInitHandle()
	apolloClient := Client{Url: "10.200.204.46:8080", AppId: "appc-demo-dl", Cluster: "default", DefaultNs: "demo.namespace-appc"}
	cc := &ContextClient{ApolloClient: &apolloClient}
	if err := cc.ApolloClient.ClientInit(); err != nil {
		return
	}
	cc.errorCh = cc.ApolloClient.GetErrorChannel()
	cc.watchCh = cc.ApolloClient.GetWatchChannel()
	cc.exitCh = make(chan bool)
	contextClient.Set(apolloClient, cc)

	go func(contextClient *ContextClient) {
		for {
			select {
			case err := <-contextClient.errorCh:
				fmt.Println("Error:", err)
			case update := <-contextClient.watchCh:
				fmt.Println("Apollo Update:", update)
			case <-contextClient.exitCh:
				fmt.Println("Apollo monitor exit:")
				return
			}
		}
	}(cc)

	//clientHandle := GetClientCacheInitHandle()
	//clientHandle.Set("dd", &Client{Url: "10.200.204.46:8080", AppId: "appc-demo-dl", Cluster: "default", DefaultNs: "demo.namespace-appc"})
	//clientHandle.Set("dl", &Client{Url: "10.200.204.52:8080", AppId: "appc-demo-dl", Cluster: "default", DefaultNs: "demo.namespace-appc"})
	//
	//ll := make(map[string]*ContextClient)
	//
	//ll["dd"] = &ContextClient{ApolloClient: clientHandle.Get("dd")}
	//time.Sleep(time.Second * 1)
	//ll["dl"] = &ContextClient{ApolloClient: clientHandle.Get("dl")}
	////todo get all apollo config form dolphin
	//for k, cl := range ll {
	//	cl.ApolloClient.ClientInit()
	//	cl.errorCh = cl.ApolloClient.GetErrorChannel()
	//	cl.watchCh = cl.ApolloClient.GetWatchChannel()
	//	logs.Error("####", k)
	//	logs.Error("####", cl.ApolloClient)
	//	go func(cc *ContextClient) {
	//		for {
	//			select {
	//			case err := <-cc.errorCh:
	//				fmt.Println("Error:", err)
	//			case update := <-cc.watchCh:
	//				fmt.Println("Apollo Update:", update)
	//			case <-cc.exitCh:
	//				fmt.Println("Apollo monitor exit:")
	//				return
	//			}
	//		}
	//	}(cl)
	//}
	for {
		time.Sleep(time.Second * 20)
		contextClient := GetContextClientCacheInitHandle()
		apolloClient := Client{Url: "10.200.204.46:8080", AppId: "appc-demo-dl", Cluster: "default", DefaultNs: "demo.namespace-appc"}
		cc := contextClient.Get(apolloClient)
		if cc == nil {
			fmt.Println("exit")
			return
		}
		cc.exitCh <- true
		apolloClient.CleanClient()
		contextClient.Del(apolloClient)
	}

	//
}

var (
	clientCache   *ClientCache
	contextClient *ContextClientCache
)

type ClientCache struct {
	data map[string]*Client
}

func GetClientCacheInitHandle() *ClientCache {
	if clientCache == nil {
		clientCache = &ClientCache{
			data: make(map[string]*Client),
		}
	}

	return clientCache
}

func (cc *ClientCache) Set(key string, data *Client) {
	cc.data[key] = data
}

func (cc *ClientCache) Get(key string) *Client {

	if v, ok := cc.data[key]; ok {
		return v
	} else {
		logs.Debug("client [%s] not in Cache.", key)
		return nil
	}
}

func (cc *ClientCache) Del(key string) {
	delete(cc.data, key)
}

func (cc *ClientCache) Clean() {
	for key := range cc.List() {
		delete(cc.data, key)
	}
}

func (cc *ClientCache) List() map[string]*Client {
	return cc.data
}

type Client struct {
	Url       string
	AppId     string
	Cluster   string
	DefaultNs string
	//ExtendNs  []string
}

const (
	YamlFile = "yaml"
	XmlFile  = "xml"
	YmlFile  = "yml"
	JsonFile = "json"
)

func (c *Client) ClientInit() error {
	file, err := newBackFile("")
	if err != nil {
		logs.Error("Back File Create Failed!")
		return err
	}
	if err := agollo.Init(c.Url, c.AppId,
		agollo.Cluster(c.Cluster),
		agollo.PreloadNamespaces(c.DefaultNs),
		agollo.AutoFetchOnCacheMiss(),       // 在配置未找到时，去apollo的带缓存的获取配置接口，获取配置
		agollo.FailTolerantOnBackupExists(), // 在连接apollo失败时，如果在配置的目录下存在.agollo备份配置，会读取备份在服务器无法连接的情况下
		agollo.BackupFile(file),
	); err != nil {
		return err
	}

	//for _, ns := range c.ExtendNs {
	//	agollo.PreloadNamespaces(ns)
	//}

	return nil
}

// 如果想监听并同步服务器配置变化，启动apollo长轮训
// 返回一个期间发生错误的error channel,按照需要去处理
func (c *Client) GetErrorChannel() <-chan *agollo.LongPollerError {
	return agollo.Start()
}

// 监听apollo配置更改事件
// 返回namespace和其变化前后的配置,以及可能出现的error
func (c *Client) GetWatchChannel() <-chan *agollo.ApolloResponse {
	return agollo.Watch()
}

func (c *Client) GetPublicNamespaceConfigurations(ns string) agollo.Configurations {
	return agollo.GetNameSpace(ns)
}

func (c *Client) GetPrivateNamespaceConfigurations(ns, fileType string) agollo.Configurations {
	ns = ns + "." + fileType
	return agollo.GetNameSpace(ns)
}

func (c *Client) GetPrivateNamespaceConfigurationsContent(ns, fileType string) interface{} {
	ns = ns + "." + fileType
	return agollo.GetNameSpace(ns)["content"]
}

func (c *Client) CleanClient() {
	agollo.Stop()
	os.RemoveAll(agollo.GetAgollo().Options().BackupFile)
}

var DirTemp = "temp"
var PREFIX = "prefix"

func init() {
	os.MkdirAll(DirTemp, 0777)
}

// todo when don't watch namespace os.RemoveAll(path)
func newBackFile(content string) (string, error) {
	file, err := ioutil.TempFile(DirTemp, PREFIX)

	if err != nil {
		logs.Error("create file failed!")
		return "", err
	}

	defer file.Close()

	path := file.Name()
	file.WriteString(content)

	return path, nil
}
