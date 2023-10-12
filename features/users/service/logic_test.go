package service

import (
	"Hannon-app/features/users"
	"Hannon-app/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Buat mock untuk UserDataInterface
	mockUserData := new(mocks.UserData)

	// Inisialisasi UserService dengan mock UserData
	userService := &UserService{
		userData: mockUserData,
	}

	// Data yang diharapkan dari mock
	returnData := users.UserCore{
		ID:           1,
		Name:         "userHannon",
		Email:        "user@hannon.com",
		PhoneNumber:  "085751151123",
		Password:     "password",
		Address:      "malang",
		ProfilePhoto: "default.jpg",
		UploadKTP:    "default.jpg",
		Role:         "users",
		MembershipID: 0,
		Membership:   users.MembershipCore{},
	}

	t.Run("test case success login", func(t *testing.T) {
		email := "user@hannon.com"
		password := "user123"

		// Definisikan argumen dan hasil yang diharapkan dari mock
		mockUserData.On("Login", email, password).Return(returnData, nil).Once()

		// Panggil fungsi yang diuji
		dataLogin, token, err := userService.Login(email, password)

		// Periksa hasil dengan menggunakan assert
		assert.NoError(t, err)
		assert.Equal(t, returnData, dataLogin)
		assert.NotNil(t, token)

		// Periksa bahwa metode mock dipanggil dengan argumen yang benar
		mockUserData.AssertCalled(t, "Login", email, password)
	})

	t.Run("test case failed login", func(t *testing.T) {
		email := "user@hanon.com"
		password := "wrongPassword"

		// Definisikan argumen dan hasil yang diharapkan dari mock
		mockUserData.On("Login", email, password).Return(users.UserCore{}, errors.New("error login")).Once()

		// Panggil fungsi yang diuji
		dataLogin, token, err := userService.Login(email, password)

		// Periksa hasil dengan menggunakan assert
		assert.NotNil(t, err)
		assert.Equal(t, users.UserCore{}, dataLogin) // Check for an empty UserCore
		assert.Empty(t, token)
		assert.EqualError(t, err, "error login")

		// Periksa bahwa metode mock dipanggil dengan argumen yang benar
		mockUserData.AssertCalled(t, "Login", email, password)
	})
}

func TestGetAll(t *testing.T) {
	mockUserData := new(mocks.UserData)
	userService := &UserService{
		userData: mockUserData,
	}

	// Data yang diharapkan dari mock
	expectedUsers := []users.UserCore{
		{
			ID:           1,
			Name:         "User1",
			Email:        "user1@example.com",
			PhoneNumber:  "",
			Password:     "user1234",
			Address:      "malang",
			ProfilePhoto: "default.jpg",
			UploadKTP:    "default.jpg",
			Role:         "user",
			MembershipID: 0,
			Membership:   users.MembershipCore{},
		},
		{
			ID:           2,
			Name:         "User2",
			Email:        "user2@example.com",
			PhoneNumber:  "",
			Password:     "user1234",
			Address:      "lumajang",
			ProfilePhoto: "default.jpg",
			UploadKTP:    "default.jpg",
			Role:         "user",
			MembershipID: 0,
			Membership:   users.MembershipCore{},
		},
	}
	expectedCount := int64(len(expectedUsers))

	t.Run("test case success GetAll", func(t *testing.T) {
		adminID := uint(1)
		page := uint(1)
		userPerPage := uint(10)
		searchName := "searchName"

		// Definisikan argumen dan hasil yang diharapkan dari mock
		mockUserData.On("ReadAll", adminID, page, userPerPage, searchName).Return(expectedUsers, expectedCount, nil).Once()

		// Panggil fungsi yang diuji
		users, hasNext, err := userService.GetAll(adminID, page, userPerPage, searchName)

		// Periksa hasil dengan menggunakan assert
		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)
		assert.True(t, hasNext)

		// Periksa bahwa metode mock dipanggil dengan argumen yang benar
		mockUserData.AssertCalled(t, "ReadAll", adminID, page, userPerPage, searchName)
	})

	t.Run("test case failed GetAll", func(t *testing.T) {
		// Implementasikan pengujian ketika pemanggilan userData.ReadAll mengembalikan error
		adminID := uint(1)
		page := uint(1)
		userPerPage := uint(10)
		searchName := "searchName"

		// Definisikan argumen yang sama dengan yang berhasil, tapi hasil dari mock adalah error
		mockUserData.On("ReadAll", adminID, page, userPerPage, searchName).Return(nil, int64(0), errors.New("error getAll")).Once()

		// Panggil fungsi yang diuji
		users, hasNext, err := userService.GetAll(adminID, page, userPerPage, searchName)

		// Periksa hasil dengan menggunakan assert
		assert.NotNil(t, err)
		assert.Empty(t, users)
		assert.False(t, hasNext)
		assert.EqualError(t, err, "error getAll")

		// Periksa bahwa metode mock dipanggil dengan argumen yang benar
		mockUserData.AssertCalled(t, "ReadAll", adminID, page, userPerPage, searchName)
	})
}

