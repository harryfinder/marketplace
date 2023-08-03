package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddStatuses godoc
// @Summary      AddStatuses
// @Description  AddStatuses
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) AddStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		statuses models.Statuses
		response = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&statuses)
	if err != nil {
		w.WriteHeader(500)
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}
	response = s.usecase.AddStatuses(r.Context(), statuses)
}

// GetAllStatuses godoc
// @Summary      GetAllStatuses
// @Description  GetAllStatuses
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) GetAllStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	response = s.usecase.GetAllStatuses(r.Context())
}
