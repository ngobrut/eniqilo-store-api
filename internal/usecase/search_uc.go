package usecase

import (
	"context"
	"net/http"

	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_error"
)

func (u *Usecase) SearchSKU(ctx context.Context, req *request.SearchQuery) ([]*response.SearchSKU, error) {
	res, err := u.repo.SearchSKU(ctx, req)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}
	return res, nil
}
