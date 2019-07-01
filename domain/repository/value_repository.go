package repository

import "github.com/kazufusa/go-clean-architecture/domain/model"

type StoreRepository interface {
	Save(*model.Value) error
	Load() (*model.Value, error)
}
