package common

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sync"
)

var csInstance *kubernetes.Clientset = nil
var k8sOnce sync.Once

func InitK8sClientSet() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return KubernetesErr{Action: "InitConfig", Message: err.Error()}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return KubernetesErr{Action: "InitClientSet", Message: err.Error()}
	}

	if clientSet != nil {
		csInstance = clientSet
	} else {
		return KubernetesErr{Action: "InitClientSet", Message: "get nil client set"}
	}
	return nil
}

func GetK8sClientSet() (cs *kubernetes.Clientset, err error) {
	// if there is some wrong with init, try again, MAY NOT USE
	if csInstance != nil {
		 cs = csInstance
	} else {
		k8sOnce.Do(func() {
			err = InitK8sClientSet()
			cs = csInstance
		})
	}
	return
}
