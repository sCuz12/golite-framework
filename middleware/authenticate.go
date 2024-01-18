package middleware

import (
	"fmt"
	"net/http"
)


func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request) {

		fmt.Println("Authenticating..")
		 // If authenticated, call the next handler
		 next.ServeHTTP(w, r)
	})
}