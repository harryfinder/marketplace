package models

import "time"

type Subcategories struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	CategoriesID int64     `json:"categories_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
