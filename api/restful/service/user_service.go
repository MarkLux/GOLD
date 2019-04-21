package service

import (
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/utils"
	)

// interface for user service
type UserService struct {
	BaseService
}

func (s UserService) Register(user orm.User) (id int64, err error) {
	// 校验用户是否已经存在（用户名 + 邮箱）
	nameCnt, err := s.Engine.Count(&orm.User{Name: user.Name})
	if nameCnt > 0 {
		err = errors.GenRegisteredError("用户名")
		return
	}
	emailCnt, err := s.Engine.Count(&orm.User{Email: user.Email})
	if emailCnt > 0 {
		err = errors.GenRegisteredError("邮箱")
		return
	}
	// 加密密码
	user.Password = utils.GenMD5(user.Password)
	user.InitTime()
	_, err  = s.Engine.Insert(user)
	if err != nil {
		return
	}
	id = user.Id
	return
}

func (s UserService) Login(email string, pwd string)  {

}
