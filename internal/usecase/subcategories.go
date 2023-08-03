package usecase

import (
	"context"
	"marketplace/internal/models"
)

func (u *usecase) AddSubcategories(ctx context.Context, subcategories models.Subcategories) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.AddSubcategories(ctx, subcategories)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Категория успешно добавлена!"
	response.Payload = auth
	return response
}
func (u *usecase) GetAllSubcategories(ctx context.Context) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.GetAllSubcategories(ctx)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Найдена несколько subcategories"
	response.Payload = auth
	return response
}
