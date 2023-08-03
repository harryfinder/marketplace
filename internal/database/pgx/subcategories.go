package pgx

import (
	"context"
	"errors"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateSubcategories(ctx context.Context, categories models.Subcategories) (s models.Subcategories, err error) {

	sqlQuery := `Insert into subcategories(name,categories_id,created_at,updated_at)values($1,$2,$3,$4) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, categories.Name, categories.CategoriesID, time.Now(), time.Now()).Scan(&s.ID, &s.Name)

	return
}

func (d *db) GetAllSubcategories(ctx context.Context) (subcategories []models.Subcategories, err error) {
	sqlQuery := `SELECT
		id,
		name,
		categories_id
		FROM subcategories;`

	rows, err := d.postgres.Query(ctx, sqlQuery)

	subcategories = make([]models.Subcategories, 0)
	for rows.Next() {
		var subcategory models.Subcategories
		err = rows.Scan(
			&subcategory.ID,
			&subcategory.Name,
			&subcategory.CategoriesID,
		)
		if err != nil {
			err = errors.New("rows.Scan error: " + err.Error())
			return
		}
		subcategories = append(subcategories, subcategory)
	}

	return
}
