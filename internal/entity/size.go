package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddSize ...
func (e *entity) AddSize(ctx context.Context, categories models.Size) (response models.Size, code int, err error) {

	response, err = e.database.CreateSize(ctx, categories.Name)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
func (e *entity) GetAllSize(ctx context.Context) (response []models.Size, code int, err error) {

	response, err = e.database.GetAllSize(ctx)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
