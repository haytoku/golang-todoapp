package core_http_response

import (
	"net/http"
)

var (
	StatusCodeInitialized = -1
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriterWrapper {

	return &ResponseWriterWrapper{
		ResponseWriter: w,
		statusCode:     StatusCodeInitialized,
	}
}

func (rw *ResponseWriterWrapper) WriteHeader(StatusCode int) {

	rw.ResponseWriter.WriteHeader(StatusCode)
	rw.statusCode = StatusCode
}

func (rw *ResponseWriterWrapper) GetStatusCodeOrpanic() int {
	if rw.statusCode == StatusCodeInitialized {
		panic("status code is not set")
	}

	return rw.statusCode
}
