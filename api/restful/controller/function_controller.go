package controller

import (
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/MarkLux/GOLD/api/restful/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreateFunctionServiceRequest struct {
	ServiceName   string `json:"serviceName" binding:"required"`
	GitRepo       string `json:"gitRepo" binding:"required"`
	GitBranch     string `json:"gitBranch" binding:"required"`
	GitMaintainer string `json:"gitMaintainer" binding:"required"`
	GitHead       string `json:"gitHead"`
	MinInstance   int    `json:"minInstance" binding:"required"`
	MaxInstance   int    `json:"maxInstance" binding:"required"`
}

type FunctionServiceController struct {
	functionService *service.FunctionService
	tokenService    *service.TokenService
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
		CreatorId:     user.Id,
		CreatorName:   user.Name,
		ServiceName:   req.ServiceName,
		GitRepo:       req.GitRepo,
		GitBranch:     req.GitBranch,
		GitMaintainer: req.GitMaintainer,
		GitHead:       req.GitHead,
		MinInstance:   req.MinInstance,
		MaxInstance:   req.MaxInstance,
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

func (c FunctionServiceController) ListFunctionService(ctx *gin.Context) {
	page := 1
	size := 10
	var err error
	pageStr := ctx.Query("page")
	if pageStr == "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
	}
	sizeStr := ctx.Query("size")
	if sizeStr == "" {
		size, err = strconv.Atoi(sizeStr)
		if err != nil {
			size = 10
		}
	}
	total, results, err := c.functionService.ListFunctionService(page, size)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.GenUnknownError())
		return
	}
	data := make(map[string]interface{})
	data["results"] = results
	data["total"] = total
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func NewFunctionServiceController() FunctionServiceController {
	return FunctionServiceController{
		functionService: service.GetFunctionService(),
		tokenService:    service.GetTokenService(),
	}
}
