package controller

import (
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/MarkLux/GOLD/api/restful/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateFunctionServiceRequest struct {
	ServiceName string `json:"serviceName" binding:"required"`
	GitRemote   string `json:"gitRemote" binding:"required"`
	GitBranch   string `json:"gitBranch" binding:"required"`
	GitHead     string `json:"gitHead"`
	MinInstance int    `json:"minInstance" binding:"required"`
	MaxInstance int    `json:"maxInstance" binding:"required"`
}

type FunctionServiceController struct {
	functionService *service.FunctionService
	tokenService *service.TokenService
}

func (c FunctionServiceController) CreateFunctionService(ctx *gin.Context) {
	var req CreateFunctionServiceRequest
	var err error
	err = ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.GenValidationError())
		return
	}
	// check login user
	hasLogin, user := service.NeedLoginCheck(ctx)
	if !hasLogin {
		return
	}
	// create do
	f := orm.FunctionService{
		CreatorId: user.Id,
		CreatorName: user.Name,
		ServiceName: req.ServiceName,
		GitRemote: req.GitRemote,
		GitBranch: req.GitBranch,
		GitHead: req.GitHead,
		MinInstance: req.MinInstance,
		MaxInstance: req.MaxInstance,
	}
	err = c.functionService.CreateFunctionService(&f)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	data := make(map[string]interface{})
	data["serviceId"] = f.Id
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func NewFunctionServiceController() FunctionServiceController {
	return FunctionServiceController{
		functionService: service.GetFunctionService(),
		tokenService: service.GetTokenService(),
	}
}