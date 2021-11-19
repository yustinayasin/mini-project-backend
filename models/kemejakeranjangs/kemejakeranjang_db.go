package kemejakeranjangs

import (
	"time"

	"gorm.io/gorm"
)

type KemejaKeranjang struct {
	Id          int    `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	IdKemeja    int    `json:"idKemeja"`
	IdKeranjang int    `json:"idKeranjang"`
	Jumlah      int    `json:"jumlah"`
	Size        string `json:"size"`
	// Keranjang   keranjangs.Keranjang `gorm:"foreignKey:IdKeranjang"`
	// Kemeja      kemejas.Kemeja       `gorm:"foreignKey:IdKemeja"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm"index" json:"deletedAt"`
}
