package models

import (
	"errors"
	"time"

	"github.com/foekall/cattle-management/pkg/config"
	"github.com/foekall/cattle-management/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	ID          int64      `gorm:"primaryKey"`
	FullName    string     `gorm:"type:varchar(30)" binding:"required"`
	DateOfBirth *time.Time `validate:"required"`
	PhoneNumber string     `gorm:"type:varchar(20)" binding:"required"`
	Email       *string    `gorm:"type:varchar(30);unique_index" binding:"required,email" validate:"is-duplicate-email"`
	Password    string     `gorm:"type:varchar(255);->:false;<-:create" binding:"required,max=20,min=6"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

//create user
func CreateUser(u *User) (*User, error) {
	var Users []User

	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	u.Password = string(hashedPassword)

	res := db.First(&Users, "email = ?", u.Email)
	if res.RowsAffected > 0 {
		return nil, errors.New("email already registered. Please choose another one")
	}

	db.Create(&u)
	return u, nil
}

var Users []User

func GetAllUser(page int, size int) []User {
	Offset := utils.Paginate(page, size)
	db.Scopes(Offset).Find(&Users).Order("id")
	return Users
}

func GetUserById(id int64) []User {
	db.First(&Users, id)
	return Users
}

func UpdateUser(u *User) (*User, error) {
	res := db.First(&Users, u.ID)
	if res.RowsAffected > 0 {
		return nil, errors.New("record not found")
	}
	db.Updates(&u)
	return u, nil
}

func DeleteUser(id int64) (string, error) {

	res := db.First(&Users, id)
	if res.RowsAffected > 0 {
		return "", errors.New("record not found")
	}
	db.Delete(&Users)
	return "User successfully deleted", nil

}
