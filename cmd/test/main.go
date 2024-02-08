package main

import (
	"log"

	"github.com/ricardoraposo/blogs-api-go/internal/database"
	"github.com/ricardoraposo/blogs-api-go/internal/entity"
)

func main() {
	db := database.NewDatabase()

	accountDB := database.NewUserDB(db.DB)

	user, _ := entity.NewUser("Ricardo", "r@r.com", "123456", "image")

	_, err := accountDB.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	users, _ := accountDB.GetUsers()

	for _, u := range users {
		log.Println(u.DisplayName)
	}
}
