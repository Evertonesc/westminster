package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"westminster/application/adapter/rest"
	"westminster/application/adapter/rest/handler"
	"westminster/application/drivers"
	"westminster/application/drivers/database"
)

type (
	Engine struct {
		HTTPServer *http.Server
	}

	HTTPServer struct {
		r chi.Router
	}
)

func New() *Engine {
	engine := &Engine{}
	engine.loadDependencies()
	return engine
}

func (e *Engine) loadDependencies() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	mongoHost := drivers.EnsureEnv(drivers.MongoDBHost)
	mongoPort := drivers.EnsureEnv(drivers.MongoDBPort)

	_, err := database.Connect(
		context.Background(),
		fmt.Sprintf("%s:%s", mongoHost, mongoPort),
	)
	if err != nil {
		log.Fatal(err)
	}

	memberHandler := handler.NewMemberHandler(nil)

	rest.RegisterRoutes(r, memberHandler)

	e.HTTPServer = &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
}

func (e *Engine) StartHTTPServer() {
	e.HTTPServer.ListenAndServe()
}
