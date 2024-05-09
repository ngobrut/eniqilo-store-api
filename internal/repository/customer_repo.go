package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_error"
)

func (r *Repository) CreateCustomer(ctx context.Context, data *model.Customer) error {
	query := `INSERT INTO customers(name, phone) VALUES (@name, @phone) RETURNING customer_id, name, phone`
	args := pgx.NamedArgs{
		"name":  data.Name,
		"phone": data.Phone,
	}

	dest := []interface{}{
		&data.CustomerID,
		&data.Name,
		&data.Phone,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(dest...)
	if err != nil {
		if IsDuplicateError(err) {
			return custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  constant.HTTPStatusText(http.StatusConflict),
			})
		}
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return nil
}

func (r *Repository) FindCustomers(ctx context.Context, params *request.ListCustomerQuery) ([]*response.ListCustomer, error) {
	query := `SELECT customer_id, phone, name FROM customers`

	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	if params.Phone != nil {
		clause = append(clause, fmt.Sprintf(" phone ilike $%d", counter))
		args = append(args, *params.Phone+"%")
		counter++
	}
	if params.Name != nil {
		clause = append(clause, fmt.Sprintf(" name ilike $%d", counter))
		args = append(args, "%"+*params.Name+"%")
		counter++
	}

	if counter > 1 {
		query += " WHERE" + strings.Join(clause, " AND")
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*response.ListCustomer, 0)
	for rows.Next() {
		customer := &response.ListCustomer{}
		err = rows.Scan(
			&customer.CustomerID,
			&customer.Phone,
			&customer.Name,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, customer)
	}

	return res, nil
}
