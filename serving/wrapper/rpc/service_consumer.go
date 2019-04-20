package rpc

import (
	"github.com/MarkLux/GOLD/serving/common"
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
	"github.com/MarkLux/GOLD/serving/wrapper/constant"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// interface
type ServiceConsumer interface {
	Request(request map[string]interface{}) (map[string]interface{}, error)
}

// implement
type GoldServiceConsumer struct {
	// the target restful name in gold
	TargetServiceName string
	// timeout setting for client
	ClientTimeOut int64
}

func GetRemoteService(serviceName string) (*GoldServiceConsumer) {
	return GetRemoteServiceWithTimeOut(serviceName, constant.DefaultClientTimeOut)
}

func GetRemoteServiceWithTimeOut(serviceName string, timeout int64) (*GoldServiceConsumer) {
	return &GoldServiceConsumer{
		TargetServiceName: serviceName,
		ClientTimeOut: timeout,
	}
}

func (consumer *GoldServiceConsumer) Request(req map[string]interface{}) (rsp map[string]interface{}, err error) {
	// 0. parseRequest
	request := &goldrpc.GoldRequest{
		Invoker: common.GetGoldEnv().PodName,
		TimeStamp: time.Now().Unix(),
		Data: req,
	}
	// 1. using k8s api to found the restful -- deprecated
	// using kube-dns instead.
	/*
	restful, err := parseService(consumer.TargetServiceName)
	if err != nil {
		return
	}
	clusterIP := restful.Spec.ClusterIP
	if clusterIP == "" {
		err = common.ServiceNotFoundErr{TargetService: consumer.TargetServiceName, Detail: "got blank cluster ip."}
		return
	}
	*/
	// 2. make request through restful cluster ip
	rpcClient := &goldrpc.GoldRpcClient{
		TargetIP:   consumer.TargetServiceName,
		TargetPort: constant.DefaultServicePort,
		TimeOut:    consumer.ClientTimeOut,
	}
	// 3. sync get response and return
	response, err := rpcClient.RequestSync(request)
	if err != nil {
		return
	}
	rsp = response.Data
	return
}

// get restful info from k8s
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
