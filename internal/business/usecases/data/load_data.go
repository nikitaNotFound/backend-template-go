package data

import "app/internal/business/models"

func (u *DataUseCases) LoadData() (models.Data, error) {
	return models.Data{
		Location:  "Minsk",
		Temp:      25,
		Humidity:  70,
		Pressure:  100,
		CreatedAt: 1629886400,
	}, nil
}
