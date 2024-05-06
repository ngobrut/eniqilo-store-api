package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	"github.com/ngobrut/eniqlo-store-api/pkg/custom_error"
)

// CreateUser implements IFaceRepository.
func (r *Repository) CreateUser(ctx context.Context, data *model.User) error {
	query := `INSERT INTO users(name, email, password) VALUES (@name, @email, @password) RETURNING user_id, name, email, created_at`
	args := pgx.NamedArgs{
		"name":     data.Name,
		"email":    data.Email,
		"password": data.Password,
	}

	dest := []interface{}{
		&data.UserID,
		&data.Name,
		&data.Email,
		&data.CreatedAt,
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
	var res *model.User

	// todo:

	return res, nil
}

// GetOneUserByEmail implements IFaceRepository.
func (r *Repository) GetOneUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var res *model.User

	// todo:

	return res, nil
}
