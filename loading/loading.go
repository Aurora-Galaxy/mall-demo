package loading

import (
	"gin_mall/conf"
	"gin_mall/repository/db/dao"
)

func Loading() {
	conf.InitConfig()
	dao.InitMySQL()
}
