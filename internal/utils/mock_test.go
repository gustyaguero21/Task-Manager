package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMock_Ok(t *testing.T) {
	//given
	path := "../mocks/user_mocks/create_user.json"

	//act
	mocks := OpenMock(path)

	//assert
	assert.NotNil(t, mocks)
}
func TestMock_Error(t *testing.T) {
	//given
	path := "../mocks/user_mocks/error.json"

	//act
	mocks := OpenMock(path)

	//assert
	assert.Nil(t, mocks)
}
