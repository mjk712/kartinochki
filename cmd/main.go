package main

import (
	//"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mjk712/kartinochki/pkg/routes"
	//"github.com/mjk712/kartinochki/pkg/tgbot"
)

func main() {
	r := mux.NewRouter()
	//r.SkipClean(true)
	routes.KartinkiRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", r)
	//tgbot.startBot()
}
