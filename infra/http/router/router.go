package router

import (
	"go-social-action/infra/controllers"
	"go-social-action/infra/http/middleware"
	"net/http"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

func New(volunteerController controllers.VolunteerController, socialActionController controllers.SocialActionController) http.Handler {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.StripSlashes)
	r.Use(middleware.JSONContentTypeResponse)
	r.Route("/volunteers", func(r chi.Router) {
		r.Post("/", volunteerController.Create)
		r.Get("/{id}", volunteerController.GetByID)
	})
	r.Route("/social-actions", func(r chi.Router) {
		r.Post("/", socialActionController.Create)
	})
	return r
}
