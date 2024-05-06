package handler

import (
	"net/http"

	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_validator"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req request.Register
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	res, err := h.uc.Register(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, http.StatusCreated, res)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.Login
	err := custom_validator.ValidateStruct(r, &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	res, err := h.uc.Login(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.OK(w, http.StatusOK, res)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
