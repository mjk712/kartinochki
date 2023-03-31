package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mjk712/kartinochki/pkg/cash"
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
	r.router.HandleFunc("/imageredact/{imageX}/{imageY}/{imgUrl:.*}", controllers.ImageShow).Methods(http.MethodGet)
	// router.HandleFunc("/image/{imageX}/{imageY}/{imgUrl}", controllers.ImageShow).Methods(http.MethodGet)
}
