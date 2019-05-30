package main

import (
	"os"
	"github.com/shima-park/agollo"
	"github.com/astaxie/beego/logs"
	"fmt"
	"io/ioutil"
)

type Client struct {
	Url       string
	AppId     string
	Cluster   string
	DefaultNs string
	ExtendNs  []string
}

const (
	YamlFile = "yaml"
	XmlFile  = "xml"
	YmlFile  = "yml"
	JsonFile = "json"
)

func (c *Client) ClientInit(apolloClient *Client) (error, <-chan *agollo.ApolloResponse, <-chan *agollo.LongPollerError) {
	file, err := newBackFile("")
	if err != nil {
		logs.Error("Back File Create Failed!")
		return err, nil, nil
	}
	fmt.Println("##1#")

	if err := agollo.Init(apolloClient.Url, apolloClient.AppId,
		agollo.Cluster(apolloClient.Cluster),
		agollo.PreloadNamespaces(apolloClient.DefaultNs),
		agollo.AutoFetchOnCacheMiss(),       // 在配置未找到时，去apollo的带缓存的获取配置接口，获取配置
		agollo.FailTolerantOnBackupExists(), // 在连接apollo失败时，如果在配置的目录下存在.agollo备份配置，会读取备份在服务器无法连接的情况下
		agollo.BackupFile(file),
	); err != nil {
		fmt.Println("#####error", err)
		return err, nil, nil
	}
	fmt.Println("###")
	for _, ns := range apolloClient.ExtendNs {
		agollo.PreloadNamespaces(ns)
	}
	// 获取默认配置中cluster=default namespace=application key=Name的值，提供默认值返回
	//fmt.Println("YourConfigKey:", agollo.Get("YourConfigKey", agollo.WithDefault("YourDefaultValue")))

	// 获取namespace下的所有配置项
	//fmt.Println("Configuration of the demo.namespace-appc:", agollo.GetNameSpace("demo.namespace-appc"))
	//fmt.Println(agollo.Get("test", agollo.WithDefault("10.0.0.2"), agollo.WithNamespace("application")))

	// 如果想监听并同步服务器配置变化，启动apollo长轮训
	// 返回一个期间发生错误的error channel,按照需要去处理
	errorCh := agollo.Start()

	// 监听apollo配置更改事件
	// 返回namespace和其变化前后的配置,以及可能出现的error
	watchCh := agollo.Watch()
	return nil, watchCh, errorCh
}

func (c *Client) GetPublicNamespaceConfigurations(ns string) agollo.Configurations {
	return agollo.GetNameSpace(ns)
}
func (c *Client) GetPrivateNamespaceConfigurationsContent(ns, fileType string) interface{} {
	ns = ns + "." + fileType
	return agollo.GetNameSpace(ns)["content"]
}

func (c *Client) GetPrivateNamespaceConfigurations(ns, fileType string) agollo.Configurations {
	ns = ns + "." + fileType
	return agollo.GetNameSpace(ns)
}

func (c *Client) CleanClient() {
	agollo.Stop()
	os.RemoveAll(agollo.GetAgollo().Options().BackupFile)
}

func main() {
	client := &Client{Url: "10.200.204.46:8080", AppId: "appc-demo-dl", Cluster: "default", DefaultNs: "demo.namespace-appc", ExtendNs: []string{"json.test"}}
	client.ClientInit(client)
	fmt.Println(client.GetPublicNamespaceConfigurations("demo.namespace-appc"))
	fmt.Println(client.GetPrivateNamespaceConfigurations("json.test", JsonFile))
	fmt.Println(client.GetPrivateNamespaceConfigurationsContent("json.test", JsonFile))
	client.CleanClient()
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

//import (
//	"github.com/philchia/agollo"
//	"encoding/json"
//	"fmt"
//)
//
//func main() {
//	conf := &agollo.Conf{IP: "10.200.204.46:8080", AppID: "appc-demo-dl", Cluster: "default", NameSpaceNames: []string{"demo.namespace-appc", "json.test"}}
//	client := agollo.NewClient(conf)
//	client.Start()
//	event := client.WatchUpdate()
//	//event := agollo.WatchUpdate()
//	fmt.Println("##",client.GetNameSpaceContent("json.test", "12345"))
//	fmt.Println("##",client.GetAllKeys("demo.namespace-appc"))
//	//client.
//	changeEvent := <-event
//	bytes, _ := json.Marshal(changeEvent)
//	fmt.Println("event:", string(bytes))
//	go func() {
//		for {
//			select {
//			//case err := <-errorCh:
//			//	fmt.Println("Error:", err)
//			case changeEvent := <-event:
//				bytes, _ := json.Marshal(changeEvent)
//				fmt.Println("event:", string(bytes))            }
//		}
//	}()
//	select {}
//}
