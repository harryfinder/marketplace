package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddBrands ...
func (e *entity) AddBrands(ctx context.Context, brands models.Brands) (response models.Brands, code int, err error) {

	response, err = e.database.CreateBrands(ctx, brands)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
func (e *entity) GetAllBrands(ctx context.Context) (response []models.Brands, code int, err error) {

	response, err = e.database.GetAllBrands(ctx)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
