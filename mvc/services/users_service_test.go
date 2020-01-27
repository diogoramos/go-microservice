package services

import (
	"github.com/diogoramos/go-microservice/mvc/domain"
	"github.com/diogoramos/go-microservice/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func (userId uint64)(*domain.User, *utils.ApplicationError)
	)

func init(){
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct {

}

func (m *usersDaoMock) GetUser(userId uint64)(*domain.User, *utils.ApplicationError){
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T){
	getUserFunction = func(userId uint64) (*domain.User, *utils.ApplicationError){
		return nil, &utils.ApplicationError{
			Message: "User 0 not found",
			Status:  http.StatusNotFound,
			Code:    "not_found",
		}
	}

	user, err := UserService.GetUser(0)

	assert.Nil(t, user, "")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 not found", err.Message)
}
