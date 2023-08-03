package entity

import (
	"context"
	"fmt"
	jwt "github.com/AbduvokhidovRustamzhon/jwt/pkg/cmd"
	"github.com/google/uuid"
	"log"
	"marketplace/internal/models"
	"marketplace/internal/utils"
	"net/http"
	"time"
)

func (e *entity) SignUp(ctx context.Context, user models.User) (*models.ResponseUserCheckingKey, int, error) {
	var ch models.CheckEmail

	m := map[string]interface{}{"email": user.Email}
	err := utils.Validate(m)
	if err != nil {
		log.Println("ERROR", err)
		return nil, http.StatusBadRequest, err
	}
	//user.Password = utils.GeneratePasswordHash(user.Password)
	//user.RepeatPassword = utils.GeneratePasswordHash(user.RepeatPassword)
	if user.FullName == "" {
		user.FullName = "user"
	}
	user.ActivePhone = true
	user.Status = true
	user.Role = 2
	user.NumberPhone = "null"
	user.ActiveEmail = true
	user.Password = utils.RandomPasswordInt(4)
	fmt.Println(user.Password)
	user.Password = utils.GeneratePasswordHash(user.Password)
	createUser, err := e.database.CreateUser(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, http.StatusBadRequest, err
	}

	userResponse, err := e.database.GetRoleUser(ctx, &createUser)
	if err != nil {
		log.Println(err)
		return nil, http.StatusBadRequest, err
	}

	forgerHash := uuid.New().String()
	ch.Key = forgerHash
	ch.UserId = user.ID
	ch.Email = user.Email

	//check, err := e.database.GetAllCheckEmails(ctx, ch)
	//
	//if ch.UserId != check.UserId {
	//	check, err = e.database.CreateCheckEmail(ctx, check)
	//	if err != nil {
	//		log.Println(err)
	//		return nil, http.StatusBadRequest, err
	//	}
	//}
	//
	//if ch.UserId == check.UserId {
	//	err = e.database.UpdateCheckEmail(ctx, check)
	//	if err != nil {
	//		log.Println(err)
	//		return nil, http.StatusBadRequest, err
	//	}
	//}
	//
	//go func() {
	//	err = utils.ActivationSendByEmailMessage(user, forgerHash, viper.GetString("integration.url_by_activation"))
	//}()

	return userResponse, http.StatusOK, err
}

// SignIn ...
func (e *entity) SignIn(ctx context.Context, login string, password string) (response models.ResponseUser, code int, err error) {
	var responseUser models.Payload
	m := map[string]interface{}{"login": login, "password": password}
	err = utils.Validate(m)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}

	password = utils.GeneratePasswordHash(password)
	userResponse, err := e.database.GetUser(ctx, login, password)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}

	responseUser.ID = userResponse.Id
	responseUser.Exp = time.Now().Add(time.Hour * 10).Unix()
	response.Id = userResponse.Id
	response.Token, err = jwt.Encode(responseUser, jwt.Secret("SECRET"))
	response.Role = userResponse.Role
	return
}
