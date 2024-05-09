package repository

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_error"
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
