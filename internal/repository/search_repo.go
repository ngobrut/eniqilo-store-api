package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ngobrut/eniqilo-store-api/internal/types/request"
	"github.com/ngobrut/eniqilo-store-api/internal/types/response"
)

func (r *Repository) SearchSKU(ctx context.Context, params *request.SearchQuery) ([]*response.SearchSKU, error) {
	query := `SELECT product_id, name, sku, category, image_url, stock, price, location, to_char(created_at, 'YYYY-MM-DD HH24:MI:SS') as created_at
		FROM products where deleted_at IS NULL AND is_available = true`

	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

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

	if counter > 1 {
		query += " AND" + strings.Join(clause, " AND")
	}

	query += " ORDER BY"
	if params.Price != nil {
		if *params.Price == "desc" {
			query += " price DESC"
		} else if *params.Price == "asc" {
			query += " price ASC"
		} else {
			query += " created_at at time zone 'Asia/Jakarta' DESC"
		}
	} else {
		query += " created_at at time zone 'Asia/Jakarta' DESC"
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
	fmt.Println(query)
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*response.SearchSKU, 0)
	for rows.Next() {
		p := &response.SearchSKU{}
		err = rows.Scan(
			&p.ProductID,
			&p.Name,
			&p.Sku,
			&p.Category,
			&p.ImageUrl,
			&p.Stock,
			&p.Price,
			&p.Location,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}

	return res, nil
}
