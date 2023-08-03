package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddCategories ...
func (e *entity) AddCategories(ctx context.Context, categories models.Categories) (response models.Categories, code int, err error) {

	response, err = e.database.CreateCategories(ctx, categories.Name)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}

func (e *entity) GetAllCategories(ctx context.Context) (response []models.Categories, code int, err error) {

	response, err = e.database.GetAllCategories(ctx)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
