package store

import "github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"

type StoreInterface interface {
	GetById(id int) (*domain.Dentist, error)
	GetAll() (*[]domain.Dentist, error)
	Update(*domain.Dentist) error
	UpdateByField(id int, field string) error
	Delete(id int) error
	Create(*domain.Dentist) error
}
