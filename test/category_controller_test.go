package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/leftsidebrain/pzn-go-restful-api/app"
	"github.com/leftsidebrain/pzn-go-restful-api/controller"
	"github.com/leftsidebrain/pzn-go-restful-api/helper"
	"github.com/leftsidebrain/pzn-go-restful-api/middleware"
	"github.com/leftsidebrain/pzn-go-restful-api/repository"
	"github.com/leftsidebrain/pzn-go-restful-api/service"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	db,err:=sql.Open("mysql", "root@tcp(localhost:3306)/pzn_go_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler{
	db := setupTestDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService:= service.NewCategoryService(categoryRepository, db, validate)	
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

return middleware.NewAuthMiddleware(router)
}


func TestCreateCategorySuccess(t *testing.T){
	router := setupRouter()

	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/category", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("AUTH", "Token")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode) 
}

func TestCreateCategoryFailed(t *testing.T){}

func TestUpdateCategorySuccess(t *testing.T){}

func TestUpdateCategoryFailed(t *testing.T){}