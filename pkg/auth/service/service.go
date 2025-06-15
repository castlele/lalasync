package service

import (
	"errors"

	"github.com/castlele/lalasync/pkg/auth/models"
	"github.com/castlele/lalasync/pkg/storage"
)

const (
	UserRegisteredError = "UserRegisteredError"
	UserEmptyLoginError = "UserEmptyLoginError"
	UserEmptyPasswordError = "UserEmptyPasswordError"

	UnknownUserError = "UnknownUserError"
	InvalidPasswordError = "InvalidPasswordError"
)

type authService struct {
	repo *storage.UserRepo
}

func NewAuthService(repo *storage.UserRepo) *authService {
	return &authService{repo: repo}
}

func (s *authService) Register(user *models.UserLogin) error {
	if s.repo.GetUserByUserName(user.Name) != nil {
		return errors.New(UserRegisteredError)
	}

	err := validateUser(user)

	if err != nil {
		return err
	}

	token, err := NewToken(user.Name)

	if err != nil {
		return err
	}

	user.JWT = token

	err = s.repo.SetUser(&storage.UserModel{
		Name: user.Name,
		Password: user.Password,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(user *models.UserLogin) error {
	savedUser := s.repo.GetUserByUserName(user.Name)

	if savedUser == nil {
		return errors.New(UnknownUserError)
	}

	if savedUser.Password != user.Password {
		return errors.New(InvalidPasswordError)
	}

	token, err := NewToken(user.Name)

	if err != nil {
		return err
	}

	user.JWT = token

	return nil
}

func validateUser(user *models.UserLogin) error {
	if len(user.Name) == 0 {
		return errors.New(UserEmptyLoginError)
	}

	if len(user.Password) == 0 {
		return errors.New(UserEmptyPasswordError)
	}

	return nil
}
