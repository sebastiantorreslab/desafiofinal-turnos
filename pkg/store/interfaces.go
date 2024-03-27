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

type StoreShiftInterface interface {
	GetById(id int) (domain.Shift, error)
	GetAll() ([]domain.Shift, error)
	Update(shift domain.Shift, id int) error
	Delete(id int) error
	Create(shift domain.Shift) (domain.Shift, error)
}
