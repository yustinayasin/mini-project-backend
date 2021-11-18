package sales

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	Id        int            `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	KemejaId  int            `json:"kemejaId"`
	Percent   int            `json:"percent"`
	Minimum   int            `json:"minimum"`
	Period    time.Time      `json:"period"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm"index" json:"deletedAt"`
	//News []news.News //one to many
}
