package rpc

import (
	"github.com/MarkLux/GOLD/serving/wrapper/constant"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// interface
type ServiceConsumer interface {
	Request(serviceName string, clientTimeOut int32, request GoldRequest) (GoldResponse, error)
}

// implement
type RemoteServiceConsumer struct {
	// the target service name in gold
	TargetServiceName string
	// timeout setting for client
	ClientTimeOut int32
}

func (*RemoteServiceConsumer) Request(
	serviceName string, clientTimeOut int32, request GoldRequest) (*GoldResponse, error) {
	response := &GoldResponse{}
	// 1. using k8s api to found the service
	service, err := parseService(serviceName)
	if err != nil {
		return response, err
	}
	clusterIP := service.Spec.ClusterIP
	if clusterIP == "" {
		return response, ServiceNotFoundErr{TargetService: serviceName, Detail: "got blank cluster ip."}
	}
	// 2. make request through service cluster ip

	// 3. sync get response and return
	return response, nil
}

// get service info from k8s
func parseService(serviceName string) (*v1.Service, error) {

	service, err := clientSet.CoreV1().
		Services(constant.GoldNamespace).
		Get(serviceName, meta_v1.GetOptions{})
	if err != nil {
		return nil, ServiceNotFoundErr{TargetService: serviceName, Detail: err.Error()}
	}
	return service, nil
}