package models

type Contact struct {
	Name    string `db:"name" json:"name"`
	Email   string `db:"email" json:"email"`
	Message string `db:"message" json:"message"`
}
