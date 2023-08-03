package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @Summary Ping API Health
// @Security AdminAuth
// @Tags gateway-admin
// @Description ping service
// @Produce json
// @Success 200 "Ok"
// @Router /ping [get]
func (s *server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (s *server) swagger(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")).ServeHTTP(w, r)
}
