package app

import (
	"github.com/gorilla/mux"
	"github.com/stetsd/micro-brick/internal/tools"
	"net/http"
)

func NewHttpRouter(serviceCollection tools.ServiceCollection) *mux.Router {
	router := mux.NewRouter()

	if serviceCollection.UserService != nil {
		router.HandleFunc("/registration", func(w http.ResponseWriter, req *http.Request) {
			//serviceCollection.UserService.Registration
		})
	}

	return router
}
