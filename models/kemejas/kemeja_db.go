package kemejas

import (
	"time"

	"gorm.io/gorm"
)

type Kemeja struct {
	Id          int            `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	StockL      int            `json:"stockL"`
	StockM      int            `json:"stockM"`
	StockS      int            `json:"stockS"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm"index" json:"deletedAt"`
	//News []news.News //one to many
}
