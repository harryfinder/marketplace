package usecase

import (
	"context"
	"marketplace/internal/models"
)

func (u *usecase) AddSize(ctx context.Context, size models.Size) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.AddSize(ctx, size)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Категория успешно добавлена!"
	response.Payload = auth
	return response
}
func (u *usecase) GetAllSize(ctx context.Context) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.GetAllSize(ctx)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Найдена несколько sizes"
	response.Payload = auth
	return response
}
