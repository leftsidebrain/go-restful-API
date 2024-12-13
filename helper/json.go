package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFormRequestBody(r *http.Request, data interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err:=encoder.Encode(response)
	PanicIfError(err)
}