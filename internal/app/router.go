package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stetsd/micro-brick/internal/tools"
)

func NewHttpRouter(serviceCollection tools.ServiceCollection) *mux.Router {
	router := mux.NewRouter()

	for key, service := range serviceCollection {
		//switch key {
		//	case "user":
		//		router.HandleFunc("/registration", *service.)
		//}

		fmt.Printf("%v", service, key)
	}

	return router
}
