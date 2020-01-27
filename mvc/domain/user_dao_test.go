package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	//Initialization

	//Execution
	user, err := UserDao.GetUser(0)

	//Validation
	assert.Nil(t, user, "")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Diogo", user.FirstName)
	assert.EqualValues(t, "Ramos", user.LastName)
	assert.EqualValues(t, "diogo.ramos@live.com", user.Email)


}