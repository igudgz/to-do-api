package service

import (
	"todo-api/model"
	"todo-api/persistence"
	"todo-api/util"
)

type UserService interface {
	ValidateUserCredentials(email, password string) (*model.User, bool)
}

type userService struct {
	persistence persistence.User
}

func (u userService) ValidateUserCredentials(email, password string) (*model.User, bool) {

	user, err := u.persistence.FindByEmail(email)
	if err != nil || util.VerifyPassword(user.Password, password) != nil {
		return nil, false
	}
	return user, true
}

func NewUserService(persistence persistence.User) UserService {
	return &userService{
		persistence: persistence,
	}
}
