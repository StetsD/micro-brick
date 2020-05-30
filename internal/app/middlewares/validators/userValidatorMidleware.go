package validators

import (
	"fmt"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"net/http"
)

func Registration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		serv := req.Context().Value(services.ServiceUserName)

		fmt.Printf("IN VALIDATOR : %v\n", serv)
		next.ServeHTTP(w, req)
	})
}
