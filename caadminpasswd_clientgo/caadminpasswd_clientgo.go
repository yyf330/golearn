package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/api/core/v1"
)

func main() {
	config := &rest.Config{}
	//cc:=clientcmdapi.Config{}
	authInfo := &api.AuthInfo{}
	authInfo.Username = "admin"
	authInfo.Password = "UScGPRbEBC3GQ3vJ7afLTqly7dQAuyZI"
	config.CAFile = "/root/temp/hotfix.ca" //[]byte("-----BEGIN CERTIFICATE-----\nMIIDNjCCAh6gAwIBAgIIDGEu3O35yn4wDQYJKoZIhvcNAQELBQAwOTELMAkGA1UE\nBhMCQ04xEzARBgNVBAoTCnRlbmNlbnR5dW4xFTATBgNVBAMTDGNscy1qM2l2ams0\nZDAeFw0xODA4MjQwODAzNTFaFw0zODA4MjQwODAzNTFaMDkxCzAJBgNVBAYTAkNO\nMRMwEQYDVQQKEwp0ZW5jZW50eXVuMRUwEwYDVQQDEwxjbHMtajNpdmprNGQwggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDZ2Jg7xjmZlsoRTJCdvMuSQn7J\ny7oe78pDbDQQI6DCModjH/nLY94z8m2V4Z2dsj3YgkPdcoA2ToAOh60nE95W6ZAB\nyA/O9/z9CMsgPkmb0iVGJFDz21ON8gOn0R3msVs2nuCa56ozLyrNOOryzDG7NM92\nAM6UaOVS9eQnpoe/FN3BLVT3JIcUAYDekW7/2xPp3M3htNAk0/s719pQiXAjzslL\np8Ob5hsmEMU6lefIOsTmQESYfXtutEaBKIWMWMllT0AHkyWG+BCsIkn+po0qvEz0\nwPJRynY0LQd/Fsw8+vVa0lPJUsePMZd28aeu/VtffvnJJ2ls9h5PFOSMyJkJAgMB\nAAGjQjBAMA4GA1UdDwEB/wQEAwIChDAdBgNVHSUEFjAUBggrBgEFBQcDAgYIKwYB\nBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAyPBuxSuE\nHuiBs5r+uNpl/bU4P7AxdTXEwFbE6x4HfqgRKgeEqPxKD06huhk9MV3a1c0iVdhp\nVsNIa6h+JgXETGnLEEUgnRvqL7jRbzs7JoJdvjY7HIsYgeAHoq6xF5Lnh/QhXQiu\nB0asD6QvooJReFUd4OGBUKwZcPyitwr245ad0tDoaMQW/Z+ljP7ax1JbdMsUw9jU\nC9e02qXCa9cd52n9oGCsLt6hQD9BR3PkFnFosmDrLvhUoQboXI5D8PEUoxkYfPR3\nmyQ6n65s4W1+DOp8hy+sLea9IoqDuLuhooQOp8Wp/06EvWv9s3sgw6BnuytGffW2\nzSrx7D7uc/IPpQ==\n-----END CERTIFICATE-----\n")
	//config.Host = "https://10.214.214.3:10218"

	cmdCfg := api.NewConfig()
	cmdCfg.Clusters["kubernetes"] = &api.Cluster{
		Server:                   config.Host,
		CertificateAuthority:     config.TLSClientConfig.CAFile,
		CertificateAuthorityData: config.TLSClientConfig.CAData,
		InsecureSkipTLSVerify:    config.TLSClientConfig.Insecure,
	}
	cmdCfg.AuthInfos["kubernetes"] = authInfo
	cmdCfg.Contexts["kubernetes"] = &api.Context{
		Cluster:  "kubernetes",
		AuthInfo: "kubernetes",
	}
	cmdCfg.CurrentContext = "kubernetes"

	clientConfig := clientcmd.NewDefaultClientConfig(
		*cmdCfg,
		&clientcmd.ConfigOverrides{},
	)

	cfg, err := clientConfig.ClientConfig()

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("###")
	ver, e := clientset.ServerVersion()
	if e != nil {
		fmt.Println("e=", e)

	}
	fmt.Println("#version#", ver)
	vv, ee := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if ee != nil {
		fmt.Println("ee=", ee)

	}
	tt := rest.CopyConfig(cfg)
	fmt.Printf("%+v", tt)
	//vv,_:=clientset.ServerVersion()
	for _, i := range vv.Items {
		fmt.Println("###", i.Name)

	}
	fmt.Println("###end")
	ta, err := clientset.CoreV1().ConfigMaps("default").Get("yonghui.cn", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ta)
	aa := make(map[string]string, 0)
	aa["IsUseing"] = "Yes"
	cc := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "yonghui.cn",
		},
		Data: aa,
	}
	clientset.CoreV1().ConfigMaps("default").Create(cc)
	//cm, _ := clientset.CoreV1().ConfigMaps("default").Get("yonghui.cn", metav1.GetOptions{})
	//fmt.Println("#####:", cm.Data)
	//if _, ok := cm.Data["IsUseing"]; ok {
	//	fmt.Println("----", cm.Data["IsUseing"])
	//}
	clientset.CoreV1().ConfigMaps("default").Delete("yonghui.cn", &metav1.DeleteOptions{})
}
