package domain

import (
	"fmt"
	"github.com/diogoramos/go-microservice/mvc/utils"
	"net/http"
)

var (
	users = map[uint64] *User{
		123: {
			Id:        123,
			FirstName: "Diogo",
			LastName:  "Ramos",
			Email:     "diogo.ramos@live.com",
		},
	}
	UserDao userDaoInterface
)

func init(){
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(uint64)(*User, *utils.ApplicationError)
}

type userDao struct {}


func (u *userDao) GetUser(userId uint64) (*User, *utils.ApplicationError){
	if user := users[userId]; user != nil{
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User %d not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
