package pangu

import "amusingx.fit/amusingx/services/pangu/conf"

func MockGormDB() {
	conf.Mock()
	InitGormDB()
}
