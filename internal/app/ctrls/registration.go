package ctrls

import (
	"github.com/stetsd/micro-brick/internal/domain/services"
	"net/http"
)

func Registration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		service, ok := req.Context().Value(services.ServiceUserName).(services.ServiceUser)

		if ok {
			service.Test()
		}

		w.Write([]byte("Mazafaka"))
	})
}
