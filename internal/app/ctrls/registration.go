package ctrls

import (
	"fmt"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"net/http"
)

func Registration(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) http.HandlerFunc {

	serv := req.Context().Value(services.ServiceUserName)
	fmt.Printf("From CONTROLLER : %v\n", serv)
	return next
}
