package main

import (
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/conf"
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/mysql/amusingriskcontrol"
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/router"
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/rpcserv"
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/xredis"
	"github.com/ItsWewin/superfactory/powertrain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	powertrain.Run(conf.Conf, func(o *powertrain.Options) {
		o.InitFunc = InitFunc
		o.DeferFunc = DeferFunc
		o.RegisterRouter = func(mux *mux.Router) {
			router.Register(mux)
		}
	})

	conf.Conf.Print()
}

// 服务初始化时候执行
func InitFunc() {
	log.Println("amusing api sever listen:", conf.Conf.Addr)
	amusingriskcontrol.InitMySQL()

	xredis.InitRedis(conf.Conf.RedisAddr, conf.Conf.RedisPassword, conf.Conf.RedisDB)

	InitRPCServer()
}

// 服务执行完毕时候执行
func DeferFunc() {
	amusingriskcontrol.MysqlDisConnect()

	xredis.CloseRedis()
}

func InitRPCServer() {
	err := rpcserv.InitRPCServer()
	if err != nil {
		panic(err)
	}
}