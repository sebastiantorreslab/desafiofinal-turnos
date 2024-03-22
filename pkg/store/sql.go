package store

import (
	"database/sql"

	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Read() (*[]domain.Dentist, error) {

	var dentist []domain.Dentist
	return &dentist, nil
}
