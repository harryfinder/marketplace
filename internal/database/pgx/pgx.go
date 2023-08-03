package pgx

import (
	"context"
	"marketplace/internal/database"
	pkgpostgres "marketplace/pkg/storage/postgres"
)

type db struct {
	ctx      context.Context
	postgres pkgpostgres.Database
}

func New(postgres pkgpostgres.Database) database.Database {
	return &db{postgres: postgres}
}
