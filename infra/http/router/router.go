package router

import (
	"go-social-action/infra/controllers"
	"go-social-action/infra/http/middleware"
	"net/http"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

func New(volunteerController controllers.VolunteerController) http.Handler {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.StripSlashes)
	r.Use(middleware.ContentTypeResponse)
	r.Route("/volunteers", func(r chi.Router) {
		r.Post("/", volunteerController.Create)
	})
	return r
}
