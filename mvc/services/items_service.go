package services

import (
	"github.com/diogoramos/go-microservice/mvc/domain"
	"github.com/diogoramos/go-microservice/mvc/utils"
	"net/http"
)

type itemService struct {
}

var (ItemService itemService)

func (i *itemService) GetItem(itemId string)(*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "Implement me",
		Status:  http.StatusInternalServerError,
		Code:    "internal_server_error",
	}
}