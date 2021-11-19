package keranjangs

import (
	"kemejaku/models/kemejakeranjangs"
	"time"

	"gorm.io/gorm"
)

type Keranjang struct {
	Id              int                                `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	UserId          int                                `json:"userId"`
	Status          bool                               `json:"status"`
	KemejaKeranjang []kemejakeranjangs.KemejaKeranjang `gorm:"foreignKey:IdKeranjang"`
	CreatedAt       time.Time                          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time                          `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt                     `gorm"index" json:"deletedAt"`
}
