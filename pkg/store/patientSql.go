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
