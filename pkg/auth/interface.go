package service

import "github.com/castlele/lalasync/pkg/auth/models"

type Service interface {
	Register(user *models.UserLogin) error
	Login(user *models.UserLogin) error
}
