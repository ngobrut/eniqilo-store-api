package repository

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngobrut/eniqlo-store-api/config"
)

type Repository struct {
	cnf *config.Config
	db  *pgxpool.Pool
}

func New(cnf *config.Config, db *pgxpool.Pool) IFaceRepository {
	return &Repository{
		cnf: cnf,
		db:  db,
	}
}

func IsDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key")
}

func IsRecordNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
