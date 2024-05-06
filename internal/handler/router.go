package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	// chi_middleware "github.com/go-chi/chi/v5/middleware"

	"github.com/ngobrut/eniqlo-store-api/config"
	"github.com/ngobrut/eniqlo-store-api/internal/middleware"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/internal/usecacse"
)

type Handler struct {
	cnf *config.Config
	uc  usecacse.IFaceUsecase
}

func InitHTTPHandler(cnf *config.Config, uc usecacse.IFaceUsecase, logger *logrus.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger)
	r.Use(middleware.Recover(logger))

	h := Handler{
		cnf: cnf,
		uc:  uc,
	}

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(response.CONTENT_TYPE_HEADER, response.CONTENT_TYPE_JSON)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Message: "Error",
			Error: &response.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "please check url",
			},
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(response.CONTENT_TYPE_HEADER, response.CONTENT_TYPE_JSON)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.JsonResponse{
			Success: true,
			Message: "Success",
			Data: map[string]interface{}{
				"app-name": "eniqlo-store-api-api",
			},
		})
	})

	r.Route("/v1", func(r chi.Router) {
		r.Route("/user", func(user chi.Router) {
			user.Post("/register", h.Register)
			user.Post("/login", h.Login)

			user.Group(func(profile chi.Router) {
				profile.Use(middleware.Authorize(cnf.JWTSecret))
				profile.Get("/profile", h.GetProfile)
			})
		})

		r.Route("/product", func(product chi.Router) {
			product.Post("/", h.CreateProduct)
		})

		r.Route("/example", func(example chi.Router) {
			example.Use(middleware.Authorize(cnf.JWTSecret))
			example.Get("/", h.Example)
		})
	})

	return r
}
