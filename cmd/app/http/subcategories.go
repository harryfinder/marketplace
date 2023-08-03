package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddSubcategories godoc
// @Summary      AddSubcategories
// @Description  AddSubcategories
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) AddSubcategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		subcategories models.Subcategories
		response      = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&subcategories)
	if err != nil {
		w.WriteHeader(500)
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}
	response = s.usecase.AddSubcategories(r.Context(), subcategories)
}

// GetAllSubcategories godoc
// @Summary      GetAllSubcategories
// @Description  GetAllSubcategories
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) GetAllSubcategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	response = s.usecase.GetAllSubcategories(r.Context())
}
