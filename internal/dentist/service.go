package dentist

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type IDentistService interface {
	GetById(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(dentist domain.Dentist, id int) (domain.Dentist, error)
	UpdateByField(dentist domain.Dentist, id int) error
	Delete(id int) error
	Create(dentist domain.Dentist) (domain.Dentist, error)
}

type denstistService struct {
	r IDentistRepository
}

func NewService(r IDentistRepository) IDentistService {
	return &denstistService{r}
}

func (s *denstistService) Create(dentist domain.Dentist) (domain.Dentist, error) {

	d, err := s.r.Create(dentist)
	if err != nil {
		return domain.Dentist{}, nil
	}

	return d, nil

}

func (s *denstistService) GetAll() ([]domain.Dentist, error) {

	dentists, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

func (s *denstistService) GetById(id int) (domain.Dentist, error) {

	d, err := s.r.GetById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s *denstistService) Update(dentist domain.Dentist, id int) (domain.Dentist, error) {

	d, err := s.r.GetById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	if dentist.License != "" {
		d.License = dentist.License
	}
	if dentist.Name != "" {
		d.Name = dentist.Name
	}
	if dentist.LastName != "" {
		d.LastName = dentist.LastName
	}

	return d, nil

}
func (s *denstistService) UpdateByField(dentist domain.Dentist, id int) error {

	d, err := s.r.GetById(id)
	if err != nil {
		return err
	}
	if dentist.License != "" {
		d.License = dentist.License
	}
	if dentist.Name != "" {
		d.Name = dentist.Name
	}
	if dentist.LastName != "" {
		d.LastName = dentist.LastName
	}
	return nil

}
func (s *denstistService) Delete(id int) error {

	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil

}
