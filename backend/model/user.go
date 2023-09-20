package model

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Profile  string `db:"profile"`
	Auth     int    `db:"auth"`
}
