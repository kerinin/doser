package graph

import (
	badger "github.com/dgraph-io/badger/v2"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *badger.DB
}

func NewResolver(db *badger.DB) *Resolver {
	return &Resolver{db}
}
