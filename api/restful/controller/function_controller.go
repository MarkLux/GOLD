package controller

import "github.com/MarkLux/GOLD/api/restful/service"

type CreateFunctionServiceRequest struct {
	ServiceName string `json:"serviceName" binding:"required"`
	GitRemote   string `json:"gitRemote" binding:"required"`
	GitBranch   string `json:"gitBranch" binding:"required"`
	GitHead     string `json:"gitHead" binding:"required"`
	MinInstance int    `json:"minInstance" binding:"required"`
	MaxInstance int    `json:"maxInstance" binding:"required"`
}

type FunctionServiceController struct {
	functionService service.FunctionService
}

