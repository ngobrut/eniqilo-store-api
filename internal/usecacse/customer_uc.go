package usecacse

import (
	"context"

	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
)

func (u *Usecase) RegisterCustomer(ctx context.Context, req *request.RegisterCustomer) (*response.RegisterCustomer, error) {
	customer := &model.Customer{
		Name:  req.Name,
		Phone: req.Phone,
	}

	err := u.repo.CreateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	res := &response.RegisterCustomer{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Phone:      customer.Phone,
	}

	return res, nil

}
