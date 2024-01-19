// Kostas.go

package middleware

import (
    "net/http"
)

// Kostas is a middleware that performs some operations.
func Kostas(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Middleware logic here

        // Call the next handler in the chain
        next.ServeHTTP(w, r)
    })
}
