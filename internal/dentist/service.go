package dentist

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type IDentistService interface {
	GetById(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(dentist domain.Dentist) error
	UpdateByField(id int, field string) error
	Delete(id int) error
	Create(dentist domain.Dentist) error
	Exists(codeValue string) bool
}

type denstistService struct {
	r IDentistRepository
}

func NewService(r IDentistRepository) IDentistService {
	return &denstistService{r}
}

func (s *denstistService) GetAll() ([]domain.Dentist, error) {

	dentists, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

func (s *denstistService) GetById(id int) (domain.Dentist, error) {
	return domain.Dentist{}, nil
}

func (s *denstistService) Update(dentist domain.Dentist) error {
	return nil
}
func (s *denstistService) UpdateByField(id int, field string) error {
	return nil
}
func (s *denstistService) Delete(id int) error {
	return nil

}
func (s *denstistService) Create(dentist domain.Dentist) error {
	return nil

}
func (s *denstistService) Exists(codeValue string) bool {
	return false

}
