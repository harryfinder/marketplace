package usecase

import (
	"context"
	"marketplace/internal/models"
)

func (u *usecase) SignUp(ctx context.Context, user models.User) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.SignUp(ctx, user)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Вы успешно зарегистрировались!"
	response.Payload = auth
	response.Code = statusCode
	return response
}

func (u *usecase) SignIn(ctx context.Context, login string, password string) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.SignIn(ctx, login, password)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}

	response.Payload = auth
	return response
}
