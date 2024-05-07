package handler

import (
	"net/http"
	"strconv"

	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_validator"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req request.CreateProduct
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	res, err := h.uc.CreateProduct(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, http.StatusCreated, res)
}

func (h *Handler) GetListProduct(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()

	params := &request.ListProductQuery{
		ID:        stringPtr(qp.Get("id")),
		Name:      stringPtr(qp.Get("name")),
		Category:  stringPtr(qp.Get("category")),
		Sku:       stringPtr(qp.Get("sku")),
		Price:     stringPtr(qp.Get("price")),
		CreatedAt: stringPtr(qp.Get("createdAt")),
	}

	if limit, err := strconv.Atoi(qp.Get("limit")); err == nil {
		params.Limit = &limit
	}
	if offset, err := strconv.Atoi(qp.Get("offset")); err == nil {
		params.Offset = &offset
	}
	if isAvailable, err := strconv.ParseBool(qp.Get("isAvailable")); err == nil {
		params.IsAvailable = &isAvailable
	}
	if inStock, err := strconv.ParseBool(qp.Get("inStock")); err == nil {
		params.InStock = &inStock
	}

	res, err := h.uc.GetListProduct(r.Context(), params)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.OK(w, http.StatusOK, res)

}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
