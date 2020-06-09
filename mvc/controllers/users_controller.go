package controllers

import (
	"encoding/json"
	"github.com/KatherineEbel/go-microservices/mvc/services"
	"github.com/KatherineEbel/go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(response http.ResponseWriter, request *http.Request) {
	idParam := request.URL.Query().Get("id")
	id, apiErr := strconv.ParseInt(idParam, 10, 64)
	if apiErr != nil {
		userErr := &utils.ApplicationError{
			Message:    "id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonVal, _ := json.Marshal(userErr)
		response.WriteHeader(userErr.StatusCode)
		if _, err := response.Write(jsonVal); err != nil {
			return
		}
	}

	user, err := services.GetUser(id)
	if err != nil {
		jsonVal, _ := json.Marshal(err)
		response.WriteHeader(err.StatusCode)
		if _, err := response.Write(jsonVal); err != nil {
			return
		}
	}

	jsonVal, _ := json.Marshal(user)
	if _, err := response.Write(jsonVal); err != nil {
		return
	}
}
