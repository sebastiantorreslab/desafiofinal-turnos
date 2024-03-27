package dentist

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

type IDentistRepository interface {
	GetById(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(dentist domain.Dentist, id int) error
	Delete(id int) error
	Create(dentist domain.Dentist) (domain.Dentist, error)
}

type dentistrepository struct {
	storage store.StoreDentistInterface
}

func NewRepository(storage store.StoreDentistInterface) IDentistRepository {
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

	dentist, err := r.storage.GetById(id)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

func (r *dentistrepository) Update(dentist domain.Dentist, id int) error {

	dentist, err := r.storage.GetById(id)
	if err != nil {
		return err
	}
	err = r.storage.Update(dentist, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *dentistrepository) Delete(id int) error {

	err := r.storage.Delete(id)
	if err != nil {
		return err
	}

	return nil

}
func (r *dentistrepository) Create(dentist domain.Dentist) (domain.Dentist, error) {

	_, err := r.storage.Create(dentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil

}
