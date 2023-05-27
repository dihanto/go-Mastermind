package main

import (
	"fmt"
	"net/http"

	"github.com/dihanto/go-mastermind/app"
	"github.com/dihanto/go-mastermind/controller"
	"github.com/dihanto/go-mastermind/helper"
	"github.com/dihanto/go-mastermind/repository"
	"github.com/dihanto/go-mastermind/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDb()
	repository := repository.NewCartRepository()
	service := service.NewCartService(repository, db)
	controller := controller.NewCartContnroller(service)
	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
	fmt.Println("server running")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
