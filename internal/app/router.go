package app

import (
	"github.com/gorilla/mux"
	"github.com/stetsd/micro-brick/internal/app/ctrls"
	"github.com/stetsd/micro-brick/internal/app/middlewares"
	"github.com/stetsd/micro-brick/internal/app/middlewares/validators"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"github.com/stetsd/micro-brick/internal/tools"
)

func NewHttpRouter(serviceCollection *tools.ServiceCollection) *mux.Router {
	router := mux.NewRouter()

	if &serviceCollection.ServiceUser != nil {
		// Registration user
		routeRegistration := router.
			Path("/registration").
			Subrouter()
		routeRegistration.Methods("POST")
		routeRegistration.Use(
			middlewares.Log,
			middlewares.BodyParser,
			validators.Registration,
			middlewares.ServiceCtxInjector(services.ServiceUserName, serviceCollection),
			ctrls.Registration,
		)
	}

	return router
}
