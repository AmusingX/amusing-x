package model

import (
	"amusingx.fit/amusingx/services/ganymede/conf"
	"github.com/ItsWewin/superfactory/db/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)

var GanymedeDB *sqlx.DB

var once sync.Once

// InitMySQL 初始化数据库
func InitMySQL() {
	ganymedeDB := conf.Conf.Mysql.Ganymede
	mysqlDB := mysql.NewMysqlDB(ganymedeDB.DB, ganymedeDB.User, ganymedeDB.Password,
		ganymedeDB.Host, ganymedeDB.Port, ganymedeDB.Protocol, ganymedeDB.MaxOpenConns, ganymedeDB.MaxIdleConns, ganymedeDB.MaxLifeTime)

	if GanymedeDB != nil {
		return
	}
	
	once.Do(func() {
		GanymedeDB = mysqlDB.Connect()
	})
}

// MysqlDisConnect 关闭数据库
func MysqlDisConnect() {
	if GanymedeDB == nil {
		return
	}

	err := GanymedeDB.Close()
	if err != nil {
		panic(err)
	}
}
