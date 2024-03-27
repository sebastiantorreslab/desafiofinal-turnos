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

func (s *patientService) Create(patient domain.Patient) (domain.Patient, error) {

	d, err := s.r.Create(patient)
	if err != nil {
		return domain.Patient{}, nil
	}

	return d, nil

}

func (s *patientService) GetAll() ([]domain.Patient, error) {

	patients, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (s *patientService) GetById(id int) (domain.Patient, error) {

	d, err := s.r.GetById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return d, nil
}

func (s *patientService) Update(patient domain.Patient, id int) (domain.Patient, error) {

	d, err := s.r.GetById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	if patient.DNI != 0 {
		d.DNI = patient.DNI
	}
	if patient.Name != "" {
		d.Name = patient.Name
	}
	if patient.LastName != "" {
		d.LastName = patient.LastName
	}
	if patient.Address != "" {
		d.Address = patient.Address
	}
	if !patient.AdmissionDate.IsZero() {
		d.AdmissionDate = patient.AdmissionDate
	}

	err = s.r.Update(d, id)
	if err != nil {
		return domain.Patient{}, err
	}

	return d, nil

}
