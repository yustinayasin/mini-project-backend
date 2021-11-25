package request

import (
	"kemejaku/business/sales"
	"time"
)

type Sale struct {
	Percent          float64   `json:"percent"`
	MinimumPembelian int       `json:"minimumPembelian"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
}

func (sale *Sale) ToUsecase() *sales.Sale {
	return &sales.Sale{
		Percent:          sale.Percent,
		MinimumPembelian: sale.MinimumPembelian,
		StartDate:        sale.StartDate,
		EndDate:          sale.EndDate,
	}
}
