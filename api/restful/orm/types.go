package orm

import "time"

type BaseDO struct {
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

func (b *BaseDO) InitTime() {
	t := time.Now().Unix()
	b.CreatedAt = t
	b.UpdatedAt = t
}

func (b *BaseDO) UpdateTime() {
	b.UpdatedAt = time.Now().Unix()
}

// the data mapping structures for orm

type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	AddOn string `json:"addOn"`
	BaseDO `xorm:"extends"`
}

type FunctionService struct {
	Id int64 `json:"id"`
	CreatorId int64 `json:"creatorId"`
	CreatorName string `json:"creatorName"`
	ServiceName string `json:"serviceName"`
	GitRepo string `json:"gitRepo"`
	GitBranch string `json:"gitBranch"`
	GitHead string `json:"gitHead"`
	GitMaintainer string `json:"gitMaintainer"`
	Status string `json:"statue"`
	LastOperation int64 `json:"lastOperation"`
	AddOn string `json:"addOn"`
	MinInstance int `json:"minInstance"`
	MaxInstance int `json:"maxInstance"`
	BaseDO `xorm:"extends"`
}

type OperateLogs struct {
	Id int64 `json:"id"`
	ServiceId int64 `json:"serviceId"`
	OperatorId int64 `json:"operatorId"`
	Type string `json:"type"`
	Start int64 `json:"start"`
	Update int64 `json:"update"`
	End int64 `json:"end"`
	CurrentAction string `json:"currentAction"`
	Log string `json:"log"`
	OriginBranch string `json:"originBranch"`
	OriginVersion string `json:"originVersion"`
	TargetBranch string `json:"targetBranch"`
	TargetVersion string `json:"targetVersion"`
	BaseDO `xorm:"extends"`
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