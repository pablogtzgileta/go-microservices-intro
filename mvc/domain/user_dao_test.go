package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:

	// Execution
	user, err := GetUser(0)

	// Validation
	assert.Nil(t, user, "Not expecting a user with id 0")
	assert.NotNil(t, err, "We ere expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUser(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Pablo", user.FirstName)
	assert.EqualValues(t, "Gutierrez", user.LastName)
	assert.EqualValues(t, "pablo@gmail.com", user.Email)
}