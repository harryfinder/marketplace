package pgx

import (
	"context"
	"errors"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateStatuses(ctx context.Context, statuses string) (s models.Statuses, err error) {

	sqlQuery := `Insert into status(name,created_at,updated_at)values($1,$2,$3) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, statuses, time.Now(), time.Now()).Scan(&s.ID, &s.Name)

	return
}

func (d *db) GetAllStatuses(ctx context.Context) (statuses []models.Statuses, err error) {
	sqlQuery := `SELECT
		id,
		name,
		created_at,
		updated_at
		FROM status;`

	rows, err := d.postgres.Query(ctx, sqlQuery)

	statuses = make([]models.Statuses, 0)
	for rows.Next() {
		var status models.Statuses
		err = rows.Scan(
			&status.ID,
			&status.Name,
			&status.CreatedAt,
			&status.UpdatedAt,
		)
		if err != nil {
			err = errors.New("rows.Scan error: " + err.Error())
			return
		}
		statuses = append(statuses, status)
	}

	return
}
