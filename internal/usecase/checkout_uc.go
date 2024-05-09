package usecase

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

func (u *Usecase) Checkout(ctx context.Context, req *request.Checkout) error {
	customerID, err := uuid.Parse(req.CustomerID)
	if err != nil {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "customerId is not found",
		})
		return err
	}
	_, err = u.repo.FindOneCustomerByID(ctx, customerID)
	if err != nil {
		return err
	}

	// check if product_id is uuid
	productIDs := make([]string, 0)
	for _, product := range req.ProductDetails {
		_, err := uuid.Parse(product.ProductID)
		if err != nil {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  "one of productIds is not found",
			})
			return err
		}
		productIDs = append(productIDs, product.ProductID)
	}
	// check if product_id exists in db
	dbProducts, err := u.repo.CheckProductChekoutByIDs(ctx, productIDs)
	if err != nil {
		return err
	}

	// check stock
	var invoiceProducts = make([]*model.InvoiceProduct, 0)
	var totalPrice int = 0
	for _, reqProduct := range req.ProductDetails {
		reqProductID, _ := uuid.Parse(reqProduct.ProductID)
		dbProduct := dbProducts[reqProduct.ProductID]
		if dbProduct.Stock < reqProduct.Quantity {
			return custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusBadRequest,
				Message:  "one of productIds stock is not enough",
			})
		}
		invoiceProducts = append(invoiceProducts,
			&model.InvoiceProduct{
				ProductID: reqProductID,
				Quantity:  reqProduct.Quantity,
				Price:     dbProduct.Price,
			})
		totalPrice += dbProduct.Price * reqProduct.Quantity
	}

	// check total price
	if totalPrice > req.Paid {
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "paid is not enough based on all bought product",
		})
	}

	// check if change is not right
	if change := req.Paid - totalPrice; change != *req.Change {
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "change is not right",
		})
	}

	invoice := &model.Invoice{
		CustomerID: customerID,
		Paid:       req.Paid,
		Change:     *req.Change,
		TotalPrice: totalPrice,
	}

	err = u.repo.CreateInvoice(ctx, invoice, invoiceProducts)
	if err != nil {
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return nil
}

func (u *Usecase) GetListInvoice(ctx context.Context, req *request.ListInvoiceQuery) ([]*response.ListInvoice, error) {
	res, err := u.repo.FindInvoices(ctx, req)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}
	return res, nil
}
