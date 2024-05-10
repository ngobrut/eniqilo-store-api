package usecase

import (
	"context"
	"net/http"

	"github.com/ngobrut/eniqilo-store-api/internal/model"
	"github.com/ngobrut/eniqilo-store-api/internal/types/request"
	"github.com/ngobrut/eniqilo-store-api/internal/types/response"
	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
	"github.com/ngobrut/eniqilo-store-api/pkg/custom_error"
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

func (u *Usecase) GetListCustomer(ctx context.Context, req *request.ListCustomerQuery) ([]*response.ListCustomer, error) {
	res, err := u.repo.FindCustomers(ctx, req)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return res, nil
}
