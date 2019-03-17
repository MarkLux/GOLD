package rpc

import (
	"github.com/MarkLux/GOLD/serving/common"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/constant"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// interface
type ServiceConsumer interface {
	Request(
		request goldrpc.GoldRequest,
		clientTimeOut int64) (goldrpc.GoldResponse, error)
}

// implement
type RemoteServiceConsumer struct {
	// the target service name in gold
	TargetServiceName string
	// timeout setting for client
	ClientTimeOut int64
}

func GetRemoteService(serviceName string) (*RemoteServiceConsumer) {
	return GetRemoteServiceWithTimeOut(serviceName, constant.DefaultClientTimeOut)
}

func GetRemoteServiceWithTimeOut(serviceName string, timeout int64) (*RemoteServiceConsumer) {
	return &RemoteServiceConsumer{
		TargetServiceName: serviceName,
		ClientTimeOut: timeout,
	}
}

func (consumer *RemoteServiceConsumer) Request(request goldrpc.GoldRequest) (response *goldrpc.GoldResponse, err error) {
	// 1. using k8s api to found the service
	service, err := parseService(consumer.TargetServiceName)
	if err != nil {
		return
	}
	clusterIP := service.Spec.ClusterIP
	if clusterIP == "" {
		err = common.ServiceNotFoundErr{TargetService: consumer.TargetServiceName, Detail: "got blank cluster ip."}
		return
	}
	// 2. make request through service cluster ip
	rpcClient := &goldrpc.GoldRpcClient{
		TargetIP:   clusterIP,
		TargetPort: constant.DefaultServicePort,
		TimeOut:    consumer.ClientTimeOut,
	}
	// 3. sync get response and return
	response, err = rpcClient.RequestSync(&request)
	if err != nil {
		return
	}
	return
}

// get service info from k8s
func parseService(serviceName string) (service *v1.Service, err error) {
	clientSet, err := common.GetK8sClientSet()
	if err != nil {
		return
	}
	service, err = clientSet.CoreV1().
		Services(constant.GoldNamespace).
		Get(serviceName, meta_v1.GetOptions{})
	if err != nil {
		return nil, common.ServiceNotFoundErr{TargetService: serviceName, Detail: err.Error()}
	}
	return service, nil
}
