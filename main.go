package main

import (
	"fmt"
	"net/http"

	"github.com/dihanto/go-mastermind/app"
	"github.com/dihanto/go-mastermind/controller"
	"github.com/dihanto/go-mastermind/exception"
	"github.com/dihanto/go-mastermind/helper"
	"github.com/dihanto/go-mastermind/middleware"
	"github.com/dihanto/go-mastermind/repository"
	"github.com/dihanto/go-mastermind/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	repository := repository.NewCartRepository()
	service := service.NewCartService(repository, db, validate)
	controller := controller.NewCartContnroller(service)
	router := app.NewRouter(controller)
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware.NewCartMiddleware(router),
	}
	fmt.Println("server running")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
