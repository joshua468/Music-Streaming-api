package models

type User struct {
	Name     string     `json:"name"`
	ID       int        `json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Platform []Platform `json:"platform"`
}
