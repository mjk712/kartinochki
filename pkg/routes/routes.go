package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mjk712/kartinochki/cash"
	"github.com/mjk712/kartinochki/pkg/controllers"
)

type Router struct {
	router *mux.Router
	cache  *cash.LRU
}

func NewRouter(router *mux.Router, cache *cash.LRU) Router {
	return Router{
		router: router,
		cache:  cache,
	}
}

func (r Router) KartinkiRoutes() {
	c := controllers.NewController(r.cache)
	r.router.HandleFunc("/imageredact/{imageX}/{imageY}/{imgUrl:.*}", c.ImageShow).Methods(http.MethodGet)
}
