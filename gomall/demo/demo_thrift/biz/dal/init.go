package dal

import (
	"github.com/readlnh/biz-demo/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/readlnh/biz-demo/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
