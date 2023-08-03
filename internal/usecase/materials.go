package usecase

import (
	"context"
	"marketplace/internal/models"
)

func (u *usecase) AddMaterials(ctx context.Context, size models.Materials) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.AddMaterials(ctx, size)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Категория успешно добавлена!"
	response.Payload = auth
	return response
}
func (u *usecase) GetAllMaterials(ctx context.Context) models.Response {
	var (
		response = models.Response{
			Code: 200,
		}
	)

	auth, statusCode, err := u.entity.GetAllMaterials(ctx)
	if err != nil {
		response.Message = err.Error()
		response.Code = statusCode
		return response
	}
	response.Message = "Найдена несколько materials"
	response.Payload = auth
	return response
}
