package trigger

import (
	"encoding/json"
	"fmt"
	"github.com/MarkLux/GOLD/serving/wrapper/rpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpInvokeRequest struct {
	ServiceName string `json:"serviceName"`
	Data string `json:"data"`
}

func HandleInvoke(c *gin.Context) {
	// parse request from json
	var request HttpInvokeRequest
	var err error
	if err = c.BindJSON(&request); err != nil {
		returnError(c, fmt.Sprintf("fail to bind request, %s", err.Error()))
		return
	}
	// parse json data
	reqData := make(map[string]interface{})
	err = json.Unmarshal([]byte(request.Data), &reqData)
	if err != nil {
		returnError(c, fmt.Sprintf("fail to parse request json, %s", err.Error()))
		return
	}
	// send rpc result
	service := rpc.GetRemoteService(request.ServiceName)
	rpcResponse, err := service.Request(reqData)
	if err != nil {
		returnError(c, fmt.Sprintf("invoke error: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rpcResponse,
	})
}

func returnError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 500,
		"data": msg,
	})
}