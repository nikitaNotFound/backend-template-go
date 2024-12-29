package business

import "app/internal/business/models"

type BusinessDeps struct {
	UserRepo      UserRepo
	UserCacheRepo UserCacheRepo
	DataCacheRepo DataCacheRepo
}

type UserRepo interface {
	GetUserByLogin(login string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserCacheRepo interface {
	GetUserByLogin(login string) (*models.User, error)
}

type DataCacheRepo interface {
	GetData() (*models.Data, error)
}
