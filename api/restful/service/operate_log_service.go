package service

import (
	"bufio"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"io"
	"log"
	"sync"
	"time"
)

const (
	ActionStart       = "START"
	ActionImgBuilding = "IMG_BUILDING"
	ActionImgPushing  = "IMG_PUSHING"
	ActionIniting 	  = "INITING"
	ActionPublishing  = "PUBLISHING"
	ActionRollBacking = "ROLLBACKING"
	ActionFinish      = "FINISH"
	ActionFailed      = "FAILED"
)

type OperateLogService struct {
	engine  *xorm.Engine
	rClient *redis.Client
}

var operateInstance *OperateLogService
var operateOnce sync.Once


func (s OperateLogService) CreateOperateLogService(action Action) (opLog *orm.OperateLogs, err error) {
	current := time.Now().Unix()
	// insert operate log
	operateLog := &orm.OperateLogs {
		ServiceId:     action.FunctionService.Id,
		OperatorId:    action.Operator.Id,
		Type:          action.Type,
		Start:         current,
		OriginBranch:  action.FunctionService.GitBranch,
		OriginVersion: action.FunctionService.GitHead,
		TargetBranch:  action.TargetBranch,
		TargetVersion: action.TargetVersion,
		CurrentAction: ActionStart,
	}
	operateLog.CreatedAt = current
	operateLog.Update = current
	_,err = s.engine.Insert(operateLog)
	if err != nil {
		log.Println("fail to create operate log, ", err)
		return
	}

	// update last operation
	_, err = s.engine.Table(orm.FunctionService{}).
		ID(action.FunctionService.Id).Cols("last_operation").
		Update(map[string]interface{}{"last_operation": operateLog.Id})
	if err != nil {
		log.Println("fail to attach operation to fs, ", err)
		return
	}
	opLog = operateLog
	return
}

func (s OperateLogService) ContinueOperateLog(opLog *orm.OperateLogs, currentAction string, output io.Reader, hasOutput bool) (lastOutput string, err error) {
	// update operate log
	current := time.Now().Unix()
	updateLog := &orm.OperateLogs{
		CurrentAction: currentAction,
		Update: current,
	}
	updateLog.UpdatedAt = current
	_, err = s.engine.Table(orm.OperateLogs{}).ID(opLog.Id).Update(updateLog)
	if err != nil {
		log.Println("fail to update op log, ", err)
	}
	// start recording output
	if !hasOutput {
		return
	}
	var fullOutput string
	bufReader := bufio.NewReader(output)
	for {
		line, _, err := bufReader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("fail to read from output, ", err)
			break
		}
		lastOutput = string(line)
		log.Println("[output]", lastOutput)
		fullOutput += lastOutput
		fullOutput += "\n"
	}
	// rewrite into db
	opLog.Log += fullOutput
	_, err = s.engine.Table(orm.OperateLogs{}).ID(opLog.Id).Update(&orm.OperateLogs{
		Log: opLog.Log,
		Update: time.Now().Unix(),
	})
	if err != nil {
		log.Println("fail to update log output, ", err)
	}
	return
}

func (s OperateLogService) FinishOperateLog(opLog *orm.OperateLogs) (err error) {
	current := time.Now().Unix()
	_, err = s.engine.Table(orm.OperateLogs{}).ID(opLog.Id).Update(&orm.OperateLogs{
		Update: current,
		End: current,
		CurrentAction: ActionFinish,
	})
	if err != nil {
		log.Println("fail to finish operate log.")
	}
	return
}

func (s OperateLogService) FailOperateLog(opLog *orm.OperateLogs, reason string) (err error) {
	current := time.Now().Unix()
	_, err = s.engine.Table(orm.OperateLogs{}).ID(opLog.Id).Update(&orm.OperateLogs{
		Update: current,
		End: current,
		CurrentAction: ActionFailed,
		Reason: reason,
	})
	if err != nil {
		log.Println("fail to fail operate log.")
	}
	return
}

func GetOperateService() *OperateLogService {
	operateOnce.Do(func() {
		operateInstance = &OperateLogService{
			engine: orm.GetOrmEngine(),
		}
	})
	return operateInstance
}