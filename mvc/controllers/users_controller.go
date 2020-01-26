package controllers

import (
	"encoding/json"
	"github.com/diogoramos/go-microservice/mvc/services"
	"github.com/diogoramos/go-microservice/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userId, err :=strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {

		apiErr := &utils.ApplicationError{
			Message: "User_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.Status)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.GetUser(uint64(userId));

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.Status)
		resp.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}