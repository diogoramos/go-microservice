package services

import (
	"github.com/diogoramos/go-microservice/mvc/domain"
	"github.com/diogoramos/go-microservice/mvc/utils"
)

type userService struct{}

var (UserService userService)

func (u *userService) GetUser(userId uint64) (*domain.User, *utils.ApplicationError){
	return domain.UserDao.GetUser(userId)
}