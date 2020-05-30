package tools

import (
	"context"
	"net/http"
)

func ServiceCtxInjector(serviceName string, req *http.Request, service interface{}) {
	ctx := context.WithValue(req.Context(), serviceName, service)
	req.WithContext(ctx)
}
