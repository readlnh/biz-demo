package dal

import (
	"github.com/readlnh/biz-demo/gomall/app/frontend/biz/dal/mysql"
	"github.com/readlnh/biz-demo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
