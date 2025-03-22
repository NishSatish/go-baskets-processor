package Basket

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-basket-processor/pkg/dto"
)

// To compare this w Java/Oops
// The struct is kind of the class def with the global fields
// The interface is like an abstract class
// We then complete the class structure by going ahead and implementing
// all the method signatures in the interface/abstract class

// Also create a constructor for the struct that kind of instantiates it

// Similarily create for other entites, and register all of them in the DBLayer interface

type basketRepository struct {
	db *sql.DB
}

type IBasketOps interface {
	CreateBasket(ctx *gin.Context, basketDetails dto.BasketDTO)
	UpdateBasket(ctx *gin.Context)
	DeleteBasket(ctx *gin.Context)
}

func NewBasketRepository(db *sql.DB) IBasketOps {
	return &basketRepository{db: db}
}

func (b *basketRepository) CreateBasket(ctx *gin.Context, basketDetails dto.BasketDTO) {
	// $1 and $2 refer to the args passed in the Exec method
	query := `INSERT INTO basket (name, min_price) VALUES ($1, $2)`

	res, err := b.db.Exec(query, basketDetails.Name, basketDetails.MinPrice)
	if err != nil {
		panic("Query failed ")
	}
	fmt.Println("ARE YOU STORING", res)
}

func (b *basketRepository) UpdateBasket(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (b *basketRepository) DeleteBasket(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

//func CreateTable() {
//	db := dbconnection.GetDB()
//	exists := true
//	err :=
//		db.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'basket')")
//	if err != nil {
//		exists = false
//		return
//	}
//
//	if !exists {
//		createTableQuery := `
//	        CREATE TABLE basket (
//	            id SERIAL PRIMARY KEY,
//	            name VARCHAR(255),
//	            min_price FLOAT8
//	        );`
//		_, err := db.Exec(createTableQuery)
//		if err != nil {
//			panic("Query failed " + err.Error())
//		}
//	}
//}
