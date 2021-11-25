package users_test

import (
	"context"
	"errors"
	"kemejaku/business/users"
	"kemejaku/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//buat mock yang seolah-olah interface dari database
var userRepoInterfaceMock mocks.UserRepoInterface
var userUseCaseInterface users.UserUseCaseInterface
var userDataDummyLogin, userDataDummyEdit users.User
var userDataDummyGetAllUsers []users.User

func setup() {
	userUseCaseInterface = users.NewUseCase(&userRepoInterfaceMock, time.Hour*1)

	//data mock hasil login
	userDataDummyLogin = users.User{
		Id:       1,
		Name:     "Yustina Yasin",
		Email:    "yustinayasin@gmail.com",
		Password: "1234",
		Token:    "",
	}

	userDataDummyGetAllUsers = []users.User{
		{
			Id:       1,
			Name:     "Yustina Yasin",
			Email:    "yustinayasin@gmail.com",
			Password: "1234",
			Token:    "",
		},
		{
			Id:       1,
			Name:     "Jeong Jaehyun",
			Email:    "jeongjaehyun@gmail.com",
			Password: "1234",
			Token:    "",
		},
	}

	userDataDummyEdit = users.User{
		Id:          1,
		Name:        "Yustina Yasin",
		Email:       "yustinayasin@gmail.com",
		Password:    "1234",
		PhoneNumber: "085723252648",
		Street:      "Jl Magelang 19",
		Address:     "Kota Magelang",
		PostalCode:  "56191",
	}
}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Success Login", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.User"), mock.Anything).Return(userDataDummyLogin, nil).Once()

		var requestLoginUser = users.User{
			Email:    "yustinayasin@gmail.com",
			Password: "1234",
		}

		user, err := userUseCaseInterface.Login(requestLoginUser, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, user)
	})

	t.Run("User not found in database", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.User"), mock.Anything).Return(users.User{}, errors.New("User not found")).Once()

		var requestLoginUser = users.User{
			Email:    "Alterra2@gmail.com",
			Password: "123",
		}
		user, err := userUseCaseInterface.Login(requestLoginUser, context.Background())

		assert.Equal(t, errors.New("User not found"), err)
		assert.Equal(t, users.User{}, user)
	})
}

func TestGetAllUsers(t *testing.T) {
	setup()
	t.Run("Success Get All Users", func(t *testing.T) {
		userRepoInterfaceMock.On("GetAllUsers", mock.Anything, mock.Anything).Return(userDataDummyGetAllUsers, nil).Once()

		user, err := userUseCaseInterface.GetAllUsers(context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyGetAllUsers, user)
	})

	t.Run("Users not found in database", func(t *testing.T) {
		userRepoInterfaceMock.On("GetAllUsers", mock.Anything, mock.Anything).Return([]users.User{}, errors.New("There is no user column"))

		user, err := userUseCaseInterface.GetAllUsers(context.Background())

		assert.Equal(t, errors.New("There is no user column"), err)
		assert.Equal(t, []users.User{}, user)
	})
}

func TestSignUp(t *testing.T) {
	setup()
	t.Run("Success Signup", func(t *testing.T) {
		userRepoInterfaceMock.On("SignUp", mock.AnythingOfType("users.User"), mock.Anything).Return(userDataDummyLogin, nil).Once()

		var requestSignUpUser = users.User{
			Email:    "yustinayasin@gmail.com",
			Password: "1234",
		}

		user, err := userUseCaseInterface.SignUp(requestSignUpUser, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, user)
	})

	t.Run("User not found in database", func(t *testing.T) {
		userRepoInterfaceMock.On("SignUp", mock.AnythingOfType("users.User"), mock.Anything).Return(users.User{}, errors.New("Email already exist")).Once()

		var requestSignUpUser = users.User{
			Email:    "yustinayasin@gmail.com",
			Password: "123",
		}
		user, err := userUseCaseInterface.SignUp(requestSignUpUser, context.Background())

		assert.Equal(t, errors.New("Email already exist"), err)
		assert.Equal(t, users.User{}, user)
	})
}

func TestGetUserDetail(t *testing.T) {
	setup()
	t.Run("Success Get User Detail", func(t *testing.T) {
		userRepoInterfaceMock.On("GetUserDetail", mock.Anything, mock.Anything).Return(userDataDummyLogin, nil).Once()

		user, err := userUseCaseInterface.GetUserDetail(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, user)
	})

	t.Run("Users not found in database", func(t *testing.T) {
		userRepoInterfaceMock.On("GetUserDetail", mock.Anything, mock.Anything).Return(users.User{}, errors.New("User not found")).Once()

		user, err := userUseCaseInterface.GetUserDetail(-1, context.Background())

		assert.Equal(t, errors.New("User not found"), err)
		assert.Equal(t, users.User{}, user)
	})
}

func TestEditUser(t *testing.T) {
	setup()
	t.Run("Success Edit", func(t *testing.T) {
		userRepoInterfaceMock.On("EditUser", mock.AnythingOfType("users.User"), mock.Anything, mock.Anything).Return(userDataDummyEdit, nil).Once()

		var requestEditUser = users.User{
			Name:  "Yustina Yasin",
			Email: "yustinayasin@gmail.com",
		}

		user, err := userUseCaseInterface.EditUser(requestEditUser, 1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyEdit, user)
	})

	t.Run("User not found", func(t *testing.T) {
		userRepoInterfaceMock.On("EditUser", mock.AnythingOfType("users.User"), mock.Anything, mock.Anything).Return(users.User{}, errors.New("User not found")).Once()

		var requestEditUser = users.User{
			Email:    "yustinayasin@gmail.com",
			Password: "123",
		}
		user, err := userUseCaseInterface.EditUser(requestEditUser, 1, context.Background())

		assert.Equal(t, errors.New("User not found"), err)
		assert.Equal(t, users.User{}, user)
	})
}

func TestDeleteUser(t *testing.T) {
	setup()
	t.Run("Success delete", func(t *testing.T) {
		userRepoInterfaceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(userDataDummyLogin, nil).Once()

		user, err := userUseCaseInterface.DeleteUser(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, user)
	})

	t.Run("Users not found", func(t *testing.T) {
		userRepoInterfaceMock.On("DeleteUser", mock.Anything, mock.Anything).Return(users.User{}, errors.New("User not found")).Once()

		user, err := userUseCaseInterface.DeleteUser(-1, context.Background())

		assert.Equal(t, errors.New("User not found"), err)
		assert.Equal(t, users.User{}, user)
	})
}
