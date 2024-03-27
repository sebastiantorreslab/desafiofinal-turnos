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

	





	return domain.Shift{}, nil

}
func (s *sqlShiftStore) GetAll() ([]domain.Shift, error) {
	return []domain.Shift{}, nil

}
func (s *sqlShiftStore) Update(shift domain.Shift, id int) error {
	return nil

}
func (s *sqlShiftStore) Delete(id int) error {
	return nil

}
func (s *sqlShiftStore) Create(shift domain.Shift) (domain.Shift, error) {
	return domain.Shift{}, nil

}
