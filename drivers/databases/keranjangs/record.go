package keranjangs

import (
	"kemejaku/business/kemejakeranjangs"
	"kemejaku/business/keranjangs"
	"time"

	"gorm.io/gorm"
)

type Keranjang struct {
	Id              int                                `gorm:"primaryKey;unique;autoIncrement:true" json:"id"`
	IdUser          int                                `json:"userId"`
	Status          bool                               `json:"status"`
	KemejaKeranjang []kemejakeranjangs.KemejaKeranjang `gorm:"foreignKey:IdKeranjang"`
	CreatedAt       time.Time                          `gorm:"autoCreateTime"`
	UpdatedAt       time.Time                          `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt                     `gorm:"index"`
}

func (keranjang Keranjang) ToUsecase() keranjangs.Keranjang {
	return keranjangs.Keranjang{
		Id:              keranjang.Id,
		IdUser:          keranjang.IdUser,
		Status:          keranjang.Status,
		KemejaKeranjang: keranjang.KemejaKeranjang,
		CreatedAt:       keranjang.CreatedAt,
		UpdatedAt:       keranjang.UpdatedAt,
		DeletedAt:       keranjang.DeletedAt,
	}
}

func ToUsecaseList(keranjang []Keranjang) []keranjangs.Keranjang {
	var newKeranjangs []keranjangs.Keranjang

	for _, v := range keranjang {
		newKeranjangs = append(newKeranjangs, v.ToUsecase())
	}

	return newKeranjangs
}

func FromUsecase(keranjang keranjangs.Keranjang) Keranjang {
	return Keranjang{
		Id:        keranjang.Id,
		IdUser:    keranjang.IdUser,
		Status:    keranjang.Status,
		CreatedAt: keranjang.CreatedAt,
		UpdatedAt: keranjang.UpdatedAt,
		DeletedAt: keranjang.DeletedAt,
	}
}
