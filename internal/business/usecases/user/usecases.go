package user

import "app/internal/business"

type UserUseCases struct {
	userRep      *business.UserRepository
	userCacheRep *business.UserCacheRepository
}

func NewUserUseCases(
	userRep *business.UserRepository,
	userCacheRep *business.UserCacheRepository,
) *UserUseCases {
	return &UserUseCases{
		userRep:      userRep,
		userCacheRep: userCacheRep,
	}
}
