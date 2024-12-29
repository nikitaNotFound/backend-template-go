package data

import "app/internal/domain"

func (u *DataUseCases) LoadData() (domain.Data, error) {
	return domain.Data{
		Location:  "Minsk",
		Temp:      25,
		Humidity:  70,
		Pressure:  100,
		CreatedAt: 1629886400,
	}, nil
}
