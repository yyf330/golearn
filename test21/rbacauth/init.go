package rbacauth

import (
	"path/filepath"
	"flag"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)
var clientset kubernetes.Interface
func Init() (kubernetes.Interface, error) {
	var kubeconfig *string
	var err error
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//panic(err)
		return nil, err
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		//panic(err)
		return nil, err
	}

	return clientset,err
}
