package pgx

import (
	"context"
	"errors"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateColor(ctx context.Context, name string) (categories models.Color, err error) {

	sqlQuery := `Insert into colors(name,created_at,updated_at)values($1,$2,$3) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, name, time.Now(), time.Now()).Scan(&categories.ID, &categories.Name)

	return
}
func (d *db) GetAllColor(ctx context.Context) (colors []models.Color, err error) {
	sqlQuery := `SELECT
		id,
		name,
		created_at,
		updated_at
		FROM colors;`

	rows, err := d.postgres.Query(ctx, sqlQuery)

	colors = make([]models.Color, 0)
	for rows.Next() {
		var color models.Color
		err = rows.Scan(
			&color.ID,
			&color.Name,
			&color.CreatedAt,
			&color.UpdatedAt,
		)
		if err != nil {
			err = errors.New("rows.Scan error: " + err.Error())
			return
		}
		colors = append(colors, color)
	}

	return
}
