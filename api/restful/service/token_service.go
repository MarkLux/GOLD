package service

import (
	"encoding/json"
	"github.com/MarkLux/GOLD/api/restful/cache"
	"github.com/MarkLux/GOLD/api/restful/constant"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/MarkLux/GOLD/api/restful/utils"
	"github.com/satori/go.uuid"
	"sync"
)

// the login token service implements by cache
type TokenService struct {
	rClient *cache.RedisClient
}

var tokenInstance *TokenService
var tokenOnce sync.Once

func (s TokenService) createToken(user *orm.User) (token string, err error) {
	// generate user token
	uuid := uuid.NewV4()
	if err != nil {
		return
	}
	token = utils.GenMD5(uuid.String())
	b, err := json.Marshal(user)
	if err != nil {
		return
	}
	// save into redis
	err = s.rClient.Set(token, string(b), constant.LoginTokenExpiredTime)
	return
}

func (s TokenService) getUserByToken(token string) (user orm.User, err error) {
	// read from redis
	userJson, err := s.rClient.Get(token)
	if err != nil {
		// if the key not exist, an error would be thrown, handle it as not login.
		return
	}
	err = json.Unmarshal([]byte(userJson), &user)
	return
}

func GetTokenService() *TokenService {
	tokenOnce.Do(func() {
		tokenInstance = &TokenService{
			rClient: cache.GetRedisClient(),
		}
	})
	return tokenInstance
}