package main

import (
	"encoding/json"
	"golite/internal/router"
	"golite/internal/server"
	"net/http"
)



func main() {
	routes :=  router.New()

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