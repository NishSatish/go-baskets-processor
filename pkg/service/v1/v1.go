package v1

import (
	db "go-basket-processor/pkg/db"
	"go-basket-processor/pkg/service/v1/basket"
)

// Register all controllers here
type v1Service struct {
	basket.BasketServiceInterface
}

type V1ServiceInterface interface {
	basket.BasketServiceInterface
}

func NewV1Service(dbService db.DBServiceInterface) V1ServiceInterface {
	return &v1Service{
		basket.NewBasketsService(dbService),
	}
}
