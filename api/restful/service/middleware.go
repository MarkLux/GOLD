package service

import (
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NeedLoginCheck(c *gin.Context) (bool, *orm.User) {
	header := c.Request.Header
	tokenStr := header.Get("token")
	if len(tokenStr) <= 0 {
		c.JSON(http.StatusOK, errors.GenNeedLoginError())
		return false, nil
	}
	tokenService := GetTokenService()
	user, err := tokenService.GetUserByToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, errors.GenNeedLoginError())
		return false, nil
	}
	return true, &user
}