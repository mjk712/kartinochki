package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mjk712/kartinochki/pkg/controllers"
)

var KartinkiRoutes = func(router *mux.Router) {

	router.HandleFunc("/imageredact/{imageX}/{imageY}/{imgUrl:.*}", controllers.ImageShow).Methods(http.MethodGet)
	//router.HandleFunc("/image/{imageX}/{imageY}/{imgUrl}", controllers.ImageShow).Methods(http.MethodGet)
}
