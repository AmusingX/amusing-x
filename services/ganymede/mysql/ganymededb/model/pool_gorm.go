package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func InitGormDB() {
	InitMySQL()

	var err error
	GormDB, err = gorm.Open(mysql.New(mysql.Config{Conn: GanymedeDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
