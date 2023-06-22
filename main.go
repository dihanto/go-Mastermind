package main

import (
	"fmt"
	"net/http"

	"github.com/dihanto/go-mastermind/config"
	"github.com/dihanto/go-mastermind/controller"
	"github.com/dihanto/go-mastermind/exception"
	"github.com/dihanto/go-mastermind/helper"
	"github.com/dihanto/go-mastermind/repository"
	"github.com/dihanto/go-mastermind/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.NewDb()
	validate := validator.New()
	var timeout int
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	repository := repository.NewCustomerRepositoryImpl()
	usecase := usecase.NewCustomerUsecaseImpl(repository, db, validate, timeout)
	controller.NewCustomerControllerImpl(usecase, router)

	server := http.Server{
		Addr:    "localhost:2000",
		Handler: router,
	}
	fmt.Println("server running")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
