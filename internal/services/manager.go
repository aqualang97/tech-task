package services

import (
	"errors"
	"techTask/internal/store"
)

type Manager struct {
	Data DataService
}

func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("no store provided")
	}
	return &Manager{
		Data: NewDataWebService(store),
	}, nil
}
