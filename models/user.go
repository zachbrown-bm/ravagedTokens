package models

type UserType string

const (
	Admin UserType = "admin"
	Paid  UserType = "paid"
	Free  UserType = "free"
)

type User struct {
	Id    int
	Name  string
	Email string
	Type  UserType
}
