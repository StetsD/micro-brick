package app

import "github.com/gorilla/mux"

func NewHttpRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}
