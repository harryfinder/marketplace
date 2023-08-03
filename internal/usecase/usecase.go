package usecase

import (
	"context"
	"marketplace/internal/entity"
	"marketplace/internal/models"
)

type Usecase interface {
	SignUp(ctx context.Context, user models.User) models.Response
	SignIn(ctx context.Context, login string, password string) models.Response
	AddCategories(context.Context, models.Categories) models.Response
	GetAllCategories(ctx context.Context) models.Response
	AddSubcategories(context.Context, models.Subcategories) models.Response
	GetAllSubcategories(ctx context.Context) models.Response
	AddBrands(context.Context, models.Brands) models.Response
	GetAllBrands(ctx context.Context) models.Response
	AddSpecies(context.Context, models.Species) models.Response
	GetAllSpecies(ctx context.Context) models.Response
	AddColor(context.Context, models.Color) models.Response
	GetAllColor(ctx context.Context) models.Response
	AddSize(context.Context, models.Size) models.Response
	GetAllSize(ctx context.Context) models.Response
	AddStatuses(context.Context, models.Statuses) models.Response
	GetAllStatuses(ctx context.Context) models.Response
	AddMaterials(context.Context, models.Materials) models.Response
	GetAllMaterials(ctx context.Context) models.Response
	AddProducts(context.Context, models.Products) models.Response
	GetAllProducts(ctx context.Context, products models.Products) models.Response
}

type usecase struct {
	entity entity.Entity
}

func New(entity entity.Entity) Usecase {
	return &usecase{
		entity: entity,
	}
}
