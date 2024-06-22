package main


import (
   "fmt"
   "github.com/google/uuid"
   "gorm.io/gorm"
   //"gorm.io/driver/sqlite"
)



type User struct{
	ent.Schema
}


//useruuid
//userid
//created
//groupbelong to

//create group add later
	// true or false
// role add later
	// this should be an enum of manager and contributor and developer and viewer
//first name
//last name
// username
// email


func (User) Fields() []ent.Field{
	return []ent.Field{
		field.UUID("user_uuid", uuid.UUID{}).Default(uuid.New),
		field.Int("id"),
		field.String("first_name").NotEmpty().MaxLen(30),
		field.String("last_name").NotEmpty().MaxLen(30),
		field.String("username").NotEmpty().Unique().Sensitive(),
		field.String("password").NotEmpty().Unique().MinLen(10).Sensitive(),
	}
}







func main() {
	var userOne User
	userOne.Username = "Larry"
	userOne.Password = "12345"
	userOne.UserId= 23311
	userOne.UserUUID = uuid.New()

	fmt.Println(userOne.Username)
	fmt.Println(userOne.Password)
	fmt.Println(userOne.UserId)
	fmt.Println(userOne.UserUUID.String())

}



