package dentist

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

type IDentistRepository interface {
	GetById(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(dentist domain.Dentist, id int) (domain.Dentist, error)
	Delete(id int) error
	Create(dentist domain.Dentist) (domain.Dentist, error)
}

type dentistrepository struct {
	storage store.StoreInterface
}

func NewRepository(storage store.StoreInterface) IDentistRepository {
	return &dentistrepository{storage}
}

func (r *dentistrepository) GetAll() ([]domain.Dentist, error) {

	dentists, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

func (r *dentistrepository) GetById(id int) (domain.Dentist, error) {

	return domain.Dentist{}, nil
}

func (r *dentistrepository) Update(dentist domain.Dentist, id int) (domain.Dentist, error) {
	return domain.Dentist{}, nil
}

func (r *dentistrepository) Delete(id int) error {
	return nil

}
func (r *dentistrepository) Create(dentist domain.Dentist) (domain.Dentist, error) {

	_, err := r.storage.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil

}
