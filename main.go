package main

import (
	"encoding/json"
	"golite/router"
	"golite/server"
	"log"
	"net/http"
)



func main() {
	routes :=  router.New()


	routes.Use(LoggingMiddleware)
	routes.Use(okMiddleware)

	routes.Group("/user").
	Add("/info", testHandler).
	Add("/check", testHandler)


	routes.AddRoute("/test",testHandler)

	// routes.ListRoutes()

	httpServer := server.NewHTTPServer(":8000",routes)

	httpServer.Start()
	
}


func testHandler(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ad")
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request received: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r) // Call the next handler
    })
}

func okMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("ok: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r) // Call the next handler
    })
}