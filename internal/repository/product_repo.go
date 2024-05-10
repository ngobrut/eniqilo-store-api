package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/eniqilo-store-api/internal/model"
	"github.com/ngobrut/eniqilo-store-api/internal/types/request"
	"github.com/ngobrut/eniqilo-store-api/internal/types/response"
	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
	"github.com/ngobrut/eniqilo-store-api/pkg/custom_error"
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

func (r *Repository) FindOneProductByID(ctx context.Context, ID uuid.UUID) (*model.Product, error) {
	product := &model.Product{}
	query := `SELECT  *
		FROM products WHERE product_id = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(ctx, query, ID).
		Scan(
			&product.ProductID,
			&product.Name,
			&product.Sku,
			&product.Category,
			&product.ImageUrl,
			&product.Notes,
			&product.Price,
			&product.Stock,
			&product.Location,
			&product.IsAvailable,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  constant.HTTPStatusText(http.StatusNotFound),
			})
		}

		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}
	return product, nil

}

func (r *Repository) FindProducts(ctx context.Context, params *request.ListProductQuery) ([]*response.ListProduct, error) {
	query := `SELECT product_id, name, sku, category, image_url, stock, notes, price, location, is_available, to_char(created_at, 'YYYY-MM-DD HH24:MI:SS') as created_at
		FROM products where deleted_at IS NULL`

	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	if params.ID != nil {
		clause = append(clause, fmt.Sprintf(" product_id = $%d", counter))
		args = append(args, params.ID)
		counter++
	} else {
		if params.Name != nil {
			clause = append(clause, fmt.Sprintf(" name ilike $%d", counter))
			args = append(args, "%"+*params.Name+"%")
			counter++
		}
		if params.Sku != nil {
			clause = append(clause, fmt.Sprintf(" sku = $%d", counter))
			args = append(args, params.Sku)
			counter++
		}
		if params.Category != nil {
			clause = append(clause, fmt.Sprintf(" category = $%d", counter))
			args = append(args, params.Category)
			counter++
		}
		if params.IsAvailable != nil {
			clause = append(clause, fmt.Sprintf(" is_available = $%d", counter))
			args = append(args, *params.IsAvailable)
			counter++
		}
		if params.InStock != nil {
			if *params.InStock {
				clause = append(clause, fmt.Sprintf(" stock > $%d", counter))
				args = append(args, 0)
				counter++
			} else {
				clause = append(clause, fmt.Sprintf(" stock = $%d", counter))
				args = append(args, 0)
				counter++
			}
		}
	}
	if counter > 1 {
		query += " AND" + strings.Join(clause, " AND")
	}

	query += " ORDER BY"
	if params.Price != nil {
		if *params.Price == "desc" {
			query += " price DESC,"
		}
		if *params.Price == "asc" {
			query += " price ASC,"
		}
	}

	if params.CreatedAt != nil && *params.CreatedAt == "asc" {
		query += " created_at ASC"
	} else {
		query += " created_at DESC"
	}
	if params.Limit != nil && *params.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", counter)
		args = append(args, params.Limit)
		counter++
	} else {
		query += fmt.Sprintf(" LIMIT $%d", counter)
		args = append(args, 5)
		counter++
	}

	if params.Offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", counter)
		args = append(args, params.Offset)
		counter++
	} else {
		query += fmt.Sprintf(" OFFSET $%d", counter)
		args = append(args, 0)
		counter++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*response.ListProduct, 0)
	for rows.Next() {
		p := &response.ListProduct{}
		err = rows.Scan(
			&p.ProductID,
			&p.Name,
			&p.Sku,
			&p.Category,
			&p.ImageUrl,
			&p.Stock,
			&p.Notes,
			&p.Price,
			&p.Location,
			&p.IsAvailable,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}

	return res, nil
}

func (r *Repository) UpdateProduct(ctx context.Context, req *request.UpdateProduct) error {
	query := `UPDATE products 
		SET name = @name, 
			sku = @sku,
			category = @category, 
			image_url = @image_url, 
			notes = @notes, 
			price = @price, 
			stock = @stock, 
			location = @location, 
			is_available = @is_available, 
			updated_at = @updated_at 
		WHERE product_id = @product_id`

	args := pgx.NamedArgs{
		"name":         req.Name,
		"sku":          req.Sku,
		"category":     req.Category,
		"image_url":    req.ImageUrl,
		"notes":        req.Notes,
		"price":        req.Price,
		"stock":        req.Stock,
		"location":     req.Location,
		"is_available": req.IsAvailable,
		"updated_at":   time.Now(),
		"product_id":   req.ProductID,
	}
	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		return err
	}
	return nil

}

func (r *Repository) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	query := `UPDATE products SET deleted_at = $1 WHERE product_id = $2`
	_, err := r.db.Exec(ctx, query, time.Now(), productID)
	if err != nil {
		return err
	}

	return nil
}
