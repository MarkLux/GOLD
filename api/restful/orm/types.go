package orm

// the data mapping structures for orm

type User struct {
	Id int64
	Name string
	Email string
	CreatedAt int64
	UpdatedAt int64
	AddOn string
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
	CreatedAt int64
	UpdatedAt int64
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