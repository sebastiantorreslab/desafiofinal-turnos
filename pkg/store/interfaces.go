package store

import "github.com/sebastiantorreslab/desafiofinal-turnos/internal/domain"

type StoreInterface interface {
	Read() (*[]domain.Dentist, error)
}
