package service

import (
	"encoding/json"
	"github.com/MarkLux/GOLD/api/restful/cache"
	"github.com/MarkLux/GOLD/api/restful/constant"
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/MarkLux/GOLD/api/restful/utils"
	"github.com/satori/go.uuid"
	"log"
	"sync"
)

const TokenPrefix = "user_token_"

// the login token service implements by cache
type TokenService struct {
	rClient *cache.RedisClient
}

var tokenInstance *TokenService
var tokenOnce sync.Once

func (s TokenService) CreateToken(user *orm.User) (token string, err error) {
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
	log.Println(TokenPrefix + token)
	err = s.rClient.Set(TokenPrefix+token, string(b), constant.LoginTokenExpiredTime)
	return
}

func (s TokenService) GetUserByToken(token string) (user orm.User, err error) {
	// read from redis
	userJson, err := s.rClient.Get(TokenPrefix + token)
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
