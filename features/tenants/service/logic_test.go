package service

import (
	"Hannon-app/features/tenants"
	"Hannon-app/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	mockTenantDataLayer := new(mocks.TenantData)
	returnData := tenants.TenantCore{
		ID:       1,
		Name:     "heiger",
		Email:    "admin@heiger.com",
		Password: "heiger123",
		Role:     "tenant",
	}
	t.Run("test case success login", func(t *testing.T) {
		email := "admin@heiger.com"
		password := "heiger123"
		mockTenantDataLayer.On("Login", email, password).Return(returnData, nil).Once()
		srv := New(mockTenantDataLayer)
		dataLogin, token, err := srv.Login(email, password)
		assert.Equal(t, returnData.ID, dataLogin.ID, returnData.Name, dataLogin.Name, returnData.Role, dataLogin.Role)
		assert.Nil(t, err)
		assert.NotNil(t, token)
		mockTenantDataLayer.AssertExpectations(t)
	})
	t.Run("test case failed login", func(t *testing.T) {
		var email string
		var password string
		mockTenantDataLayer.On("Login", email, password).Return(tenants.TenantCore{}, errors.New("error login")).Once()
		srv := New(mockTenantDataLayer)
		dataLogin, token, err := srv.Login(email, password)
		assert.NotNil(t, err)
		assert.Equal(t, tenants.TenantCore{}, dataLogin)
		assert.Empty(t, token)
		assert.EqualError(t, err, "error login")
		mockTenantDataLayer.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	var id uint
	mockTenantDataLayer := new(mocks.TenantData)
	t.Run("test case success delete", func(t *testing.T) {
		id = 1
		mockTenantDataLayer.On("Delete", id).Return(nil).Once()
		srv := New(mockTenantDataLayer)
		err := srv.Remove(id)
		assert.Nil(t, err)
		mockTenantDataLayer.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockTenantDataLayer := new(mocks.TenantData)
	returnData := []tenants.TenantCore{{
		ID:    1,
		Name:  "sumber berkah outdoor",
		Email: "sumberberkah@gmail.com",
	}}
	t.Run("test case success get all", func(t *testing.T) {
		address := "malang"
		mockTenantDataLayer.On("GetAll", address).Return(returnData, nil).Once()
		srv := New(mockTenantDataLayer)
		result, err := srv.ReadAll(address)
		assert.Equal(t, returnData[0].ID, result[0].ID)
		assert.Nil(t, err)
		mockTenantDataLayer.AssertExpectations(t)
	})
}

func TestGetTenantById(t *testing.T) {
	mockTenantDataLayer := new(mocks.TenantData)
	returnData := tenants.TenantCore{
		ID:    1,
		Name:  "sumber berkah outdoor",
		Email: "sumberberkah@gmail.com",
	}
	t.Run("test case succes get by id", func(t *testing.T) {
		id := uint(1)
		mockTenantDataLayer.On("GetTenantById", id).Return(returnData, nil).Once()
		srv := New(mockTenantDataLayer)
		result, err := srv.ReadTenantById(id)
		assert.Nil(t, err)
		assert.Equal(t, returnData.ID, result.ID)
		assert.Equal(t, returnData.Name, result.Name)
		mockTenantDataLayer.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	mockTenantDataLayer := new(mocks.TenantData)
	t.Run("test case success register", func(t *testing.T) {
		var image multipart.File
		var IDCard multipart.File
		inputData := tenants.TenantCore{
			Name:      "Sumber Berkah Outdoor",
			Email:     "sumberberkah@outdoor.com",
			Password:  "berkah123",
			Phone:     "085549663213",
			Images:    "logo.png",
			Address:   "Lumajang",
			IDcard:    "tenantID.png",
			OpenTime:  "08.00",
			CloseTime: "22.00",
		}
		mockTenantDataLayer.On("Register", inputData, image, IDCard, "logo", "IDcard").Return(nil).Once()
		srv := New(mockTenantDataLayer)
		err := srv.Create(inputData, image, IDCard, "logo", "IDcard")
		assert.Nil(t, err)
		mockTenantDataLayer.AssertExpectations(t)
	})
}
