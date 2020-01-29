package controllers

import (
	"github.com/diogoramos/go-microservice/mvc/services"
	"github.com/diogoramos/go-microservice/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context){
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err != nil {

		apiErr := &utils.ApplicationError{
			Message: "User_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		c.JSON(apiErr.Status, apiErr)
		//utils.Respond(c, apiErr.Status, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(uint64(userId));

	if apiErr != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusOK, user)
	return
}