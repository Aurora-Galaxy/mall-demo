package dao

import (
	"gin_mall/conf"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"strings"
	"time"
)

var (
	_db *gorm.DB
)

func InitMySQL() {
	mConfig := conf.Config.MySql["default"]
	// 读写分离，实现需要将读写设置不同数据库，此处只设置一个
	pathRead := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":",
		mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")
	pathWrite := strings.Join([]string{mConfig.UserName, ":", mConfig.Password, "@tcp(", mConfig.DbHost, ":",
		mConfig.DbPort, ")/", mConfig.DbName, "?charset=" + mConfig.Charset + "&parseTime=true"}, "")

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead, // 设置数据库的一些配置信息
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     //  通过删除并新建的方式重命名，用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置数据库连接
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{ //使用单数命名数据库表，mysql自动迁移创建表时，默认是复数
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 连接池中的最大空闲连接数，超出数量会关闭
	sqlDB.SetMaxOpenConns(100)                 // 同时可以有最多100个活动连接
	sqlDB.SetConnMaxLifetime(time.Second * 30) //连接在闲置30秒后将被关闭并重新创建
	_db = db
	_ = _db.Use(dbresolver. //dbresolver实现读写分离的库
				Register(dbresolver.Config{
			// `db2` 作为 sources，`db3`、`db4` 作为 replicas
			Sources:  []gorm.Dialector{mysql.Open(pathWrite)},                      // 写操作
			Replicas: []gorm.Dialector{mysql.Open(pathRead), mysql.Open(pathRead)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                    // 设置了负载均衡策略，这里使用了 dbresolver.RandomPolicy{}，表示随机选择一个副本进行读操作
		}))

	_db = _db.Set("gorm:table_options", "charset=utf8mb4")
	err = migrate() //自动迁移
	if err != nil {
		panic(err)
	}
}

// NewDBClient 在不修改全局 _db 对象的情况下创建一个具有特定上下文的数据库客户端。
// 每个上下文可以具有不同的超时、取消等属性，而不会影响全局数据库连接。
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
