package main

import (
	"net/http"

	"github.com/music/controller"
	"github.com/music/middleware"
	"github.com/music/store"
)

func main() {
	store := store.MockStore{}
	track := controller.NewTrack(store)
	router := middleware.NewRouter(track.Routes)

	http.ListenAndServe(":3000", router)
}
