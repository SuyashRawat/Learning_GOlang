package models

import (
	"errors"
	"time"
	// "log"
	// "taskmanage/db"
	"taskmanage/pkg/utils"

	// "github.com/golang-migrate/migrate/v4/database"
	"gorm.io/gorm"
)

// import "time"
var DB *gorm.DB
type User struct {
    Id            int `db:"id" json:"id"`
    Username      string `db:"username" json:"username"`
    Email         string `db:"email" json:"email"`
    PhoneNo string `db:"phone_no" json:"phoneNo"`
    PasswordHash  string `db:"password_hash" json:"password_hash"`
    CreatedAt     string `db:"created_at" json:"created_at"`
}


func SetDB(database *gorm.DB){

	DB=database
}

func CreateUser(user *User)error{
	if user.Username == "" || user.Email == "" || user.PasswordHash == "" || user.PhoneNo == "" {
		return errors.New("missing fields")
	}
	
	var existinguser User
	result:=DB.Where("email = ?",user.Email).First(&existinguser)
	if result.RowsAffected>0{
		return errors.New("user already exists")
	}
	hashed_password,err:=utils.HashPassword(user.PasswordHash)
	if err!=nil{
		return errors.New("failed to hash")
	}
	user.PasswordHash=hashed_password
	user.CreatedAt = time.Now().Format(time.RFC3339)
creation := DB.Create(&user) // Ensure user.ID is omitted during insert
if creation.Error != nil {
    return errors.New("failed to create user")
}
return nil



}

func Getuserbyemail(email string)(*User,error){
	var user User
	result:=DB.Where("email = ?",email).First(&user)
	if result.Error!=nil{
		return nil,errors.New("User does not exists")
	}
	return &user,nil
}