package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddSubcategories ...
func (e *entity) AddSubcategories(ctx context.Context, categories models.Subcategories) (response models.Subcategories, code int, err error) {

	response, err = e.database.CreateSubcategories(ctx, categories)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
func (e *entity) GetAllSubcategories(ctx context.Context) (response []models.Subcategories, code int, err error) {

	response, err = e.database.GetAllSubcategories(ctx)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
