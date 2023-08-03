package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddColor ...
func (e *entity) AddColor(ctx context.Context, categories models.Color) (response models.Color, code int, err error) {

	response, err = e.database.CreateColor(ctx, categories.Name)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
func (e *entity) GetAllColor(ctx context.Context) (response []models.Color, code int, err error) {

	response, err = e.database.GetAllColor(ctx)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
