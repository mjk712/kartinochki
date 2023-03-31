package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mjk712/kartinochki/pkg/cash"
	"github.com/mjk712/kartinochki/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	// r.SkipClean(true)

	// cash это прекол про бабки или ошибся?)
	// докинь .gitignore и закинь туда .exe, не хранят в репе обычно
	cache := cash.NewLru(100) // взять из конфига бы
	router := routes.NewRouter(r, cache)
	router.KartinkiRoutes()
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Panic("boom")
	}
	// tgbot.startBot()
}
