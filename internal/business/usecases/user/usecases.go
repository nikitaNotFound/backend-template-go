package user

import "app/internal/business"

type AuthUseCases struct {
	userRep      *business.UserRepo
	userCacheRep *business.UserCacheRepo
}

func NewUserUseCases(
	userRep *business.UserRepo,
	userCacheRep *business.UserCacheRepo,
) *AuthUseCases {
	return &AuthUseCases{
		userRep:      userRep,
		userCacheRep: userCacheRep,
	}
}
