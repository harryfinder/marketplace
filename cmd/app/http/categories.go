package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddCategories godoc
// @Summary      AddCategories
// @Description  AddCategories
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) AddCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		categories models.Categories
		response   = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&categories)
	if err != nil {
		w.WriteHeader(500)
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}
	response = s.usecase.AddCategories(r.Context(), categories)
}

// GetAllCategories godoc
// @Summary      GetAllCategories
// @Description  GetAllCategories
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) GetAllCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	response = s.usecase.GetAllCategories(r.Context())
}
