package model

import (
	"html"
	"strings"
	"swiftstock_api/app/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
// Holds user related info, handles saving and loading of user data
// Also used by auth system to login users in

type User struct {
	gorm.Model
	Username string `gorm:"size:255; not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
}


// Save the user to database
func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

// Gorm Hook
// Called before we save user data, process the input
func (user *User) BeforeSave(*gorm.DB) error {
	// TODO: Cost configuration
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User

	err := database.Database.Where("username=?", username).Find(&user).Error

	if(err != nil) {
		return User{}, err 
	}

	return user, nil
}