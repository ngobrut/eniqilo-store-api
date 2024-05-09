package handler

import (
	"net/http"
	"strconv"

	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_validator"
)

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req request.Checkout
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	err = h.uc.Checkout(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, http.StatusOK, nil)
}

func (h *Handler) GetListInvoice(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()

	params := &request.ListInvoiceQuery{
		CustomerID: StringPtr(qp.Get("customerId")),
		CreatedAt:  StringPtr(qp.Get("createdAt")),
	}

	if limit, err := strconv.Atoi(qp.Get("limit")); err == nil {
		params.Limit = &limit
	}
	if offset, err := strconv.Atoi(qp.Get("offset")); err == nil {
		params.Offset = &offset
	}

	res, err := h.uc.GetListInvoice(r.Context(), params)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, http.StatusOK, res)
}
