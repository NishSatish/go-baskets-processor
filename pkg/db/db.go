package db

import (
	"database/sql"
	"go-basket-processor/pkg/db/models/Basket"
)

// This struct is supposed to instantiate and provide the dbObj of all entities
type dbService struct {
	Basket.IBasketOps
}

// I know this is redundant but it's needed simply so you can have an exportable interface
// And directly adding this interfacein the struct directly implements it in the struct
type DBServiceInterface interface {
	Basket.IBasketOps
}

func NewDBService(db *sql.DB) DBServiceInterface {
	return &dbService{
		Basket.NewBasketRepository(db),
	}
}
