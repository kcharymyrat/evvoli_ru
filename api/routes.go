package api

import (
	"github.com/go-chi/chi/v5"
)

func (app *ApiApplication) Routes() *chi.Mux {

	mux := chi.NewRouter()

	// Healthcheck for the API
	mux.Get("/v1/healthcheck", app.HealthcheckHandler)

	// Languages
	mux.Post("/v1/languages", app.CreateLanguageHandler)
	mux.Get("/v1/languages", app.GetAllLanguagesHandler)

	// Categories
	mux.Post("/v1/categories", app.CreateCategoryHandler)
	mux.Get("/v1/categories", app.GetAllCategoriesHandler)
	mux.Get("/v1/category/{id}", app.GetCategoryHandler)

	// Products
	mux.Post("/v1/products", app.CreateProductHandler)
	mux.Get("/v1/products", app.GetAllProductsHandler)
	mux.Get("/v1/product/{id}", app.GetProductHandler)

	return mux
}
