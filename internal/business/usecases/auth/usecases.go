package auth

import "app/internal/business"

type AuthUseCases struct {
	*business.BusinessDeps
}

func NewAuthUseCases(deps *business.BusinessDeps) *AuthUseCases {
	return &AuthUseCases{
		BusinessDeps: deps,
	}
}
