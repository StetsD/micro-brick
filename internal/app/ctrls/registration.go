package ctrls

import (
	"encoding/json"
	"github.com/stetsd/micro-brick/internal/app/constants"
	"github.com/stetsd/micro-brick/internal/app/schemas"
	"github.com/stetsd/micro-brick/internal/domain/errors"
	"github.com/stetsd/micro-brick/internal/domain/services"
	"github.com/stetsd/micro-brick/internal/tools/helpers"
	"net/http"
)

func Registration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		serviceUser, ok := req.Context().Value(services.ServiceUserName).(services.ServiceUser)

		if !ok {
			helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			return
		}

		err := serviceUser.Registration()

		if err != nil {
			_, ok := err.(errors.ErrorUser)
			if ok {
				errorText := err.Error()
				helpers.Throw(w, http.StatusBadRequest, &errorText)
			} else {
				helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			}
			return
		}

		result := schemas.HttpResult{Result: "success"}
		jsonBody, err := json.Marshal(result)

		if err != nil {
			helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonBody)

		if err != nil {
			helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			return
		}
	})
}
