package http

import (
	"context"
	"marketplace/cmd/app/controllers"
	"marketplace/internal/config"
	"marketplace/internal/usecase"
	pkghttp "marketplace/pkg/controller/http"
	"net/http"
)

type server struct {
	usecase usecase.Usecase
	srv     pkghttp.Server
}

func NewController(usecase usecase.Usecase, srv pkghttp.Server) controllers.Controller {
	return &server{
		usecase: usecase,
		srv:     srv,
	}
}
func (s *server) Serve(ctx context.Context, address string, tlsCfg config.Config) error {
	return s.srv.Serve(ctx, address, tlsCfg, []pkghttp.Route{

		{Method: http.MethodGet, Path: "/ping", Handler: s.ping},
		// swagger
		{Method: http.MethodGet, Path: "/swagger/*any", Handler: s.swagger},

		// <--------- auth ---------->
		{Method: http.MethodPost, Path: "/auth/sign-up", Handler: s.SignUp},
		{Method: http.MethodPost, Path: "/auth/sign-in", Handler: s.SignIn},

		// <-------Categories ----------->
		{Method: http.MethodPost, Path: "/app/categories", Handler: s.AddCategories},
		{Method: http.MethodGet, Path: "/app/categories/all", Handler: s.GetAllCategories},

		// <-------Subcategories ----------->
		{Method: http.MethodPost, Path: "/app/subcategories", Handler: s.AddSubcategories},
		{Method: http.MethodGet, Path: "/app/subcategories/all", Handler: s.GetAllSubcategories},

		// <-------Brands ----------->
		{Method: http.MethodPost, Path: "/app/brands", Handler: s.AddBrands},
		{Method: http.MethodGet, Path: "/app/brands/all", Handler: s.GetAllBrands},

		// <-------Species ----------->
		{Method: http.MethodPost, Path: "/app/species", Handler: s.AddSpecies},
		{Method: http.MethodGet, Path: "/app/species/all", Handler: s.GetAllSpecies},

		// <-------Color ----------->
		{Method: http.MethodPost, Path: "/app/color", Handler: s.AddColor},
		{Method: http.MethodGet, Path: "/app/color/all", Handler: s.GetAllColor},

		// <-------Size ----------->
		{Method: http.MethodPost, Path: "/app/size", Handler: s.AddSize},
		{Method: http.MethodGet, Path: "/app/size/all", Handler: s.GetAllSize},

		// <-------Statuses ----------->
		{Method: http.MethodPost, Path: "/app/status", Handler: s.AddStatuses},
		{Method: http.MethodGet, Path: "/app/status/all", Handler: s.GetAllStatuses},

		// <-------Materials ----------->
		{Method: http.MethodPost, Path: "/app/material", Handler: s.AddMaterials},
		{Method: http.MethodGet, Path: "/app/material/all", Handler: s.GetAllMaterials},

		// <-------Materials ----------->
		{Method: http.MethodPost, Path: "/app/products", Handler: s.AddProducts},
		{Method: http.MethodGet, Path: "/app/products/all", Handler: s.GetAllProducts},
	})
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
