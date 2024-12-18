package business

import "app/internal/domain"

type UserRepository interface {
	GetUserByLogin(login string) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type UserCacheRepository interface {
	GetUserByLogin(login string) (*domain.User, error)
}
