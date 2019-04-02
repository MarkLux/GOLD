package gold

import (
	"github.com/MarkLux/GOLD/serving/rpc/goldrpc"
)

// the model of user
// use annotation `bson` to control the key saved in db(for mongo).
type UserModel struct {
	Name string `bson:"name"`
	Sex  string `bson:"sex"`
	Mail string `bson:"mail"`
}

// the biz function
func (s *GoldService) Handle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error {

}
