package middleware

import (
	"net/http"

	"github.com/leftsidebrain/pzn-go-restful-api/helper"
	"github.com/leftsidebrain/pzn-go-restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}


func NewAuthMiddleware(handler http.Handler) *AuthMiddleware{
return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if "Token" == r.Header.Get("AUTH"){

		middleware.Handler.ServeHTTP(w,r)
	}else{
		w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code:http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}
	helper.WriteToResponseBody(w,webResponse)
	}
}