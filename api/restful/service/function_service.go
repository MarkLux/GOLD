package service

import (
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/docker/docker/client"
	"github.com/go-xorm/xorm"
	"k8s.io/client-go/kubernetes"
	"sync"
)

// instance for function service
type FunctionService struct {
	engine       *xorm.Engine
	dockerCli    *client.Client
	k8sCli		 *kubernetes.Clientset
}

var functionInstance *FunctionService
var functionOnce sync.Once

type UpdateAction struct {
	FunctionService orm.FunctionService
	TargetBranch string
	TargetVersion string
}

func (s FunctionService) CreateFunctionService(f *orm.FunctionService) (err error)  {
	// check if the function existed
	nameCnt, err := s.engine.Count(&orm.FunctionService{ServiceName: f.ServiceName})
	if err != nil {
		return
	}
	if nameCnt > 0 {
		err = errors.GenFunctionServiceExistedError()
		return
	}
	// validate?
	// then create
	f.InitTime()
	_, err = s.engine.Insert(f)
	return
}

func (s FunctionService) PublishFunctionService(action UpdateAction) (opId int64, err error) {
	// 1. build
}

func (s FunctionService) ListFunctionService(page int, size int) (total int64, results []orm.FunctionService, err error) {
	total, err = s.engine.Table("function_services").Count()
	if err != nil {
		return
	}
	err = s.engine.Limit(size, (page-1) * size).Find(&results)
	return
}

func GetFunctionService() *FunctionService {
	functionOnce.Do(func() {
		functionInstance = &FunctionService {
			engine: orm.GetOrmEngine(),
		}
	})
	return functionInstance
}

// atomic actions
func buildImage(f orm.FunctionService) error {
	// check parameters
}