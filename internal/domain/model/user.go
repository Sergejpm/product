package model

type User struct {
	Id       uint   `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
