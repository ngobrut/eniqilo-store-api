package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngobrut/eniqilo-store-api/internal/model"
	"github.com/ngobrut/eniqilo-store-api/internal/types/request"
	"github.com/ngobrut/eniqilo-store-api/internal/types/response"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User) error
	GetOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetOneUserByPhone(ctx context.Context, phone string) (*model.User, error)

	// product
	CreateProduct(ctx context.Context, data *model.Product) error
	FindOneProductByID(ctx context.Context, ID uuid.UUID) (*model.Product, error)
	FindProducts(ctx context.Context, params *request.ListProductQuery) ([]*response.ListProduct, error)
	UpdateProduct(ctx context.Context, req *request.UpdateProduct) error
	DeleteProduct(ctx context.Context, productID uuid.UUID) error

	// search
	SearchSKU(ctx context.Context, params *request.SearchQuery) ([]*response.SearchSKU, error)

	// customer
	CreateCustomer(ctx context.Context, data *model.Customer) error
	FindOneCustomerByID(ctx context.Context, ID uuid.UUID) (*model.Customer, error)
	FindCustomers(ctx context.Context, params *request.ListCustomerQuery) ([]*response.ListCustomer, error)

	// checkout
	CheckProductChekoutByIDs(ctx context.Context, IDs []string) (map[string]*model.Product, error)
	CreateInvoice(ctx context.Context, invoice *model.Invoice, invoiceProducts []*model.InvoiceProduct) error
	FindInvoices(ctx context.Context, params *request.ListInvoiceQuery) ([]*response.ListInvoice, error)
}
