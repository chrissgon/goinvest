package stock

type stockSearchRepoMock struct{}

var StockEntityMockPETR4 = StockEntity{
	ID:         "PETR4",
	Company:    "Petrobras",
	Price:      36.93,
	NetProfit:  78760000000.00,
	NetRevenue: 499060000000.00,
	NetEquity:  373480000000.00,
	NetDebt:    7143243975.00,
	Shares:     13044496930,
	Dividend: 6.3937,
}

var StockEntityMockVALE3 = StockEntity{
	ID:         "VALE3",
	Company:    "Vale",
	Price:      60.76,
	NetProfit:  48730000000.00,
	NetRevenue: 210090000000.00,
	NetEquity:  205450000000.00,
	NetDebt:    47760000000.00,
	Shares:     786043449,
	Dividend: 7.164,
}

var StockEntityMockYDUQ3 = StockEntity{
	ID:         "YDUQ3",
	Company:    "YDUQS",
	Price:      10.24,
	NetProfit:  146400000.00,
	NetRevenue: 5310000000.00,
	NetEquity:  2140000000.00,
	NetDebt:    4560000000.00,
	Shares:     445312500,
	Dividend: 0.1745,
}

func NewStockSearchRepoMock() StockSearchRepo {
	return &stockSearchRepoMock{}
}

func (v *stockSearchRepoMock) Run(ID string) (*StockEntity, error) {
	stock := StockEntityMockVALE3
	return &stock, nil
}
