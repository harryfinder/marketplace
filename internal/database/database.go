package database

import (
	"context"
	"marketplace/internal/models"
)

// Database Можно ли разделит по интерфейсам а там храним методов?
type Database interface {
	GetUser(ctx context.Context, login string, password string) (models.ResponseUser, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetRoleUser(ctx context.Context, user *models.User) (*models.ResponseUserCheckingKey, error)
	CreateCheckEmail(ctx context.Context, ch *models.CheckEmail) (*models.CheckEmail, error)
	GetUsersByID(ctx context.Context, id int64) (*models.User, error)
	CreateCategories(context.Context, string) (models.Categories, error)
	GetAllCategories(context.Context) ([]models.Categories, error)
	CreateSubcategories(context.Context, models.Subcategories) (models.Subcategories, error)
	GetAllSubcategories(context.Context) ([]models.Subcategories, error)
	CreateBrands(context.Context, models.Brands) (models.Brands, error)
	GetAllBrands(context.Context) ([]models.Brands, error)
	CreateSpecies(context.Context, models.Species) (models.Species, error)
	GetAllSpecies(context.Context) ([]models.Species, error)
	CreateColor(context.Context, string) (models.Color, error)
	GetAllColor(context.Context) ([]models.Color, error)
	CreateSize(context.Context, string) (models.Size, error)
	GetAllSize(context.Context) ([]models.Size, error)
	CreateStatuses(context.Context, string) (models.Statuses, error)
	GetAllStatuses(context.Context) ([]models.Statuses, error)
	CreateMaterials(context.Context, string) (models.Materials, error)
	GetAllMaterials(context.Context) ([]models.Materials, error)
	CreateProducts(context.Context, models.Products) (models.Products, error)
}

//Get user by id
//Delete user by id
//update user by id
//get all users
