package services

import (
	"fmt"
	"github.com/stetsd/micro-brick/internal/domain/repositoryInterfaces"
	"net/http"
)

type ServiceUser struct {
	UserStore repositoryInterfaces.UserStore
}

const ServiceUserName = "ServiceUser"

func (su *ServiceUser) Login(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) http.HandlerFunc {
	// TODO: validation module, logging
	// TODO: DB include
	// TODO: id, email hardcode

	//user := &models.User{1, name, ""}

	//return &models.User{}
	return next
}

func (su *ServiceUser) Registration(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) http.HandlerFunc {
	// TODO: validation module, logging
	// TODO: DB include
	// TODO: id, email hardcode

	//user := &models.User{1, name, email}

	//return &models.User{}
	return next
}

func (su *ServiceUser) Test() {
	fmt.Println("TEEEEESt")
}
