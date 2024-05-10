package repository

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/eniqilo-store-api/internal/model"
	"github.com/ngobrut/eniqilo-store-api/internal/types/request"
	"github.com/ngobrut/eniqilo-store-api/internal/types/response"
	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
	"github.com/ngobrut/eniqilo-store-api/pkg/custom_error"
)

func (r *Repository) CheckProductChekoutByIDs(ctx context.Context, IDs []string) (map[string]*model.Product, error) {
	query := `SELECT product_id, price, stock, is_available FROM products
		WHERE deleted_at IS NULL AND product_id = ANY($1)`

	rows, err := r.db.Query(ctx, query, IDs)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	products := make(map[string]*model.Product)
	cnt := 0
	for rows.Next() {
		p := &model.Product{}
		err = rows.Scan(
			&p.ProductID,
			&p.Price,
			&p.Stock,
			&p.IsAvailable,
		)
		if err != nil {
			return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusInternalServerError,
				Message:  constant.HTTPStatusText(http.StatusInternalServerError),
			})
		}
		if !p.IsAvailable {
			return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusBadRequest,
				Message:  "one of productIds isAvailable == false",
			})
		}

		cnt++
		products[p.ProductID.String()] = p
	}
	if cnt != len(IDs) {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusNotFound,
			Message:  "one of productIds is not found",
		})
	}

	return products, nil
}

func (r *Repository) CreateInvoice(ctx context.Context, invoice *model.Invoice, invoiceProducts []*model.InvoiceProduct) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		fmt.Println("error1: " + err.Error())
		return err
	}
	// Setup a deferred function to handle transaction rollback
	defer func() {
		// Check the current error, and if not nil, attempt to rollback
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				fmt.Println("failed to rollback transaction:", rbErr)
			}
		}
	}()

	// insert into table invoices
	invoiceQuery := `INSERT INTO invoices(customer_id, total_price, paid, change)
        VALUES (@customer_id, @total_price, @paid, @change) returning invoice_id`
	args := pgx.NamedArgs{
		"customer_id": invoice.CustomerID,
		"total_price": invoice.TotalPrice,
		"paid":        invoice.Paid,
		"change":      invoice.Change,
	}

	err = tx.QueryRow(ctx, invoiceQuery, args).Scan(&invoice.InvoiceID)
	if err != nil {
		fmt.Println("error2: " + err.Error())
		return err
	}

	// insert into table invoice_products
	// update stock in table products
	insertQuery := `INSERT INTO invoice_products(invoice_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`
	updateQuery := `UPDATE products SET stock = stock - $1, updated_at = $2 WHERE product_id = $3`
	for _, ip := range invoiceProducts {
		_, err = tx.Exec(ctx, insertQuery, invoice.InvoiceID, ip.ProductID, ip.Quantity, ip.Price)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, updateQuery, ip.Quantity, time.Now(), ip.ProductID)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(ctx); err != nil {
		fmt.Printf("failed to commit transaction: %s", err.Error())
		return err
	}
	return nil
}

func (r *Repository) FindInvoices(ctx context.Context, params *request.ListInvoiceQuery) ([]*response.ListInvoice, error) {
	query := `SELECT
			i.invoice_id,
			i.customer_id,
			ip.product_id,
			ip.quantity,
			i.paid,
			i.change	
		FROM invoice_products ip
		LEFT JOIN invoices i ON ip.invoice_id = i.invoice_id `

	subQuery := ` SELECT i.invoice_id FROM invoices i`
	var args = make([]interface{}, 0)
	var counter int = 1
	if params.CustomerID != nil {
		subQuery += fmt.Sprintf(" WHERE i.customer_id = $%d", counter)
		args = append(args, params.CustomerID)
		counter++
	}

	queryOrderby := " ORDER BY i.created_at DESC"
	if params.CreatedAt != nil && *params.CreatedAt == "asc" {
		queryOrderby = " ORDER BY i.created_at ASC"
	}
	subQuery += queryOrderby

	if params.Limit != nil && *params.Limit != 0 {
		subQuery += fmt.Sprintf(" LIMIT $%d", counter)
		args = append(args, params.Limit)
		counter++
	} else {
		subQuery += fmt.Sprintf(" LIMIT $%d", counter)
		args = append(args, 5)
		counter++
	}

	if params.Offset != nil {
		subQuery += fmt.Sprintf(" OFFSET $%d", counter)
		args = append(args, params.Offset)
		counter++
	} else {
		subQuery += fmt.Sprintf(" OFFSET $%d", counter)
		args = append(args, 0)
		counter++
	}

	query += " WHERE i.invoice_id IN (" + subQuery + " )" + queryOrderby

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*response.ListInvoice, 0)

	var isFirst = true
	productDetails := make([]response.ProductDetail, 0)
	prevIN := &response.ListInvoice{}
	for rows.Next() {
		pd := &response.ProductDetail{}
		currIN := &response.ListInvoice{}
		err = rows.Scan(
			&currIN.InvoiceID,
			&currIN.CustomerID,
			&pd.ProductID,
			&pd.Quantity,
			&currIN.Paid,
			&currIN.Change,
		)
		if err != nil {
			return nil, err
		}

		if isFirst {
			prevIN = currIN
			isFirst = false
		}

		if prevIN.InvoiceID != currIN.InvoiceID {
			prevIN.ProductDetail = productDetails
			res = append(res, prevIN)
			productDetails = make([]response.ProductDetail, 0)
			prevIN = currIN
		}
		productDetails = append(productDetails, *pd)

	}
	prevIN.ProductDetail = productDetails
	res = append(res, prevIN)
	if prevIN.InvoiceID.String() == "00000000-0000-0000-0000-000000000000" {
		res = make([]*response.ListInvoice, 0)
	}

	return res, nil

}
