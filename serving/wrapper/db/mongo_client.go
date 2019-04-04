package db

import (
	"fmt"
	"github.com/MarkLux/GOLD/serving/wrapper/constant"
	"gopkg.in/mgo.v2"
	"time"
)

// to avoid insert empty object id, create a do for insert
type GoldInsertDO struct {
	Data      interface{} `bson:"data"`
	CreatedAt int64       `bson:"created_at"`
	UpdatedAt int64       `bson:"updated_at"`
}

// export method
func NewMongoClient(db string, user string, pwd string) *GoldMongoClient {
	return &GoldMongoClient{
		DataBase:  db,
		AuthUser:  user,
		AuthToken: pwd,
	}
}

// the mongodb client using mgo.
type GoldMongoClient struct {
	DataBase  string
	AuthUser  string
	AuthToken string
}

func (c *GoldMongoClient) NewSession(table string) (GoldDataBaseSession, error) {
	session := &GoldMongoSession{
		authUser:  c.AuthUser,
		authToken: c.AuthToken,
		dataBase:  c.DataBase,
		table:     table,
	}
	if err := session.init(); err != nil {
		return nil, err
	}
	return session, nil
}

// the mongodb session using mgo.
type GoldMongoSession struct {
	authUser  string
	authToken string
	dataBase  string
	table     string
	// for each request, use one single session
	session *mgo.Session
}

func (s *GoldMongoSession) init() (err error) {
	s.session, err = mgo.Dial(s.getMongoDialUrl())
	return
}

func (s *GoldMongoSession) Get(id string) (data GoldDO, err error) {
	data = GoldDO{}
	err = s.session.DB(s.dataBase).C(s.table).FindId(id).One(&data)
	return
}

func (s *GoldMongoSession) Insert(data interface{}) (err error) {
	now := time.Now().Unix()
	err = s.session.DB(s.dataBase).C(s.table).Insert(&GoldInsertDO{
		CreatedAt: now,
		UpdatedAt: now,
		Data:      data,
	})
	return
}

func (s *GoldMongoSession) Update(do GoldDO) (err error) {
	do.UpdatedAt = time.Now().Unix()
	err = s.session.DB(s.dataBase).C(s.table).UpdateId(do.Id, do)
	return
}

func (s *GoldMongoSession) Delete(id string) error {
	return s.session.DB(s.dataBase).C(s.table).RemoveId(id)
}

func (s *GoldMongoSession) Query(q GoldDBQuery) (data []GoldDO, err error) {
	err = s.session.DB(s.dataBase).C(s.table).Find(q.Param).Skip(q.Skip).Limit(q.Limit).All(&data)
	return
}

func (s *GoldMongoSession) Close() {
	s.session.Close()
}

func (s *GoldMongoSession) getMongoDialUrl() string {
	dbHost := constant.GoldMongoPrimaryEndPoint + ":" + constant.GoldMongoServicePort
	return fmt.Sprintf("mongodb://%s:%s@%s/%s", s.authUser, s.authToken, dbHost, s.dataBase)
}
