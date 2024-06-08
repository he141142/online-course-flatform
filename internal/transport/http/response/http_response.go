package response

import (
	"net/http"
)

type HttpResponse interface {
	ResponseError(status int, err error)
	ResponseSuccess(status int, data interface{})
}

type httpResponse struct {
	http.ResponseWriter
}

func NewHttpResponse(w http.ResponseWriter) HttpResponse {
	return &httpResponse{w}
}
