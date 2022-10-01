package models

import (
	"github.com/BarunBlog/Golang_Jwt_Authentication/pkg/config"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	DB = config.GetDb()
	DB.AutoMigrate(&User{})
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
