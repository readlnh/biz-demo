package dal

import (
	"github.com/readlnh/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/readlnh/biz-demo/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
