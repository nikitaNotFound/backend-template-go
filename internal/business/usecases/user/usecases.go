package user

import "app/internal/business"

type UserUseCases struct {
	userRep      *business.UserRepo
	userCacheRep *business.UserCacheRepo
}

func NewUserUseCases(
	userRep *business.UserRepo,
	userCacheRep *business.UserCacheRepo,
) *UserUseCases {
	return &UserUseCases{
		userRep:      userRep,
		userCacheRep: userCacheRep,
	}
}
