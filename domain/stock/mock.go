package stock

import "github.com/chrissgon/goinvest/domain"

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
	Dividend:   6.3937,
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
	Dividend:   7.164,
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
	Dividend:   0.1745,
}

var StockIndicatorsMockVALE3 = map[string]domain.Indicator{
	PER_NAME: {
		Name:  PER_NAME,
		Label: PER_LABEL,
		Mark:  PER_MARK,
		Value: 0.9801580900145184,
		Operator: PER_OPERATOR,
		Good:  true,
	},
	PBV_NAME: {
		Name:  PBV_NAME,
		Label: PBV_LABEL,
		Mark:  PBV_MARK,
		Value: 0.23246738340283887,
		Operator: PBV_OPERATOR,
		Good:  true,
	},
	PROFIT_MARGIN_NAME: {
		Name:  PROFIT_MARGIN_NAME,
		Label: PROFIT_MARGIN_LABEL,
		Mark:  PROFIT_MARGIN_MARK,
		Value: 23.194821267076016,
		Operator: PROFIT_MARGIN_OPERATOR,
		Good:  true,
	},
	ROE_NAME: {
		Name:  ROE_NAME,
		Label: ROE_LABEL,
		Mark:  ROE_MARK,
		Value: 23.718666342175712,
		Operator: ROE_OPERATOR,
		Good:  true,
	},
	DEBIT_RATIO_NAME: {
		Name:  DEBIT_RATIO_NAME,
		Label: DEBIT_RATIO_LABEL,
		Mark:  DEBIT_RATIO_MARK,
		Value: 23.24653200292042,
		Operator: DEBIT_RATIO_OPERATOR,
		Good:  true,
	},
	DIVIDEND_YIELD_NAME: {
		Name:  DIVIDEND_YIELD_NAME,
		Label: DIVIDEND_YIELD_LABEL,
		Mark:  DIVIDEND_YIELD_MARK,
		Value: 11.790651744568795,
		Operator: DIVIDEND_YIELD_OPERATOR,
		Good:  true,
	},
}

func NewStockSearchRepoMock() StockSearchRepo {
	return &stockSearchRepoMock{}
}

func (v *stockSearchRepoMock) Run(ID string) (StockEntity, error) {
	return StockEntityMockVALE3, nil
}
