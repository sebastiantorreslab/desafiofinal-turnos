package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type sqlShiftStore struct {
	db *sql.DB
}

func NewSqlShiftStore(db *sql.DB) StoreShiftInterface {
	return &sqlShiftStore{
		db: db,
	}
}

func (s *sqlShiftStore) GetById(id int) (domain.Shift, error) {

	var shift domain.Shift
	query := "SELECT * FROM shift WHERE id=?;"

	stmt := s.db.QueryRow(query, id)

	err := stmt.Scan(&shift.ID, &shift.Patient, &shift.Dentist, &shift.ShiftHour, &shift.ShiftDate)
	if err != nil {
		return domain.Shift{}, err
	}

	return domain.Shift{}, nil

}
func (s *sqlShiftStore) GetAll() ([]domain.Shift, error) {

	query := "SELECT * FROM `clinic-db`.shift"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shifts []domain.Shift

	for rows.Next() {
		var shift domain.Shift
		err := rows.Scan(&shift.ID, &shift.Patient, &shift.Dentist, &shift.ShiftHour, &shift.ShiftDate)
		if err != nil {
			return nil, err

		}
		shifts = append(shifts, shift)
	}
	if err := rows.Err(); err != nil {
		return []domain.Shift{}, err
	}

	return shifts, nil

}
func (s *sqlShiftStore) Update(shift domain.Shift, id int) error {

	query := "UPDATE shift SET shift_date = ?, shift_hour = ?, id_patient=?,id_dentist=? WHERE id=?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(shift.ShiftDate, shift.ShiftHour, shift.Patient.ID, shift.Dentist.ID, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}
func (s *sqlShiftStore) Delete(id int) error {

	query := "DELETE FROM shift WHERE id = ?;"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}
func (s *sqlShiftStore) Create(shift domain.Shift) (domain.Shift, error) {

	query := "INSERT INTO shift (shift_date, shift_hour, id_patient,id_dentist) VALUES (?, ?, ?,?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Shift{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(shift.ShiftDate, shift.ShiftHour, shift.Patient.ID, shift.Dentist.ID)
	if err != nil {
		return domain.Shift{}, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil

}
