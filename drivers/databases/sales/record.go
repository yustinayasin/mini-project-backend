package sales

import (
	"kemejaku/business/kemejas"
	"kemejaku/business/sales"
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	Id               int `gorm:"primaryKey;unique;autoIncrement:true"`
	Percent          float64
	MinimumPembelian int
	StartDate        time.Time
	EndDate          time.Time
	Kemejas          []kemejas.Kemeja `gorm:"foreignKey:IdSale;references:Id"`
	CreatedAt        time.Time        `gorm:"autoCreateTime"`
	UpdatedAt        time.Time        `gorm:"autoUpdateTime"`
	DeletedAt        gorm.DeletedAt   `gorm:"index"`
}

func (sale Sale) ToUsecase() sales.Sale {
	return sales.Sale{
		Id:               sale.Id,
		Percent:          sale.Percent,
		MinimumPembelian: sale.MinimumPembelian,
		StartDate:        sale.StartDate,
		EndDate:          sale.EndDate,
		Kemejas:          sale.Kemejas,
		CreatedAt:        sale.CreatedAt,
		UpdatedAt:        sale.UpdatedAt,
		DeletedAt:        sale.DeletedAt,
	}
}

func ToUsecaseList(sale []Sale) []sales.Sale {
	var newSales []sales.Sale

	for _, v := range sale {
		newSales = append(newSales, v.ToUsecase())
	}

	return newSales
}

func FromUsecase(sale sales.Sale) Sale {
	return Sale{
		Id:               sale.Id,
		Percent:          sale.Percent,
		MinimumPembelian: sale.MinimumPembelian,
		StartDate:        sale.StartDate,
		EndDate:          sale.EndDate,
		Kemejas:          sale.Kemejas,
		CreatedAt:        sale.CreatedAt,
		UpdatedAt:        sale.UpdatedAt,
		DeletedAt:        sale.DeletedAt,
	}
}
