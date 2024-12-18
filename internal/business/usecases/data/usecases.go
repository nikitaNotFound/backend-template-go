package data

import "app/internal/business"

type DataUseCases struct {
	dataCacheRep *business.DataCacheRepo
}

func NewDataUseCases(dataCacheRep *business.DataCacheRepo) *DataUseCases {
	return &DataUseCases{
		dataCacheRep: dataCacheRep,
	}
}
