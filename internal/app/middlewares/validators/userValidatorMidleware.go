package validators

import (
	"context"
	"net/http"
)

func Login() {

}

func Registration(ctx context.Context, next http.Handler) http.Handler {

	return next
}
