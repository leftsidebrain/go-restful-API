package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/leftsidebrain/pzn-go-restful-api/helper"
	"github.com/leftsidebrain/pzn-go-restful-api/model/web"
	"github.com/leftsidebrain/pzn-go-restful-api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	CategoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFormRequestBody(request, &CategoryCreateRequest)

	categoryResponses:=controller.CategoryService.Create(request.Context(), CategoryCreateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	id,err:=strconv.Atoi( params.ByName("categoryId") )
	helper.PanicIfError(err)


	controller.CategoryService.Delete(request.Context(),id) 

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	

	CategoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFormRequestBody(request, &CategoryUpdateRequest)

	
	id,err:=strconv.Atoi( params.ByName("categoryId") )
	helper.PanicIfError(err)

	CategoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(),CategoryUpdateRequest )

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	categoryResponses:=controller.CategoryService.FindAll(request.Context()) 

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id,err:=strconv.Atoi( params.ByName("categoryId") )
	helper.PanicIfError(err)


	categoryResponse:=controller.CategoryService.FindById(request.Context(),id) 

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	
	helper.WriteToResponseBody(writer, webResponse)
}