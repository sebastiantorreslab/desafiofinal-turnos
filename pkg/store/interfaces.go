package store

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type StoreInterface interface {
	GetById(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(dentist domain.Dentist, id int) error
	UpdateByField(id int, field string) error
	Delete(id int) error
	Create(dentist domain.Dentist) (domain.Dentist, error)
}
