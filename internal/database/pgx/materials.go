package pgx

import (
	"context"
	"errors"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateMaterials(ctx context.Context, materials string) (m models.Materials, err error) {

	sqlQuery := `Insert into materials(name,created_at,updated_at)values($1,$2,$3) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, materials, time.Now(), time.Now()).Scan(&m.ID, &m.Name)

	return
}
func (d *db) GetAllMaterials(ctx context.Context) (sizes []models.Materials, err error) {
	sqlQuery := `SELECT
		id,
		name,
		created_at,
		updated_at
		FROM materials;`

	rows, err := d.postgres.Query(ctx, sqlQuery)

	sizes = make([]models.Materials, 0)
	for rows.Next() {
		var size models.Materials
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
