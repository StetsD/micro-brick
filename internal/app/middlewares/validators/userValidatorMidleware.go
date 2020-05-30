package validators

import (
	"fmt"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"net/http"
)

func Login(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) http.HandlerFunc {

	return next
}

func Registration(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) http.HandlerFunc {
	serv := req.Context().Value(services.ServiceUserName)

	fmt.Printf("IN VALIDATOR : %v\n", serv)

	return next
}
