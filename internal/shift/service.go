package shift

import (
	"log"

	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type IShiftService interface {
	GetById(id int) (domain.Shift, error)
	GetAll() ([]domain.Shift, error)
	Update(shift domain.Shift, id int) error
	UpdateByField(shift domain.Shift, id int) error
	Delete(id int) error
	Create(shift domain.Shift) (domain.Shift, error)
}

type shiftService struct {
	r IShiftRepository
}

func NewShiftService(r IShiftRepository) IShiftService {
	return &shiftService{r}
}

func (s *shiftService) GetById(id int) (domain.Shift, error) {

	shift, err := s.r.GetById(id)
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil

}
func (s *shiftService) GetAll() ([]domain.Shift, error) {

	shifts, err := s.r.GetAll()
	if err != nil {
		return []domain.Shift{}, err
	}
	return shifts, nil
}
func (s *shiftService) Update(shift domain.Shift, id int) error {

	currentShift, err := s.r.GetById(id)
	if err != nil {
		log.Fatal(err)
	}

	if shift.ShiftDate != "" {
		currentShift.ShiftDate = shift.ShiftDate

	}
	if shift.ShiftHour != "" {
		currentShift.ShiftHour = shift.ShiftHour

	}
	if shift.IdPatient != 0 {
		currentShift.IdPatient = shift.IdPatient

	}
	if shift.IdDentist != 0 {
		currentShift.IdDentist = shift.IdDentist

	}

	err = s.r.Update(currentShift, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *shiftService) UpdateByField(shift domain.Shift, id int) error {

	currentShift, err := s.r.GetById(id)
	if err != nil {
		log.Fatal(err)
	}

	if shift.ShiftDate != "" {
		currentShift.ShiftDate = shift.ShiftDate

	}
	if shift.ShiftHour != "" {
		currentShift.ShiftHour = shift.ShiftHour

	}
	if shift.IdPatient != 0 {
		currentShift.IdPatient = shift.IdPatient

	}
	if shift.IdDentist != 0 {
		currentShift.IdDentist = shift.IdDentist

	}

	err = s.r.Update(currentShift, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *shiftService) Delete(id int) error {

	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil

}

func (s *shiftService) Create(shift domain.Shift) (domain.Shift, error) {

	shift, err := s.r.Create(shift)
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil

}
