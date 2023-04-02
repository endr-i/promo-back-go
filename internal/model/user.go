package model

type User struct {
	Id    string `db:"id"`
	Email string `db:"email"`
	Phone string `db:"phone"`
}
