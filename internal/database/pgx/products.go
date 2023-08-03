package pgx

import (
	"context"
	"marketplace/internal/models"
	"time"
)

func (d *db) CreateProducts(ctx context.Context, products models.Products) (m models.Products, err error) {

	sqlQuery := `Insert into products(name,price,old_price,season,inventory_number,gender_id,color_id,size_id,material_id,species_id,brands_id,status_id,created_at,updated_at)
					values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) returning id,name;
		`
	err = d.postgres.QueryRow(ctx, sqlQuery,
		products.Name,
		products.Price,
		products.OldPrice,
		products.Season,
		products.InventoryNumber,
		products.GenderID,
		products.ColorID,
		products.SizeID,
		products.MaterialID,
		products.SpeciesID,
		products.BrandsID,
		products.StatusesID,
		time.Now(),
		time.Now()).
		Scan(
			&m.ID,
			&m.Name)

	return
}
