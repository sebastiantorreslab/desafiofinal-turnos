package store

import(

	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	
)

type StoreInterface interface {
	GetById(id int) (*domain.Dentist, error)
	GetAll() (*[]domain.Dentist, error)
	Update(*domain.Dentist) error
	UpdateByField(id int, field string) error
	Delete(id int) error
	Create(*domain.Dentist) error
}

func (s *SqlStore) GetAll() (*[]domain.Dentist, error) {

	query := "SELECT * FROM dentists"
	rows, err := s.StorageDB.Query(query)
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
