package common

import "fmt"

// common errors
type ClientTimeOutErr struct {
	TargetService string
	ClientTimeOut int32
}

func (e ClientTimeOutErr) Error() string {
	return fmt.Sprintf("request restful %s timeout after waiting %d",
		e.TargetService, e.ClientTimeOut)
}

type ServiceNotFoundErr struct {
	TargetService string
	Detail string
}

func (e ServiceNotFoundErr) Error() string {
	return fmt.Sprintf("no restful provider of %s found, detail info: %s", e.TargetService, e.Detail)
}

type KubernetesErr struct {
	Action string
	Message string
}

func (e KubernetesErr) Error() string {
	return fmt.Sprintf("k8s err, action: %s, message: %s ", e.Action, e.Message)
}

type UnknownErr struct {
	Message string
}

func (e UnknownErr) Error() string {
	return e.Message
}