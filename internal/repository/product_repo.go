package repository

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_error"
)

func (r *Repository) CreateProduct(ctx context.Context, data *model.Product) error {
	query := `INSERT INTO products(name, sku, category, image_url, notes, price, stock, location, is_available)
            VALUES (@name, @sku, @category, @image_url, @notes, @price, @stock, @location, @is_available)
            RETURNING product_id, created_at, updated_at`
	args := pgx.NamedArgs{
		"name":         data.Name,
		"sku":          data.Sku,
		"category":     data.Category,
		"image_url":    data.ImageUrl,
		"notes":        data.Notes,
		"price":        data.Price,
		"stock":        data.Stock,
		"location":     data.Location,
		"is_available": data.IsAvailable,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(&data.ProductID, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})

	}

	return nil
}
