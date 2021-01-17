package mygin

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc using func return
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

func test() {

}

// New is the constructor of myGin.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// AddRoute Rule
func (engin *Engine) AddRoute(method string, pattern string, handleFunc HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("%s", key)
	engin.router[key] = handleFunc
}

// GET using
func (engin *Engine) GET(pattern string, handleFunc HandlerFunc) {
	engin.AddRoute("GET", pattern, handleFunc)
}

// POST Method
func (engin *Engine) POST(pattern string, handleFunc HandlerFunc) {
	engin.AddRoute("POST", pattern, handleFunc)
}

// Run Serve
func (engin *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engin)
}

// ServerHTTP
func (engin *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handleFunc, ok := engin.router[key]; ok {
		handleFunc(w, req)
	} else {
		fmt.Fprintf(w, "404 not found\n %s", req.URL.Path)
	}
}
