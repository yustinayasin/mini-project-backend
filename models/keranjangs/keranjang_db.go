package keranjangs

import (
	"time"

	"gorm.io/gorm"
)

type Keranjang struct {
	Id        int            `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	UserId    int            `json:"userId"`
	Status    bool           `json:"status"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm"index" json:"deletedAt"`
	//News []news.News //one to many
}
