package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	// chi_middleware "github.com/go-chi/chi/v5/middleware"

	"github.com/ngobrut/eniqilo-store-api/config"
	"github.com/ngobrut/eniqilo-store-api/internal/middleware"
	"github.com/ngobrut/eniqilo-store-api/internal/types/response"
	"github.com/ngobrut/eniqilo-store-api/internal/usecase"
)

type Handler struct {
	cnf *config.Config
	uc  usecase.IFaceUsecase
}

func InitHTTPHandler(cnf *config.Config, uc usecase.IFaceUsecase) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestLogger)
	r.Use(middleware.Recover)

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
				"app-name": "eniqilo-store-api-api",
			},
		})
	})

	r.Route("/v1", func(r chi.Router) {
		r.Route("/staff", func(user chi.Router) {
			user.Post("/register", h.Register)
			user.Post("/login", h.Login)

			user.Group(func(profile chi.Router) {
				profile.Use(middleware.Authorize(cnf.JWTSecret))
				profile.Get("/profile", h.GetProfile)
			})
		})

		r.Route("/product", func(product chi.Router) {
			product.Use(middleware.Authorize(cnf.JWTSecret))
			product.Post("/", h.CreateProduct)
			product.Get("/", h.GetListProduct)
			product.Put("/{id}", h.UpdateProduct)
			product.Delete("/{id}", h.DeleteProduct)
			product.Post("/checkout", h.Checkout)
			product.Get("/checkout/history", h.GetListInvoice)
		})

		r.Route("/product/customer", func(search chi.Router) {
			search.Get("/", h.SearchSKU)
		})

		r.Route("/customer", func(customer chi.Router) {
			customer.Use(middleware.Authorize(cnf.JWTSecret))
			customer.Post("/register", h.RegisterCustomer)
			customer.Get("/", h.GetListCustomer)
		})

	})

	return r
}
