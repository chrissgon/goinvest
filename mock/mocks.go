package mock

import (
	"github.com/chrissgon/goinvest/domain"
)

type stockSearchRepoMock struct{}

var StockEntityMockPETR4 = domain.StockEntity{
	Company:    "Petrobras",
	Price:      36.93,
	NetProfit:  78760000000.00,
	NetRevenue: 499060000000.00,
	NetEquity:  373480000000.00,
	NetDebt:    7143243975,
	Shares:     13044496930,
}

func NewStockSearchRepoMock() domain.StockSearchRepo {
	return &stockSearchRepoMock{}
}

func (v *stockSearchRepoMock) Run(ID string) (*domain.StockEntity, error) {
	stock := StockEntityMockPETR4
	return &stock, nil
}
