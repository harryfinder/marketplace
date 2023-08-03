package pgx

import (
	"context"
	"errors"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateSpecies(ctx context.Context, species models.Species) (s models.Species, err error) {

	sqlQuery := `Insert into species(name,subcategories_id,brands_id,created_at,updated_at)values($1,$2,$3,$4,$5) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery, species.Name, species.SubcategoriesID, species.BrandsID, time.Now(), time.Now()).Scan(&s.ID, &s.Name)

	return
}

func (d *db) GetAllSpecies(ctx context.Context) (species []models.Species, err error) {
	sqlQuery := `SELECT
		id,
		name,
		subcategories_id,
		brands_id,
		created_at,
		updated_at
		FROM species;`

	rows, err := d.postgres.Query(ctx, sqlQuery)

	species = make([]models.Species, 0)
	for rows.Next() {
		var specie models.Species
		err = rows.Scan(
			&specie.ID,
			&specie.Name,
			&specie.SubcategoriesID,
			&specie.BrandsID,
			&specie.CreatedAt,
			&specie.UpdatedAt,


		)
		if err != nil {
			err = errors.New("rows.Scan error: " + err.Error())
			return
		}
		species = append(species, specie)
	}

	return
}
