package apiserver

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(stausCode int) {
	w.code = stausCode
	w.ResponseWriter.WriteHeader(stausCode)
}