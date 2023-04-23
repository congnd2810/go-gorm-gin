package main

import "handleDb"

func main() {
	db := handleDb.InitDb()

	users := []handleDb.User{
		{Name: "Jane", Age: 25, Email: "jane@hostmail.com"},
		{Name: "Bob", Age: 22, Email: "bob@hostmail.com"},
	}
	db.Create(&users)
}
