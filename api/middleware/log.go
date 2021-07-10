package middleware

import (
	"fmt"
	"net/http"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rAddr := r.RemoteAddr
		method := r.Method
		path := r.URL.Path
		fmt.Printf("Remote: %s [%s] %s\n", rAddr, method, path)
		h.ServeHTTP(w, r)
	})
}
