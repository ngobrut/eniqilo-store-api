package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/eniqilo-store-api/internal/model"
	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
	"github.com/ngobrut/eniqilo-store-api/pkg/custom_error"
)

// CreateUser implements IFaceRepository.
func (r *Repository) CreateUser(ctx context.Context, data *model.User) error {
	query := `INSERT INTO users(name, phone, password) VALUES (@name, @phone, @password) RETURNING user_id, name, phone`
	args := pgx.NamedArgs{
		"name":     data.Name,
		"phone":    data.Phone,
		"password": data.Password,
	}

	dest := []interface{}{
		&data.UserID,
		&data.Name,
		&data.Phone,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(dest...)
	if err != nil {
		if IsDuplicateError(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  constant.HTTPStatusText(http.StatusConflict),
			})
		}

		return err
	}

	return nil
}

// GetOneUserByID implements IFaceRepository.
func (r *Repository) GetOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	res := &model.User{}

	err := r.db.
		QueryRow(ctx, "SELECT * FROM users WHERE user_id = $1", userID).
		Scan(
			&res.UserID,
			&res.Name,
			&res.Phone,
			&res.Password,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  constant.HTTPStatusText(http.StatusNotFound),
			})
		}

		return nil, err
	}

	return res, nil
}

// GetOneUserByEmail implements IFaceRepository.
func (r *Repository) GetOneUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	res := &model.User{}

	err := r.db.
		QueryRow(ctx, "SELECT * FROM users WHERE phone = $1", phone).
		Scan(
			&res.UserID,
			&res.Name,
			&res.Phone,
			&res.Password,
			&res.CreatedAt,
			&res.UpdatedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  constant.HTTPStatusText(http.StatusNotFound),
			})
		}

		return nil, err
	}

	return res, nil
}
