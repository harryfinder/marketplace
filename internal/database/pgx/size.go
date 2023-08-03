package pgx

import (
	"context"
	"errors"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateSize(ctx context.Context, size string) (sizes models.Size, err error) {

	sqlQuery := `Insert into size(name,created_at,updated_at)values($1,$2,$3) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, size, time.Now(), time.Now()).Scan(&sizes.ID, &sizes.Name)

	return
}

func (d *db) GetAllSize(ctx context.Context) (sizes []models.Size, err error) {
	sqlQuery := `SELECT
		id,
		name,
		created_at,
		updated_at
		FROM size;`

	rows, err := d.postgres.Query(ctx, sqlQuery)

	sizes = make([]models.Size, 0)
	for rows.Next() {
		var size models.Size
		err = rows.Scan(
			&size.ID,
			&size.Name,
			&size.CreatedAt,
			&size.UpdatedAt,
		)
		if err != nil {
			err = errors.New("rows.Scan error: " + err.Error())
			return
		}
		sizes = append(sizes, size)
	}

	return
}
