package sprinter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnResponseError(t *testing.T) {
	err := errors.New("error")
	message := "error"
	response := OnResponseError(InternalServerError, err, message)
	assert.Equal(t, err, response.Err)
	assert.Equal(t, message, response.Text)
	assert.Equal(t, 500, response.Code)

	response = OnResponseError(NotFoundError, err, message)
	assert.Equal(t, err, response.Err)
	assert.Equal(t, message, response.Text)
	assert.Equal(t, 404, response.Code)

	response = OnResponseError(BadRequestErr, err, message)
	assert.Equal(t, err, response.Err)
	assert.Equal(t, message, response.Text)
	assert.Equal(t, 400, response.Code)

	response = OnResponseError(ForbiddenErr, err, message)
	assert.Equal(t, err, response.Err)
	assert.Equal(t, message, response.Text)
	assert.Equal(t, 403, response.Code)

	response = OnResponseError(UnauthorizedErr, err, message)
	assert.Equal(t, err, response.Err)
	assert.Equal(t, message, response.Text)
	assert.Equal(t, 401, response.Code)
}
