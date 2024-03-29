package main

import (
	"encoding/json"
	"golite/middleware"
	"golite/router"
	"golite/server"
	"net/http"
)



func main() {
	routes :=  router.New()

	//Apply to this group only the middleware
	routes.Group("/user").
	AddMiddleware(middleware.Logging).
	Add("/info", testHandler).
	Add("/check", testHandler).
	Add("/bio",testHandler)


	routes.Group("/owner").
    Add("/check", testHandler) // Add routes within the owner group.


	routes.AddRoute("/test",testHandler)
	routes.AddRoute("/xoxo",testHandler)

	routes.ListRoutes()

	httpServer := server.NewHTTPServer(":8000",routes)

	httpServer.Start()
	
}


func testHandler(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ad")
}

// func LoggingMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         log.Printf("Request received: %s %s", r.Method, r.URL.Path)
//         next.ServeHTTP(w, r) // Call the next handler
//     })
// }
