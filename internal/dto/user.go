package dto

type User struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Email   string `json:"email"`
}