func TestDeleteByID(t *testing.T) {
	mockUserData := new(mocks.UserData)
	userService := &UserService{
		userData: mockUserData,
	}

	t.Run("test case success DeleteByID", func(t *testing.T) {
		adminID := uint(1)
		userID := uint(5)

		// Definisikan argumen yang diharapkan dari mock
		mockUserData.On("Delete", adminID, userID).Return(nil).Once()

		// Panggil fungsi yang diuji
		err := userService.Deletebyid(adminID, userID)

		// Periksa hasil dengan menggunakan assert
		assert.NoError(t, err)

		// Periksa bahwa metode mock dipanggil dengan argumen yang benar
		mockUserData.AssertCalled(t, "Delete", adminID, userID)
	})

	t.Run("test case failed DeleteByID", func(t *testing.T) {
		// Implementasikan pengujian ketika pemanggilan userData.Delete mengembalikan error
		adminID := uint(1)
		userID := uint(10)

		// Definisikan argumen yang sama dengan yang berhasil, tapi hasil dari mock adalah error
		mockUserData.On("Delete", adminID, userID).Return(errors.New("error delete")).Once()

		// Panggil fungsi yang diuji
		err := userService.Deletebyid(adminID, userID)

		// Periksa hasil dengan menggunakan assert
		assert.NotNil(t, err)
		assert.EqualError(t, err, "error delete")

		// Periksa bahwa metode mock dipanggil dengan argumen yang benar
		mockUserData.AssertCalled(t, "Delete", adminID, userID)
	})
}

func TestGetUserById(t *testing.T) {
	mockUserData := new(mocks.UserData)
	userService := &UserService{
		userData: mockUserData,
	}

	t.Run("test case success GetUserById", func(t *testing.T) {
		id := uint(1)
		expectedUser := users.UserCore{
			ID: id,
			/* Isi dengan data yang sesuai */
		}

		// Definisikan argumen yang diharapkan dari mock
		mockUserData.On("SelectById", id).Return(expectedUser, nil).Once()

		// Panggil fungsi yang diuji
		result, err := userService.GetUserById(id)

		// Periksa hasil dengan menggunakan assert
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, result)

		// Periksa bahwa metode mock dipanggil dengan argumen yang benar
		mockUserData.AssertCalled(t, "SelectById", id)
	})

	t.Run("test case error GetUserById", func(t *testing.T) {
		id := uint(2)

		// Definisikan error yang diharapkan dari mock
		expectedError := errors.New("error fetching user")

		// Definisikan argumen yang diharapkan dari mock
		mockUserData.On("SelectById", id).Return(users.UserCore{}, expectedError).Once()

		// Panggil fungsi yang diuji
		result, err := userService.GetUserById(id)

		// Periksa hasil dengan menggunakan assert
		assert.Error(t, err)
		assert.Equal(t, users.UserCore{}, result)
		assert.EqualError(t, err, "error fetching user")

		// Periksa bahwa metode mock dipanggil dengan argumen yang benar
		mockUserData.AssertCalled(t, "SelectById", id)
	})
}
