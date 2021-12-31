package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-lang-server/config"
	"net/http"
)

const portNumber = ":8080"

type Router struct {
	config config.RouterConfig
}

func NewRouter() *Router {
	return &Router{
		config: config.RouterConfig{},
	}
}

func (router *Router) Start() {
	r := mux.NewRouter()
	fmt.Println("listening on port %d", portNumber)
	http.ListenAndServe(portNumber, r)
}
