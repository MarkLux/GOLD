package k8s

import "github.com/MarkLux/GOLD/api/restful/orm"

type KubernetesService interface {
	// init
	CreateService(service orm.FunctionService) error
	//
}
