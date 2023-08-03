package http

import (
	"context"
	"marketplace/internal/config"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type server struct {
	srv *http.Server
}

type Server interface {
	Serve(context.Context, string, config.Config, []Route) error
	Shutdown(context.Context) error
}

func NewServer() Server {
	return &server{
		srv: &http.Server{
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 120 * time.Second,
		},
	}
}

type Route struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}

func (s *server) Serve(ctx context.Context, address string, tlsCfg config.Config, routes []Route) error {
	s.srv.Addr = address
	handler := httprouter.New()
	for _, route := range routes {
		handler.Handle(route.Method, tlsCfg.App.BaseAddress+tlsCfg.App.Version+route.Path, route.Handler)
	}
	s.srv.Handler = addCors(handler)

	return s.srv.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func addCors(router *httprouter.Router) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodDelete, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		MaxAge:           10,
		AllowCredentials: true,
	}).Handler(router)
}
