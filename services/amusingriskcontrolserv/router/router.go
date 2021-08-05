package router

import (
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/api/accountriskcontrol"
	"amusingx.fit/amusingx/services/amusingriskcontrolserv/api/pong"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(mux *mux.Router) {
	mux.HandleFunc("/v1/amusingriskcontrolserv/pong", pong.Pong).Methods(http.MethodGet)
	mux.HandleFunc("/v1/amusingriskcontrolserv/logincontrol", accountriskcontrol.LoginControl).Methods(http.MethodPost)
}
