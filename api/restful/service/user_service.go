package service

import (
	"github.com/MarkLux/GOLD/api/restful/errors"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/MarkLux/GOLD/api/restful/utils"
	"github.com/go-xorm/xorm"
	"sync"
)

// interface for user service
type UserService struct {
	engine       *xorm.Engine
	tokenService *TokenService
}

var userInstance *UserService
var userOnce sync.Once

func (s UserService) Register(user orm.User) (id int64, err error) {
	// check if the user is exist
	nameCnt, err := s.engine.Count(&orm.User{Name: user.Name})
	if nameCnt > 0 {
		err = errors.GenRegisteredError("用户名")
		return
	}
	emailCnt, err := s.engine.Count(&orm.User{Email: user.Email})
	if emailCnt > 0 {
		err = errors.GenRegisteredError("邮箱")
		return
	}
	// encrypt the password
	user.Password = utils.GenMD5(user.Password)
	user.InitTime()
	_, err = s.engine.Insert(&user)
	if err != nil {
		return
	}
	id = user.Id
	return
}

func (s UserService) Login(email string, pwd string) (token string, err error) {
	u := &orm.User{
		Email: email,
	}
	has, err := s.engine.Get(u)
	if err != nil {
		return
	}
	if !has {
		err = errors.GenUserNotExistedError()
		return
	}
	if utils.CheckMD5(pwd, u.Password) {
		token, err = s.tokenService.createToken(u)
	} else {
		err = errors.GenPwdError()
		return
	}
	return
}

func GetUserService() *UserService {
	userOnce.Do(func() {
		userInstance = &UserService{
			tokenService: GetTokenService(),
			engine: orm.GetOrmEngine(),
		}
	})
	return userInstance
}