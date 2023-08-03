package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddMaterials ...
func (e *entity) AddProducts(ctx context.Context, products models.Products) (response models.Products, code int, err error) {

	response, err = e.database.CreateProducts(ctx, products)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
