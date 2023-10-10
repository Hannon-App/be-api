package service

import (
	"Hannon-app/features/admins"
	"Hannon-app/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	mockAdminDataLayer := new(mocks.AdminData)
	returnData := admins.AdminCore{
		ID:       1,
		Name:     "adminHannon",
		Username: "admin",
		Email:    "admin@hannon.com",
		Password: "admin123",
		Role:     "admin",
	}
	t.Run("test case success login", func(t *testing.T) {
		email := "admin@hannon.com"
		password := "admin123"
		mockAdminDataLayer.On("Login", email, password).Return(returnData, nil).Once()
		srv := New(mockAdminDataLayer)
		dataLogin, token, err := srv.Login(email, password)
		assert.Nil(t, err)
		assert.Equal(t, returnData.ID, dataLogin.ID, returnData.Name, dataLogin.Name, returnData.Role, dataLogin.Role)
		assert.NotNil(t, token)
		mockAdminDataLayer.AssertExpectations(t)

	})
	t.Run("test case failed login", func(t *testing.T) {
		email := "admin@hanon.com"
		password := ""
		mockAdminDataLayer.On("Login", email, password).Return(admins.AdminCore{}, errors.New("error login")).Once()
		srv := New(mockAdminDataLayer)
		dataLogin, token, err := srv.Login(email, password)
		assert.NotNil(t, err)
		assert.Equal(t, admins.AdminCore{}, dataLogin) // Check for an empty AdminCore
		assert.Empty(t, token)
		assert.EqualError(t, err, "error login")
		mockAdminDataLayer.AssertExpectations(t)
	})

}
