package ports

import (
	"context"

	"github.com/ThailanTec/jogo-do-bicho/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u domain.User) (user *domain.User, err error)
	UpdateUser(ctx context.Context, u domain.User) (user *domain.User, err error)
	ListUser(ctx context.Context, id uint) (user *domain.User, err error)
	DeleteUser(ctx context.Context, id uint) (deleted bool, err error)
	ListAllUser() (user *domain.User, err error)
	LoginUser(email, password string) (login bool, err error)
}
