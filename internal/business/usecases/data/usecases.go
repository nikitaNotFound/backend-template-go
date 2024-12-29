package data

import "app/internal/business"

type DataUseCases struct {
	*business.BusinessDeps
}

func NewDataUseCases(deps *business.BusinessDeps) *DataUseCases {
	return &DataUseCases{
		BusinessDeps: deps,
	}
}
