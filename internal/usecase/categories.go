package usecase

import (
	"context"
	"marketplace/internal/models"
)

func (u *usecase) AddCategories(ctx context.Context, categories models.Categories) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.AddCategories(ctx, categories)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Категория успешно добавлена!"
	response.Payload = auth
	return response
}

func (u *usecase) GetAllCategories(ctx context.Context) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.GetAllCategories(ctx)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Найдена несколько categories"
	response.Payload = auth
	return response
}
