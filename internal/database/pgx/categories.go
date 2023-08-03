package pgx

import (
	"context"
	"errors"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateCategories(ctx context.Context, name string) (categories models.Categories, err error) {

	sqlQuery := `Insert into categories(name,created_at,updated_at)values($1,$2,$3) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, name, time.Now(), time.Now()).Scan(&categories.ID, &categories.Name)

	return
}

func (d *db) GetAllCategories(ctx context.Context) (categories []models.Categories, err error) {
	sqlQuery := `SELECT
		id,
		name
		FROM categories;`

	rows, err := d.postgres.Query(ctx, sqlQuery)

	categories = make([]models.Categories, 0)
	for rows.Next() {
		var category models.Categories
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			err = errors.New("rows.Scan error: " + err.Error())
			return
		}
		categories = append(categories, category)
	}

	return
}
