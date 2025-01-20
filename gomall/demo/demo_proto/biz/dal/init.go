package dal

import (
	"github.com/readlnh/biz-demo/gomall/demo/demo-proto/biz/dal/mysql"
	"github.com/readlnh/biz-demo/gomall/demo/demo-proto/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
