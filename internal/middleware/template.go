// {{ .MiddlewareName }}.go

package middleware

import (
    "net/http"
)

// {{ .MiddlewareName }} is a middleware that performs some operations.
func {{ .MiddlewareName }}(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Middleware logic here

        // Call the next handler in the chain
        next.ServeHTTP(w, r)
    })
}
