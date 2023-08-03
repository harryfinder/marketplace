package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddMaterials godoc
// @Summary      AddMaterials
// @Description  AddMaterials
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) AddMaterials(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		materials models.Materials
		response  = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&materials)
	if err != nil {
		w.WriteHeader(500)
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}
	response = s.usecase.AddMaterials(r.Context(), materials)
}

// GetAllMaterials godoc
// @Summary      GetAllMaterials
// @Description  GetAllMaterials
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) GetAllMaterials(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	response = s.usecase.GetAllMaterials(r.Context())
}
