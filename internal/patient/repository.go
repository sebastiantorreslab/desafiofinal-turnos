package patient

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

type IPatientRepository interface {
	GetById(id int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(dentist domain.Patient, id int) error
	Delete(id int) error
	Create(dentist domain.Patient) (domain.Patient, error)
}

type patientRepository struct {
	storage store.StorePatientInterface
}

func NewPatientRepository(storage store.StorePatientInterface) IPatientRepository {
	return &patientRepository{storage}
}
