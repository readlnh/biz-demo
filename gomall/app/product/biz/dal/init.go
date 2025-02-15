package dal

import (
	"github.com/readlnh/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/readlnh/biz-demo/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
