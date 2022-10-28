package services

import "techTask/internal/store"

type DataWebService struct {
	store *store.Store
}

func NewDataWebService(store *store.Store) *DataWebService {
	return &DataWebService{
		store: store,
	}
}
