package models

import "time"

type Species struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	SubcategoriesID int64     `json:"subcategories_id"`
	BrandsID        int64     `json:"brands_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
