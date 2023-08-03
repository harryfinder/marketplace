package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddProducts godoc
// @Summary      AddProducts
// @Description  AddProducts
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) AddProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		products models.Products
		response = models.Response{
			Code: 200,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		w.WriteHeader(500)
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}
	response = s.usecase.AddProducts(r.Context(), products)
}

// GetAllProducts godoc
// @Summary      GetAllProducts
// @Description  GetAllProducts
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
func (s *server) GetAllProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		products models.Products
		response = models.Response{
			Code: 200,
		}
	)
	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		w.WriteHeader(500)
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}

	response = s.usecase.GetAllProducts(r.Context(), products)
}
