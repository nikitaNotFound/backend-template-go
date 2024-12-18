package domain

type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	PwdHash   string `json:"pwdHash"`
	CreatedAt int64  `json:"created_at"`
}

type Data struct {
	Location  string `json:"location"`
	Temp      int    `json:"temp"`
	Humidity  int    `json:"humidity"`
	Pressure  int    `json:"pressure"`
	CreatedAt int64  `json:"created_at"`
}
