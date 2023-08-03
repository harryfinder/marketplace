package usecase

import (
	"context"
	"marketplace/internal/models"
)

func (u *usecase) AddBrands(ctx context.Context, brands models.Brands) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.AddBrands(ctx, brands)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Категория успешно добавлена!"
	response.Payload = auth
	return response
}
func (u *usecase) GetAllBrands(ctx context.Context) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.GetAllBrands(ctx)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Найдена несколько brands"
	response.Payload = auth
	return response
}
