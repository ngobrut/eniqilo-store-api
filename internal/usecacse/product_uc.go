package usecacse

import (
	"context"

	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
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
		CreatedAt: product.CreatedAt,
	}

	return res, nil
}
