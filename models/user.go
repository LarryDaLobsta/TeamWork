package main

import (
   "fmt"
   "github.com/google/uuid"
   //"gorm.io/gorm"
   //"gorm.io/driver/sqlite"
)




type User struct{
	//gorm.Model
	username string
	password string
	UserId int
	UserUUID uuid.UUID
}

func main() {
	var userOne User
	userOne.username = "Larry"
	userOne.password = "12345"
	userOne.UserId= 23311
	userOne.UserUUID = uuid.New()

	fmt.Println(userOne.username)
	fmt.Println(userOne.password)
	fmt.Println(userOne.UserId)
	fmt.Println(userOne.UserUUID.String())

}



