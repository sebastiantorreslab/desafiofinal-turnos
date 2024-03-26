package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
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
	return domain.Dentist{}, nil

}
func (s *sqlStore) Update(dentist domain.Dentist, id int) error {
	return nil

}
func (s *sqlStore) UpdateByField(id int, field string) error {
	return nil

}
func (s *sqlStore) Delete(id int) error {
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

func (s *sqlStore) Exists(codeValue string) bool {
	return false

}
