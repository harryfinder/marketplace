package models

import "time"

type Products struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Price           float64   `json:"price"`
	OldPrice        float64   `json:"old_price"`
	Season          string    `json:"season"`
	InventoryNumber string    `json:"inventory_number"`
	GenderID        int64     `json:"gender_id"`
	ColorID         int64     `json:"color_id"`
	SizeID          int64     `json:"size_id"`
	MaterialID      int64     `json:"material_id"`
	SpeciesID       int64     `json:"species_id"`
	BrandsID        int64     `json:"brands_id"`
	StatusesID      int64     `json:"statuses_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
