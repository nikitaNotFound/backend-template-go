package domain

type User struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	PwdHash   string `json:"pwdHash"`
	CreatedAt int64  `json:"created_at"`
}

type Data struct {
	Location  string `json:"location"`
	Temp      int64  `json:"temp"`
	Humidity  int64  `json:"humidity"`
	Pressure  int64  `json:"pressure"`
	CreatedAt int64  `json:"created_at"`
}
