package graph

import (
	"database/sql"

	"github.com/kerinin/doser/service/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db            *sql.DB
	atoController *controller.ATO
	awcController *controller.AWC
}

func NewResolver(
	db *sql.DB,
	atoController *controller.ATO,
	awcController *controller.AWC,
) *Resolver {
	return &Resolver{db, atoController, awcController}
}
