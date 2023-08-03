package entity

import (
	"context"
	"log"
	"marketplace/internal/models"
	"net/http"
)

// AddStatuses ...
func (e *entity) AddStatuses(ctx context.Context, categories models.Statuses) (response models.Statuses, code int, err error) {

	response, err = e.database.CreateStatuses(ctx, categories.Name)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}

func (e *entity) GetAllStatuses(ctx context.Context) (response []models.Statuses, code int, err error) {

	response, err = e.database.GetAllStatuses(ctx)
	if err != nil {
		log.Println(err)
		return response, http.StatusBadRequest, err
	}
	return
}
