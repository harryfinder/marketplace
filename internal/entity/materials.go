package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddMaterials ...
func (e *entity) AddMaterials(ctx context.Context, categories models.Materials) (response models.Materials, code int, err error) {

	response, err = e.database.CreateMaterials(ctx, categories.Name)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}

func (e *entity) GetAllMaterials(ctx context.Context) (response []models.Materials, code int, err error) {

	response, err = e.database.GetAllMaterials(ctx)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
