package store

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type StoreDentistInterface interface {
	GetById(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(dentist domain.Dentist, id int) error
	Delete(id int) error
	Create(dentist domain.Dentist) (domain.Dentist, error)
}

type StorePatientInterface interface {
	GetById(id int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(patient domain.Patient, id int) error
	Delete(id int) error
	Create(patient domain.Patient) (domain.Patient, error)
}
