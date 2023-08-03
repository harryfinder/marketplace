package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddSpecies ...
func (e *entity) AddSpecies(ctx context.Context, categories models.Species) (response models.Species, code int, err error) {

	response, err = e.database.CreateSpecies(ctx, categories)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}

func (e *entity) GetAllSpecies(ctx context.Context) (response []models.Species, code int, err error) {

	response, err = e.database.GetAllSpecies(ctx)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
