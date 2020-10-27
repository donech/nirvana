package service

import (
	"github.com/donech/nirvana/internal/domain/user/entity"
	"github.com/donech/nirvana/internal/domain/user/repository"
	"github.com/donech/tool/xdb"
)

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

type UserService struct {
	userRepository *repository.UserRepository
}

func (s UserService) Login(account, password string) entity.User {
	return entity.User{
		Entity:   xdb.Entity{ID: 1},
		Name:     "solar",
		Phone:    "180010123261",
		Email:    "solarpwx@yeah.net",
		Password: "piao1234",
		Status:   1,
		CUDTime:  xdb.CUDTime{},
	}
}

func (s UserService) Migration() {
	s.userRepository.DB.AutoMigrate(entity.User{})
	s.userRepository.DB.AutoMigrate(entity.Account{})
}
