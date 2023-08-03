package usecase

import (
	"context"
	"marketplace/internal/models"
)

func (u *usecase) AddSpecies(ctx context.Context, species models.Species) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.AddSpecies(ctx, species)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Категория успешно добавлена!"
	response.Payload = auth
	return response
}

func (u *usecase) GetAllSpecies(ctx context.Context) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.GetAllSpecies(ctx)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Найдена несколько brands"
	response.Payload = auth
	return response
}
