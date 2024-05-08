package handler

import (
	"net/http"
	"strconv"

	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
)

func (h *Handler) SearchSKU(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()

	params := &request.SearchQuery{
		Name:     StringPtr(qp.Get("name")),
		Category: StringPtr(qp.Get("category")),
		Sku:      StringPtr(qp.Get("sku")),
		Price:    StringPtr(qp.Get("price")),
	}

	if limit, err := strconv.Atoi(qp.Get("limit")); err == nil {
		params.Limit = &limit
	}
	if offset, err := strconv.Atoi(qp.Get("offset")); err == nil {
		params.Offset = &offset
	}
	if inStock, err := strconv.ParseBool(qp.Get("inStock")); err == nil {
		params.InStock = &inStock
	}

	res, err := h.uc.SearchSKU(r.Context(), params)
	if err != nil {
		response.Error(w, err)
		return
	}
	response.OK(w, http.StatusOK, res)

}
