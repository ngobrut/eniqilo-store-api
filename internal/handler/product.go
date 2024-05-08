package handler

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_error"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_validator"
)

func StringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

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
		ID:        StringPtr(qp.Get("id")),
		Name:      StringPtr(qp.Get("name")),
		Category:  StringPtr(qp.Get("category")),
		Sku:       StringPtr(qp.Get("sku")),
		Price:     StringPtr(qp.Get("price")),
		CreatedAt: StringPtr(qp.Get("createdAt")),
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

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var req request.UpdateProduct
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	req.ProductID, err = uuid.Parse(r.PathValue("id"))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  constant.HTTPStatusText(http.StatusNotFound),
		})
		response.Error(w, err)
		return
	}

	err = h.uc.UpdateProduct(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, http.StatusOK, nil)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(r.PathValue(("id")))
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  constant.HTTPStatusText(http.StatusNotFound),
		})
		response.Error(w, err)
		return
	}
	err = h.uc.DeleteProduct(r.Context(), productID)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.OK(w, http.StatusOK, nil)
}
