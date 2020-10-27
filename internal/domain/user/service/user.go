package service

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"

	"github.com/donech/nirvana/internal/common"
	"github.com/donech/nirvana/internal/domain/user/entity"
	"github.com/donech/nirvana/internal/domain/user/repository"
)

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

type UserService struct {
	userRepository *repository.UserRepository
}

func (s UserService) Login(ctx context.Context, account, password string) (entity.User, error) {
	password = common.SignPassword(password)
	user, err := s.userRepository.ItemByEmail(ctx, account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.User{}, errors.New("account or password error")
		}
		return entity.User{}, err
	}
	if user.Password != password {
		return entity.User{}, errors.New("account or password error")
	}
	return user, nil
}

func (s UserService) Migration() {
	s.userRepository.DB.AutoMigrate(entity.User{})
	s.userRepository.DB.AutoMigrate(entity.Account{})
}
