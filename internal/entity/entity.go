package entity

import (
	"context"
	"marketplace/internal/database"
	"marketplace/internal/models"
)

type Entity interface {
	SignUp(context.Context, models.User) (*models.ResponseUserCheckingKey, int, error)
	SignIn(context.Context, string, string) (models.ResponseUser, int, error)
	AddCategories(context.Context, models.Categories) (models.Categories, int, error)
	GetAllCategories(context.Context) ([]models.Categories, int, error)
	AddSubcategories(context.Context, models.Subcategories) (models.Subcategories, int, error)
	GetAllSubcategories(context.Context) ([]models.Subcategories, int, error)
	AddBrands(context.Context, models.Brands) (models.Brands, int, error)
	GetAllBrands(context.Context) ([]models.Brands, int, error)
	AddSpecies(context.Context, models.Species) (models.Species, int, error)
	GetAllSpecies(context.Context) ([]models.Species, int, error)
	AddColor(context.Context, models.Color) (models.Color, int, error)
	GetAllColor(context.Context) ([]models.Color, int, error)
	AddSize(context.Context, models.Size) (models.Size, int, error)
	GetAllSize(context.Context) ([]models.Size, int, error)
	AddMaterials(context.Context, models.Materials) (models.Materials, int, error)
	GetAllMaterials(context.Context) ([]models.Materials, int, error)
	AddStatuses(context.Context, models.Statuses) (models.Statuses, int, error)
	GetAllStatuses(context.Context) ([]models.Statuses, int, error)
	AddProducts(context.Context, models.Products) (models.Products, int, error)
}

type entity struct {
	database database.Database
}

func New(database database.Database) Entity {
	return &entity{database: database}
}
