package models

type User struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Age   int    `db:"age" json:"age"`
	Email string `db:"email" json:"email"`
}
