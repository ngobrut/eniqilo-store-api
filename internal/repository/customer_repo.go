package repository

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
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
