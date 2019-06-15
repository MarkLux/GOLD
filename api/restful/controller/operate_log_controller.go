package controller

import (
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OperateLogController struct {
	opService *service.OperateLogService
	tokenService    *service.TokenService
}

func (c OperateLogController) GetLogDetail(ctx *gin.Context) {
	opId,err := strconv.ParseInt(ctx.Query("opId"),10,64)
	if err!=nil {
		ctx.JSON(http.StatusOK,errors.GenValidationError())
		return
	}
	opLog := c.opService.GetOperateLogService(opId)
	if opLog == nil {
		ctx.JSON(http.StatusOK,errors.GenOperateLogNotFoundError())
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"code": 0,
		"data": opLog,
	})
}

func NewOperateLogController() OperateLogController {
	return OperateLogController {
		opService: service.GetOperateService(),
		tokenService: service.GetTokenService(),
	}
}
