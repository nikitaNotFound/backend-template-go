package domain

type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	PwdHash   string `json:"pwdHash"`
	CreatedAt int64  `json:"created_at"`
}
