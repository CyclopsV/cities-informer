package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func CreateRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	r.Route("/", func(r chi.Router) {
		r.Get("/", getCityByIdHandler)
		r.Put("/", createCityHandler)
		r.Delete("/", deleteCityHandler)
		r.Patch("/", updateCityHandler)
	})

	return r
}
