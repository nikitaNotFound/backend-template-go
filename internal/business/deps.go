package business

import "app/internal/domain"

type BusinessDeps struct {
	UserRepo      UserRepo
	UserCacheRepo UserCacheRepo
	DataCacheRepo DataCacheRepo
}

type UserRepo interface {
	GetUserByLogin(login string) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type UserCacheRepo interface {
	GetUserByLogin(login string) (*domain.User, error)
}

type DataCacheRepo interface {
	GetData() (*domain.Data, error)
}
