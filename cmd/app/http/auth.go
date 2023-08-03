package http

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// SignUp godoc
// @Summary      SignUp
// @Description  SignUp
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserSignUp  true  "auth"
// @Success      200   {object}  models.Response
// @Security     ApiKeyAuth
// @Router       /auth/sign-up [POST]
// SignUp ...
func (s *server) SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		user     models.User
		response = models.Response{
			Code: 200,
		}
	)

	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(500)
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}
	response = s.usecase.SignUp(r.Context(), user)
}

// SignIn godoc
// @Summary      SignIn
// @Description  SignIn
// @Tags         authorization
// @Accept       json
// @Produce      json
// @Param        auth  body      models.UserLogin  true  "auth"
// @Success      200   {object}  models.ResponseUser
// @Security     ApiKeyAuth
// @Router       /auth/sign-in [POST]
// SignIn ...
func (s *server) SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		login    models.UserLogin
		response = models.Response{
			Code: 200,
		}
	)
	// utils.CorsOptions(w, r)

	defer response.Send(w, r)

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		log.Println("controllers, auth user parsing request failed:", err)
		return
	}

	response = s.usecase.SignIn(r.Context(), login.Login, login.Password)
	return
}
