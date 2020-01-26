package services

import (
	"github.com/diogoramos/go-microservice/mvc/domain"
	"github.com/diogoramos/go-microservice/mvc/utils"
)

func GetUser(userId uint64) (*domain.User, *utils.ApplicationError){
	return domain.GetUser(userId)
}