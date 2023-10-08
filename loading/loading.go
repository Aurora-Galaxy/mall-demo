package loading

import (
	"gin_mall/conf"
	"gin_mall/repository/cache"
	"gin_mall/repository/db/dao"
)

func Loading() {
	conf.InitConfig()
	dao.InitMySQL()
	cache.InitCache() //初始化Redis连接
}
