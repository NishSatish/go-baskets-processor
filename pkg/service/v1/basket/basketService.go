package basket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db2 "go-basket-processor/pkg/db"
	"go-basket-processor/pkg/dto"
)

// Reg all basket controllers here, and register an instance of this whole
// service in the v1 file. The router will run this whole thing using the v1 service

type basketService struct {
	// Mimics global vars of a java class
	dbObj db2.DBServiceInterface
}

type BasketServiceInterface interface {
	CreateBasket(ctx *gin.Context) // only needs the request context
	UpdateBasket(ctx *gin.Context)
	DeleteBasket(ctx *gin.Context)
	GetBaskets(ctx *gin.Context)
	GetBasketByID(ctx *gin.Context)
}

func NewBasketsService(db db2.DBServiceInterface) BasketServiceInterface {
	return &basketService{
		dbObj: db,
	}
}

func (b basketService) CreateBasket(ctx *gin.Context) {
	var incomingBasket dto.BasketDTO
	err := ctx.BindJSON(&incomingBasket)
	if err != nil {
		fmt.Println("Error in binding")
	}
	b.dbObj.CreateBasket(ctx, incomingBasket)
	fmt.Println("BASKETS CONTROLLER", incomingBasket)
}

func (b basketService) UpdateBasket(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (b basketService) DeleteBasket(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (b basketService) GetBaskets(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (b basketService) GetBasketByID(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
