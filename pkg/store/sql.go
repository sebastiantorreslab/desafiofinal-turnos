package store

import (
	"database/sql"

	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Read() (*[]domain.Dentist, error) {

	query := "SELECT * FROM dentists"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dentists []domain.Dentist

	for rows.Next() {
		var d domain.Dentist
		err := rows.Scan(&d.ID, &d.License, &d.Name, &d.LastName)
		if err != nil {
			return nil, err

		}
		dentists = append(dentists, d)
	}
	if err := rows.Err(); err != nil {

	}

	return &dentists, nil
}
