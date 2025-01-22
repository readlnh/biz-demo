package dal

import (
	"github.com/readlnh/biz-demo/gomall/demo/demo-proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
