package docker

import (
	"github.com/docker/docker/client"
	"log"
	"sync"
)

var once sync.Once
var clientInstance *client.Client

func GetClient() *client.Client {
	if clientInstance != nil {
		return clientInstance
	} else {
		once.Do(buildClient)
		return clientInstance
	}
}

func buildClient() {
	var err error
	clientInstance, err = client.NewEnvClient()
	if err != nil {
		log.Fatalln("fail to init docker client", err)
	}
}