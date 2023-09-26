package services

import (
	"context"

	"github.com/ThailanTec/jogo-do-bicho/internal/core/domain"
	"github.com/ThailanTec/jogo-do-bicho/internal/core/ports"
)

type UserService struct {
	repository ports.UserRepository
}

func NewUserServices(repo ports.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}

}

func (us *UserService) CreateUser(ctx context.Context, u domain.User) (user *domain.User, err error) {
	return us.repository.CreateUser(ctx, *user)
}

func (us *UserService) ListUser(ctx context.Context, id uint) (user *domain.User, err error) {
	return us.repository.ListUser(ctx, id)
}

func (us *UserService) DeleteUser(ctx context.Context, id uint) (deleted bool, err error) {
	return us.repository.DeleteUser(ctx, id)
}

func (us *UserService) ListAllUser() (user *domain.User, err error) {
	return us.repository.ListAllUser()
}
