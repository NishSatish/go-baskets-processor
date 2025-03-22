package db

import (
	"database/sql"
	"fmt"
	"go-basket-processor/pkg/config"
)

func Connect() *sql.DB {
	var (
		HOST     = config.GetConfig().GetString("DB_HOST")
		PORT     = config.GetConfig().GetInt("DB_PORT")
		USER     = config.GetConfig().GetString("DB_USER")
		PASSWORD = config.GetConfig().GetString("DB_PWD")
		DBNAME   = config.GetConfig().GetString("DB_NAME")
	)
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", connString)
	//db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	fmt.Println("PDB Connected")

	return db
}
