package main

import (
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/leftsidebrain/pzn-go-restful-api/app"
	"github.com/leftsidebrain/pzn-go-restful-api/controller"
	"github.com/leftsidebrain/pzn-go-restful-api/helper"
	"github.com/leftsidebrain/pzn-go-restful-api/middleware"
	"github.com/leftsidebrain/pzn-go-restful-api/repository"
	"github.com/leftsidebrain/pzn-go-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
categoryRepository := repository.NewCategoryRepository()
categoryService:= service.NewCategoryService(categoryRepository, db, validate)	
categoryController := controller.NewCategoryController(categoryService)
router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}