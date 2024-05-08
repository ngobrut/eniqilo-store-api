package usecacse

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_error"
)

func (u *Usecase) CreateProduct(ctx context.Context, req *request.CreateProduct) (*response.CreateProduct, error) {
	product := &model.Product{
		Name:        req.Name,
		Sku:         req.Sku,
		Category:    req.Category,
		ImageUrl:    req.ImageUrl,
		Notes:       req.Notes,
		Price:       req.Price,
		Stock:       *req.Stock,
		Location:    req.Location,
		IsAvailable: *req.IsAvailable,
	}

	err := u.repo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	res := &response.CreateProduct{
		ProductID: product.ProductID.String(),
		CreatedAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return res, nil
}

func (u *Usecase) GetListProduct(ctx context.Context, req *request.ListProductQuery) ([]*response.ListProduct, error) {
	res, err := u.repo.FindProducts(ctx, req)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}
	return res, nil
}

func (u *Usecase) UpdateProduct(ctx context.Context, req *request.UpdateProduct) error {
	_, err := u.repo.FindOneProductByID(ctx, req.ProductID)
	if err != nil {
		return err
	}

	err = u.repo.UpdateProduct(ctx, req)
	if err != nil {
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return nil

}

func (u *Usecase) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	_, err := u.repo.FindOneProductByID(ctx, productID)
	if err != nil {
		return err
	}
	err = u.repo.DeleteProduct(ctx, productID)
	if err != nil {
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}
	return nil
}
