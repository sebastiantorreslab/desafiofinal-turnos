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
	GetByIdPatient(id int) (domain.Patient, error)
	GetAllPatient() ([]domain.Patient, error)
	UpdatePatient(patient domain.Patient, id int) error
	DeletePatient(id int) error
	CreatePatient(patient domain.Patient) (domain.Patient, error)
}
