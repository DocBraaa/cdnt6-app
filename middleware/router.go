package middleware

import (
	"github.com/gorilla/mux"
	"github.com/music/model"
)

func NewRouter(routes model.Routes) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
