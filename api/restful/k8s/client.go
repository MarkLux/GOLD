package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/MarkLux/GOLD/api/restful/constant"
	"log"
	"sync"
)

var once sync.Once
var clientInstance *kubernetes.Clientset

func buildClient() {
	config, err := clientcmd.BuildConfigFromFlags("", &constant.KubeConfigPath)
	if err != nil {
		log.Fatal("fail to init k8s client", err)
	}
	clientInstance, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("fail to init k8s client", err)
	}
}

func GetClient() *kubernetes.Clientset {
	if clientInstance != nil {
		return clientInstance
	} else {
		once.Do(buildClient)
		return clientInstance
	}
}
