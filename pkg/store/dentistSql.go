package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreDentistInterface {
	return &sqlStore{
		db: db,
	}

}

func ConnectDB() (*sql.DB, error) {

	dataSource := "user-db:pass-db@tcp(localhost:3360)/clinic-db"

	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db, nil
}

func (s *sqlStore) GetAll() ([]domain.Dentist, error) {

	query := "SELECT * FROM `clinic-db`.dentists"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dentists []domain.Dentist

	for rows.Next() {
		var d domain.Dentist
		err := rows.Scan(&d.ID, &d.License, &d.Name, &d.LastName)
		if err != nil {
			return nil, err

		}
		dentists = append(dentists, d)
	}
	if err := rows.Err(); err != nil {

	}

	return dentists, nil
}

func (s *sqlStore) GetById(id int) (domain.Dentist, error) {

	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE id=?;"

	stmt := s.db.QueryRow(query, id)

	err := stmt.Scan(&dentist.ID, &dentist.License, &dentist.Name, &dentist.LastName)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil

}
func (s *sqlStore) Update(dentist domain.Dentist, id int) error {

	query := "UPDATE dentists SET license = ?, name = ? , last_name=? WHERE id=?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(dentist.License, dentist.Name, dentist.LastName, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

func (s *sqlStore) Delete(id int) error {

	query := "DELETE FROM dentists WHERE id = ?;"

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
func (s *sqlStore) Create(dentist domain.Dentist) (domain.Dentist, error) {
	query := "INSERT INTO dentists (license, name, last_name) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domain.Dentist{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(dentist.License, dentist.Name, dentist.LastName)
	if err != nil {
		return domain.Dentist{}, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, err
}
