package models

type User struct {
	ID       int
	Username string
	Password string
}

type Book struct {
	ID     int
	Title  string
	Author string
}
