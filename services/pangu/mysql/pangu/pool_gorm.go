package pangu

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func InitGormDB() {
	InitMySQL()

	var err error
	GormDB, err = gorm.Open(mysql.New(mysql.Config{Conn: PanguDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
