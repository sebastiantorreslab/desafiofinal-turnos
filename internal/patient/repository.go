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

func (r *patientRepository) GetAll() ([]domain.Patient, error) {

	patients, err := r.storage.GetAllPatient()
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *patientRepository) GetById(id int) (domain.Patient, error) {

	patient, err := r.storage.GetByIdPatient(id)
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}
