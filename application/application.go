package application

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"

	"westminster/application/adapter/rest"
	"westminster/application/adapter/rest/handler"
	"westminster/application/drivers"
	"westminster/application/drivers/sqldb"
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
	err := godotenv.Load(".env-local")
	if err != nil {
		log.Fatalf("loading env file: %s", err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	connStr := drivers.EnsureEnv(drivers.PostgresConnString)
	_ = sqldb.DatabasePool(context.Background(), connStr)

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

func (e *Engine) ShutDownHTTPServer(ctx context.Context) {
	e.HTTPServer.Shutdown(ctx)
}
