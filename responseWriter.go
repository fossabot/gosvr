package main

import "net/http"

type loggingResponseWriter struct {
	Writer     http.ResponseWriter
	Written    bool
	StatusCode int
}

func (w *loggingResponseWriter) Header() http.Header {
	return w.Writer.Header()
}

func (w *loggingResponseWriter) Write(data []byte) (int, error) {
	if !w.Written {
		w.WriteHeader(http.StatusOK)
	}
	ret, err := w.Writer.Write(data)
	return ret, err
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.Written = true
	w.StatusCode = code
	w.Writer.WriteHeader(code)
}

func httpHandlerWrapper(h simpleHTTPServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lrw := loggingResponseWriter{
			Writer:  w,
			Written: false,
		}
		h.ServeHTTP(lrw, r)
	}
}
