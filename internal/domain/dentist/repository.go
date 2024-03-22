package dentist

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

type IDentistRepository interface {
	GetById(id int) (*domain.Dentist, error)
	GetAll() (*[]domain.Dentist, error)
	Update(*domain.Dentist) error
	UpdateByField(id int, field string) error
	Delete(id int) error
	Create(*domain.Dentist) error
}

type Repository struct {
	Store store.StoreInterface
}

func (r *Repository) getAll() (*[]domain.Dentist, error) {

	dentists, err := r.Store.GetAll()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}
