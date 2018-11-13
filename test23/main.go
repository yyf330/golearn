package main


import (
	"flag"
	"path/filepath"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/informers"
	"k8s.io/apimachinery/pkg/labels"
//"k8s.io/pkg/api"
	"fmt"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	/*test informer*/
	sharedinformerfactory := informers.NewSharedInformerFactory(clientset, time.Minute*10)

	stopch := make(chan struct{})
	sharedinformerfactory.Start(stopch)
	//sharedinformerfactory.
	podlist := sharedinformerfactory.Core().V1().Pods().Lister()
	fmt.Println(podlist.List(labels.Nothing()))
	fmt.Println(podlist.Pods("default").Get("mallbsweb"))
	fmt.Println(podlist.Pods("default").List(labels.Nothing()))
	/*end*/
}
