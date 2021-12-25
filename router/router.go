package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-lang-server/config"
	"net/http"
)

type Router struct {
	config config.RouterConfig
}

func NewRouter() *Router {
	return &Router{}
}

func (router *Router) Start() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
