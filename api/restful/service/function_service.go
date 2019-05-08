package service

import (
	"github.com/MarkLux/GOLD/api/restful/docker"
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/k8s"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/docker/docker/client"
	"github.com/go-xorm/xorm"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
)

// instance for function service
type FunctionService struct {
	engine       *xorm.Engine
	dockerCli    *client.Client
	k8sCli		 *kubernetes.Clientset
	opService 	 *OperateLogService
}

var functionInstance *FunctionService
var functionOnce sync.Once

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

func (s FunctionService) PublishFunctionService(action Action) (opId int64, err error) {
	opLog, err := s.opService.CreateOperateLogService(action)
	if err != nil {
		log.Println("fail to create opLog, ", err)
		err = errors.GenUnknownError()
		return
	}
	// update function service
	f := action.FunctionService
	f.GitHead = action.TargetVersion
	f.GitBranch = action.TargetBranch
	opId = opLog.Id
	//err = s.buildImage(f, opLog)
	// check if the service existed ?
	err = s.initK8sService(f, opLog)
	if err != nil {
		log.Println("fail to init k8s service ", err)
		return
	}
	s.opService.FinishOperateLog(opLog)
	return
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
			opService: GetOperateService(),
		}
	})
	return functionInstance
}