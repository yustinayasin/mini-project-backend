package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id          int            `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	PhoneNumber string         `json:"phoneNumber"`
	Street      string         `json:"street"`
	Address     string         `json:"address"`
	PostalCode  string         `json:"postalCode"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm"index" json:"deletedAt"`
	//News []news.News //one to many
}
