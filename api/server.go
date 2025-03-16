package api

import (
	_ "github.com/lib/pq"
	"go-basket-processor/pkg/db"
	"net/http"
)

// StartServer Start the func in caps if you want to export it
func StartServer() {
	startHttpServer()
}

// startHttpServer for some reason you need to write the name of the func in the comment
// don't start it with caps if you don't want to export it
func startHttpServer() {
	db.Connect()

	server := &http.Server{
		Addr:    ":8080",
		Handler: getRouter(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	} else {
		println("Server started")
	}
}
