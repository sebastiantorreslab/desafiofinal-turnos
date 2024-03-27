package patient

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type IPatientService interface {
	GetById(id int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(patient domain.Patient, id int) (domain.Patient, error)
	UpdateByField(patient domain.Patient, id int) error
	Delete(id int) error
	Create(patient domain.Patient) (domain.Patient, error)
}

type patientService struct {
	r IPatientRepository
}

func NewService(r IPatientRepository) IPatientRepository {
	return &patientService{r}
}
