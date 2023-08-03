package usecase

import (
	"context"
	"marketplace/internal/models"
)

func (u *usecase) AddColor(ctx context.Context, color models.Color) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.AddColor(ctx, color)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Категория успешно добавлена!"
	response.Payload = auth
	return response
}
func (u *usecase) GetAllColor(ctx context.Context) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.GetAllColor(ctx)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Найдена несколько colors"
	response.Payload = auth
	return response
}
