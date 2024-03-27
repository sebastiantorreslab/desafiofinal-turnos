package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type sqlStorePatient struct {
	db *sql.DB
}

func newSqlStorePatient(db *sql.DB) StorePatientInterface {
	return &sqlStorePatient{
		db: db,
	}
}

func (s *sqlStorePatient) CreatePatient(patient domain.Patient) (domain.Patient, error) {
	query := "INSERT INTO patients (dni, name, last_name, address, admission_date) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Patient{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(patient.DNI, patient.Name, patient.LastName, patient.Address, patient.AdmissionDate)
	if err != nil {
		return domain.Patient{}, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, err
}

func (s *sqlStorePatient) GetByIdPatient(id int) (domain.Patient, error) {

	var patient domain.Patient
	query := "SELECT * FROM patients WHERE id=?;"

	stmt := s.db.QueryRow(query, id)

	err := stmt.Scan(&patient.ID, &patient.DNI, &patient.Name, &patient.LastName, &patient.Address, &patient.AdmissionDate)
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil

}
func (s *sqlStorePatient) GetAllPatient() ([]domain.Patient, error) {

	query := "SELECT * FROM `clinic-db`.patients"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []domain.Patient

	for rows.Next() {
		var d domain.Patient
		err := rows.Scan(&d.ID, &d.DNI, &d.Name, &d.LastName, &d.Address, &d.AdmissionDate)
		if err != nil {
			return nil, err

		}
		patients = append(patients, d)
	}
	if err := rows.Err(); err != nil {

	}

	return patients, nil
}
