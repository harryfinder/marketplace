package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddSpecies godoc
// @Summary      AddSpecies
// @Description  AddSpecies
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) AddSpecies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		species  models.Species
		response = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&species)
	if err != nil {
		w.WriteHeader(500)
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}
	response = s.usecase.AddSpecies(r.Context(), species)
}

// GetAllSpecies godoc
// @Summary      GetAllSpecies
// @Description  GetAllSpecies
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) GetAllSpecies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	response = s.usecase.GetAllSpecies(r.Context())
}
