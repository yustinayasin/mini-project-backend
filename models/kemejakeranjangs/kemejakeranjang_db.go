package kemejakeranjangs

import (
	"time"

	"gorm.io/gorm"
)

type KemejaKeranjang struct {
	IdKemeja    int            `gorm:"primaryKey;unique" json:"idKemeja"`
	IdKeranjang int            `gorm:"primaryKey;unique" json:"idKeranjang"`
	Jumlah      int            `json:"jumlah"`
	Size        int            `json:"size"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm"index" json:"deletedAt"`
	//News []news.News //one to many
}
