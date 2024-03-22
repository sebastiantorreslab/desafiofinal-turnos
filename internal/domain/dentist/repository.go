package dentist

import (
	"github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"
	"github.com/sebastiantorreslab/desafiofinal-turnos/pkg/store"
)

type IRepository interface {
	GetAll()
}

type Repository struct {
	Store store.StoreInterface
}

func (r *Repository) getAll() (*[]domain.Dentist, error) {

	var dentists []domain.Dentist

	return &dentists, nil
}
