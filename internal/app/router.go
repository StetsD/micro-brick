package app

import (
	"github.com/gorilla/mux"
	"github.com/stetsd/micro-brick/internal/app/ctrls"
	"github.com/stetsd/micro-brick/internal/app/middlewares/validators"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"github.com/stetsd/micro-brick/internal/tools"
	"net/http"
)

func plug(w http.ResponseWriter, req *http.Request) {}

func NewHttpRouter(serviceCollection tools.ServiceCollection) *mux.Router {
	router := mux.NewRouter()

	if serviceCollection.ServiceUser != nil {
		// Registration user
		router.HandleFunc("/registration", func(w http.ResponseWriter, req *http.Request) {
			tools.ServiceCtxInjector(services.ServiceUserName, req, &serviceCollection.ServiceUser)
			validators.Registration(
				w,
				req,
				ctrls.Registration(w, req, plug),
			)
		})
	}

	return router
}
