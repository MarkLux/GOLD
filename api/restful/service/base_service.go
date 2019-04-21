package service

import (
	"github.com/go-xorm/xorm"
)

type BaseService struct {
	Engine *xorm.Engine
}