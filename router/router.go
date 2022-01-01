package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-lang-server/config"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

type Router struct {
	config config.RouterConfig
	logger *log.Logger
}

func NewRouter() *Router {
	return &Router{
		config: config.RouterConfig{},
		logger: log.New(os.Stdout, "", 1),
	}
}

func (router *Router) Start() {
	r := mux.NewRouter()

	srv := &http.Server{
		Handler: handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "X-Auth", "x-deviceid", "x-google-token-id", "x-isd-code", "x-phone-number"}),
			handlers.AllowedMethods([]string{http.MethodGet,
				http.MethodDelete,
				http.MethodPost,
				http.MethodPut,
				http.MethodOptions}))(r),
		Addr:         router.config.Address,
		WriteTimeout: router.config.WriteTimeout * time.Second,
		ReadTimeout:  router.config.ReadTimeout * time.Second,
	}
	router.logger.Println("Server Started")
	log.Fatal(srv.ListenAndServe())
}
