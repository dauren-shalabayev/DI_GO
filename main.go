package main

import (
	print "book/formatter"
	"fmt"
)

type User struct {
	Email    string
	Username string
}

type Admin struct {
	User
	Level int
}

// фнкция для уровня
func (a *Admin) LevelShow() string {
	return fmt.Sprintf("admin %s, Level %d", a.Username, a.Level)
}

type Printer interface {
	GetUserName() string
}

func (u *User) GetUserName() string {
	return fmt.Sprintf("Email: %s, Username: %s", u.Email, u.Username)
}

func main() {
	print.Format("sadas")
	user := &User{
		Email:    "dauren@ya.ru",
		Username: "dauren10",
	}

	adm := &Admin{
		User:  User{Username: "Aidar", Email: "aidar.admin@yandex.kz"},
		Level: 1,
	}

	fmt.Println(user.GetUserName())
	fmt.Println(adm.LevelShow())

}
