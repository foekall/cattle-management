package models

import (
	"errors"
	"time"

	"github.com/foekall/cattle-management/pkg/config"
	"gorm.io/gorm"
)

// var db *gorm.DB

type Cattle struct {
	gorm.Model
	ID      int64 `gorm:"primaryKey"`
	Breed   string
	OwnerId int64
	// Owner     models.Owner `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Cattle{})
}

func CreateCattle(c *Cattle) *Cattle {
	db.Create(&c)
	return c
}

func GetAllCattles() []Cattle {
	var Cattles []Cattle
	db.Find(&Cattles)
	return Cattles
}

func GetCattleById(id int64) []Cattle {
	var Cattles []Cattle
	db.First(&Cattles, id)
	return Cattles
}

func UpdateCattle(c *Cattle) (*Cattle, error) {
	var Cattles []Cattle
	res := db.First(&Cattles, c.ID)
	if res.RowsAffected > 0 {
		db.Updates(&c)
	} else {
		return nil, errors.New("Record not found!")
	}
	return c, nil
}

func DeleteCattle(id int64) (string, error) {
	var Cattles []Cattle
	res := db.First(&Cattles, id)
	if res.RowsAffected > 0 {
		db.Delete(&Cattles)
	} else {
		return "", errors.New("Record not found!")
	}
	return "Successfully deleted", nil
}
