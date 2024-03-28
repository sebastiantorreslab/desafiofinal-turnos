package shift

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

type IShiftRepository interface {
	GetById(id int) (domain.Shift, error)
	GetByDNI(dni int) (domain.Shift, error)
	GetAll() ([]domain.Shift, error)
	Update(shift domain.Shift, id int) error
	Delete(id int) error
	Create(shift domain.Shift) (domain.Shift, error)
}

type shiftrepository struct {
	storage store.StoreShiftInterface
}

func NewShiftRepository(storage store.StoreShiftInterface) IShiftRepository {
	return &shiftrepository{storage}
}

func (r *shiftrepository) GetById(id int) (domain.Shift, error) {

	shift, err := r.storage.GetById(id)
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil

}
func (r *shiftrepository) GetAll() ([]domain.Shift, error) {

	shifts, err := r.storage.GetAllShifts()
	if err != nil {
		return []domain.Shift{}, err
	}

	return shifts, nil

}
func (r *shiftrepository) Update(shift domain.Shift, id int) error {

	err := r.storage.Update(shift, id)
	if err != nil {
		return err
	}

	return nil

}
func (r *shiftrepository) Delete(id int) error {

	err := r.storage.Delete(id)
	if err != nil {
		return err
	}

	return nil

}

func (r *shiftrepository) Create(shift domain.Shift) (domain.Shift, error) {

	shift, err := r.storage.Create(shift)
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil

}

func (r *shiftrepository) GetByDNI(dni int) (domain.Shift, error) {

	shift, err := r.storage.GetByDNI(dni)
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil

}
