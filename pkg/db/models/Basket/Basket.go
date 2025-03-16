package Basket

import dbconnection "go-basket-processor/pkg/db"

type Basket struct {
	ID       int
	name     string
	minPrice float64
}

func AddStock(stockId string) {
	return
}

func CreateTable() {
	db := dbconnection.GetDB()
	exists := true
	err :=
		db.QueryRow("SELECT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'basket')")
	if err != nil {
		exists = false
		return
	}

	if !exists {
		createTableQuery := `
	        CREATE TABLE basket (
	            id SERIAL PRIMARY KEY,
	            name VARCHAR(255),
	            min_price FLOAT8
	        );`
		_, err := db.Exec(createTableQuery)
		if err != nil {
			panic("Query failed " + err.Error())
		}
	}
}
