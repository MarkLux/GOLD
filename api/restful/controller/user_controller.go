package controller

import (
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/MarkLux/GOLD/api/restful/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// the json request dto here:
type UserRegisterRequest struct {
	Email string `json:"email" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// the user controller
type UserController struct {
	userService *service.UserService
}

func (c UserController) RegisterUser(ctx *gin.Context) {
	var req UserRegisterRequest
	var err error
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.GenValidationError())
		return
	}
	user := orm.User{
		Name: req.UserName,
		Email: req.Email,
		Password: req.Password,
	}
	id, err := c.userService.Register(user)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	data := make(map[string]interface{})
	data["userId"] = id
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func (c UserController) LoginUser(ctx *gin.Context) {
	var req UserLoginRequest
	var err error
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, errors.GenValidationError())
		return
	}
	user, token, err := c.userService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	data["user"] = user
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func (c UserController) GetLoginUser(ctx *gin.Context) {
	hasLogin, user := service.NeedLoginCheck(ctx)
	if !hasLogin {
		return
	}
	user.Password = ""
	data := make(map[string]interface{})
	data["user"] = user
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func NewUserController() UserController {
	return UserController{userService: service.GetUserService()}
}