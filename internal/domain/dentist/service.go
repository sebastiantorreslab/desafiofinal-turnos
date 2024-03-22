package dentist

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type IDentistService interface {
	GetById(id int) (*domain.Dentist, error)
	GetAll() (*[]domain.Dentist, error)
	Update(*domain.Dentist) error
	UpdateByField(id int, field string) error
	Delete(id int) error
	Create(*domain.Dentist) error
}

type DentistServiceImpl struct {
	Repository IDentistRepository
}

func (s *DentistServiceImpl) GetAll() (*[]domain.Dentist, error) {
	dentists, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}
