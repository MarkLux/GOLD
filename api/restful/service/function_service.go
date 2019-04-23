package service

import (
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/go-xorm/xorm"
	"sync"
)

// instance for function service
type FunctionService struct {
	engine       *xorm.Engine
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
	// validate?
	// then create
	f.InitTime()
	_, err = s.engine.Insert(f)
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