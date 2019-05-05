package service

import (
	"github.com/MarkLux/GOLD/api/restful/constant"
	"github.com/MarkLux/GOLD/api/restful/docker"
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/k8s"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/docker/docker/client"
	"github.com/go-xorm/xorm"
	"k8s.io/client-go/kubernetes"
	"sync"
	"time"
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
	Operator orm.User
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
	// then create
	f.InitTime()
	_, err = s.engine.Insert(f)
	return
}

func (s FunctionService) PublishFunctionService(action UpdateAction) int64 {
	operateLog := &orm.OperateLogs{
		ServiceId: action.FunctionService.Id,
		OperatorId: action.Operator.Id,
		Type: constant.OperatePublish,
		Start: time.Now().Unix(),
		OriginBranch: action.FunctionService.GitBranch,
		OriginVersion: action.FunctionService.GitHead,
		TargetBranch: action.TargetBranch,
		TargetVersion: action.TargetVersion,
		CurrentAction: "START",
	}
	s.engine.Insert(operateLog)
	s.engine.Table(orm.FunctionService{}).
		ID(action.FunctionService.Id).Cols("last_operation").
		Update(map[string]interface{}{"last_operation": operateLog.Id})
	// update function service
	f := action.FunctionService
	f.GitHead = action.TargetVersion
	f.GitBranch = action.TargetBranch
	s.buildImage(f, *operateLog)
	return operateLog.Id
}

func (s FunctionService) ListFunctionService(page int, size int) (total int64, results []orm.FunctionService, err error) {
	total, err = s.engine.Table("function_services").Count()
	if err != nil {
		return
	}
	err = s.engine.Limit(size, (page-1) * size).Find(&results)
	return
}

func (s FunctionService) GetFunctionService(id int64) orm.FunctionService {
	f := &orm.FunctionService{
		Id: id,
	}
	s.engine.Get(f)
	return *f
}

func GetFunctionService() *FunctionService {
	functionOnce.Do(func() {
		functionInstance = &FunctionService {
			engine: orm.GetOrmEngine(),
			dockerCli: docker.GetClient(),
			k8sCli: k8s.GetClient(),
		}
	})
	return functionInstance
}