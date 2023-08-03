package pgx

import (
	"context"
	"errors"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateBrands(ctx context.Context, brands models.Brands) (b models.Brands, err error) {

	sqlQuery := `Insert into brands(name,categories_id,created_at,updated_at)values($1,$2,$3,$4) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, brands.Name, brands.CategoriesID, time.Now(), time.Now()).Scan(&b.ID, &b.Name)

	return
}

func (d *db) GetAllBrands(ctx context.Context) (brands []models.Brands, err error) {
	sqlQuery := `SELECT
		id,
		name,
		categories_id
		FROM brands;`

	rows, err := d.postgres.Query(ctx, sqlQuery)

	brands = make([]models.Brands, 0)
	for rows.Next() {
		var brandies models.Brands
		err = rows.Scan(
			&brandies.ID,
			&brandies.Name,
			&brandies.CategoriesID,
		)
		if err != nil {
			err = errors.New("rows.Scan error: " + err.Error())
			return
		}
		brands = append(brands, brandies)
	}

	return
}
