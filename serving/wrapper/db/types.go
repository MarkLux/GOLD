package db

// base do
type GoldDO struct {
	// unique data id
	Id string `bson:"_id"`
	// would be rewritten into json
	Data interface{} `bson:"data"`
	// timestamp of create
	CreatedAt int64 `bson:"created_at"`
	// timestamp of update
	UpdatedAt int64 `bson:"updated_at"`
}

// DB Query param
type GoldDBQuery struct {
	Skip  int
	Limit int
	Param map[string]string
}

// interface for db client
type GoldDataBaseClient interface {
	NewSession(table string) (GoldDataBaseSession, error)
}

// interface for database session
// not support transaction yet.
type GoldDataBaseSession interface {
	// single data handlers
	Get(id string) (data GoldDO, err error)
	Insert(data interface{}) error
	Update(do GoldDO) error
	Delete(id string) error
	// batch data handlers
	Query(q GoldDBQuery) (data []GoldDO, err error)
	// close the session and connection.
	Close()
}

// custom errors
type AuthError struct {
	Message string
}

func (e AuthError) Error() string {
	return e.Message
}

type DBCommonError struct {
	Message string
}

func (e DBCommonError) Error() string {
	return e.Message
}
