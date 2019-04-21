package orm

import "time"

// base interface
type TimeRecordable interface {
	SetCreatedAt(int64)
	SetUpdatedAt(int64)
}

type BaseDO struct {
	CreatedAt int64
	UpdatedAt int64
}

func (b BaseDO) InitTime() {
	t := time.Now().Unix()
	b.CreatedAt = t
	b.UpdatedAt = t
}

func (b BaseDO) UpdateTime() {
	b.UpdatedAt = time.Now().Unix()
}

// the data mapping structures for orm

type User struct {
	Id int64
	Name string
	Email string
	Password string
	AddOn string
	BaseDO
}

type FunctionService struct {
	Id int64
	CreatorId int64
	CreatorName string
	ServiceName string
	GitRemote string
	GitBranch string
	GitHead string
	Status string
	LastOperation int64
	AddOn string
	MinInstance int
	MaxInstance int
	BaseDO
}

type OperateLogs struct {
	Id int64
	ServiceId int64
	OperatorId int64
	Type string
	Start int64
	Update int64
	End int64
	CurrentAction string
	Log string
	BaseDO
}

// TableName() Method for X-ORM
func (User) TableName() string {
	return "users"
}

func (FunctionService) TableName() string {
	return "function_services"
}

func (OperateLogs) TableName() string {
	return "operate_logs"
}