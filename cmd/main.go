package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mjk712/kartinochki/cash"
	"github.com/mjk712/kartinochki/config"
	"github.com/mjk712/kartinochki/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	cs := config.CascheSize()

	cache := cash.NewLru(cs)
	router := routes.NewRouter(r, cache)
	router.KartinkiRoutes()
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Panic("boom")
	}

}
