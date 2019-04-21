package service

import (
	"github.com/MarkLux/GOLD/api/restful/orm"
	"github.com/go-xorm/xorm"
)

type BaseService struct {
	Engine *xorm.Engine
}

func (s BaseService) Insert(do orm.TimeRecordable) error {
	_, err := s.Engine.Insert(do)
	return err
}

func (s BaseService) Update(do interface{}) error {
	s.Engine.Update()
}