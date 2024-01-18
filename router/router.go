package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes map[string] http.HandlerFunc
	middlewareChain []MiddleWare
}

type MiddleWare func(http.Handler) http.Handler


type RouteGroup struct { 
	prefix string 
	router *Router
}


//**Act like constructor **/
func New()*Router {
	return &Router {
		routes: make(map[string] http.HandlerFunc),
	}
}

func (r *Router) AddRoute(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) Group(prefix string) *RouteGroup {
    
	return &RouteGroup{
		prefix: prefix,
		router: r,
	}
}
/**
* Adding path to specific route group prefix and return RouteGroup option to allow chaining
**/
func (rg *RouteGroup) Add(path string , handler http.HandlerFunc) *RouteGroup {
	fullPath := rg.prefix + path
    rg.router.routes[fullPath] = handler

	return rg
}

func (r *Router) Use(mw MiddleWare) {
	//append to middlewares the new middleware
	r.middlewareChain = append(r.middlewareChain,mw)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, exists := r.routes[req.URL.Path]
    
    if !exists {
        http.NotFound(w, req)
        return
    }

	finalHandler := http.Handler(handler)

    for i := len(r.middlewareChain) - 1; i >= 0; i-- {
        finalHandler = r.middlewareChain[i](finalHandler)
    }

    finalHandler.ServeHTTP(w, req)
}

func (r *Router) ListRoutes() {
	for path,_ := range r.routes {
		fmt.Println(path)
	}
}



