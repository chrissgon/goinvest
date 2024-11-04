package stock

import (
	"errors"
	"regexp"

	"github.com/chrissgon/goinvest/domain"
)

type StockEntity struct {
	ID            string
	Company       string
	NetProfit     float64
	NetRevenue    float64
	NetEquity     float64
	NetDebt       float64
	Price         float64
	Dividend      float64
	DividendYield float64
	Shares        int
}

const PER_MARK = 15
const PER_NAME = "per"
const PER_LABEL = "P/L (Preço / Lucro Líquido por Ação)"

const PBV_MARK = 2
const PBV_NAME = "pbv"
const PBV_LABEL = "P/VPA (Preço / Valor Patrimonial da Ação)"

const PROFIT_MARGIN_MARK = 10
const PROFIT_MARGIN_NAME = "profitMargin"
const PROFIT_MARGIN_LABEL = "Margem Líquida (Lucro Líquido / Receita Líquida)"

const ROE_MARK = 10
const ROE_NAME = "roe"
const ROE_LABEL = "ROE (Lucro Líquido / Patrimônio Líquido)"

const DEBIT_RATIO_MARK = 70
const DEBIT_RATIO_NAME = "debitRatio"
const DEBIT_RATIO_LABEL = "DL/PL (Dívida Líquida / Patrimônio Líquido)"

const DIVIDEND_YELD_MARK = 2
const DIVIDEND_YELD_NAME = "dividendYield"
const DIVIDEND_YELD_LABEL = "Dividend Yield (Proventos por Ação / Preço da Ação)"

var ErrStockIDInvalid = errors.New("stock ID is invalid")
var ErrStockNotFound = errors.New("stock not found")
var ErrStockCompanyInvalid = errors.New("stock company is invalid")
var ErrStockNetProfitInvalid = errors.New("stock net profit is invalid")
var ErrStockNetRevenueInvalid = errors.New("stock net revenue is invalid")
var ErrStockNetEquityInvalid = errors.New("stock net equity is invalid")
var ErrStockNetDebtInvalid = errors.New("stock net debt is invalid")
var ErrStockNetPriceInvalid = errors.New("stock net price is invalid")
var ErrStockNetSharesInvalid = errors.New("stock net shares is invalid")
var ErrStockDividendYieldInvalid = errors.New("stock dividend yield is invalid")

func (entity *StockEntity) IsValid() error {
	err := CheckStockID(entity.ID)

	if err != nil {
		return err
	}
	if entity.Company == "" {
		return ErrStockCompanyInvalid
	}
	if entity.NetProfit == 0 {
		return ErrStockNetProfitInvalid
	}
	if entity.NetRevenue == 0 {
		return ErrStockNetRevenueInvalid
	}
	if entity.NetEquity == 0 {
		return ErrStockNetEquityInvalid
	}
	if entity.NetDebt == 0 {
		return ErrStockNetDebtInvalid
	}
	if entity.Price == 0 {
		return ErrStockNetPriceInvalid
	}
	if entity.Dividend == 0 && entity.DividendYield == 0 {
		return ErrStockDividendYieldInvalid
	}
	if entity.Shares == 0 {
		return ErrStockNetSharesInvalid
	}

	return nil
}

func (entity *StockEntity) GetPER() domain.Indicator {
	vps := domain.ValuePerShare(entity.NetProfit, entity.Shares)
	per := PER(entity.Price, vps)

	return domain.Indicator{
		Name:  PER_NAME,
		Label: PER_LABEL,
		Mark:  PER_MARK,
		Value: per,
		Good:  GoodPER(per),
	}
}

func (entity *StockEntity) GetPBV() domain.Indicator {
	vps := domain.ValuePerShare(entity.NetEquity, entity.Shares)
	pbv := PBV(entity.Price, vps)

	return domain.Indicator{
		Name:  PBV_NAME,
		Label: PBV_LABEL,
		Mark:  PBV_MARK,
		Value: pbv,
		Good:  GoodPBV(pbv),
	}
}

func (entity *StockEntity) GetProfitMargin() domain.Indicator {
	margin := ProfitMargin(entity.NetProfit, entity.NetRevenue)

	return domain.Indicator{
		Name:  PROFIT_MARGIN_NAME,
		Label: PROFIT_MARGIN_LABEL,
		Mark:  PROFIT_MARGIN_MARK,
		Value: margin,
		Good:  GoodProfitMargin(margin),
	}
}

func (entity *StockEntity) GetROE() domain.Indicator {
	roe := ROE(entity.NetProfit, entity.NetEquity)

	return domain.Indicator{
		Name:  ROE_NAME,
		Label: ROE_LABEL,
		Mark:  PROFIT_MARGIN_MARK,
		Value: roe,
		Good:  GoodROE(roe),
	}
}

func (entity *StockEntity) GetDebtRatio() domain.Indicator {
	debt := DebtRatio(entity.NetDebt, entity.NetEquity)

	return domain.Indicator{
		Name:  DEBIT_RATIO_NAME,
		Label: DEBIT_RATIO_LABEL,
		Mark:  DEBIT_RATIO_MARK,
		Value: debt,
		Good:  GoodDebitRatio(debt),
	}
}

func (entity *StockEntity) GetDividenYield() domain.Indicator {
	dividend := DividendYield(entity.Dividend, entity.Price)

	if entity.DividendYield != 0 {
		dividend = entity.DividendYield
	}

	return domain.Indicator{
		Name:  DIVIDEND_YELD_NAME,
		Label: DIVIDEND_YELD_LABEL,
		Mark:  DIVIDEND_YELD_MARK,
		Value: dividend,
		Good:  GoodDividendYield(dividend),
	}
}

func CheckStockID(ID string) error {
	matched, err := regexp.MatchString("^[a-zA-Z]{4}(1|2|3|4|5|6|7|8|9|10|11)$", ID)

	if err != nil {
		return err
	}

	if !matched {
		return ErrStockIDInvalid
	}

	return nil
}
func PER(stockPrice, vps float64) float64 {
	return stockPrice / vps
}
func PBV(stockPrice, vps float64) float64 {
	return stockPrice / vps
}
func ProfitMargin(netProfit, netRevenue float64) float64 {
	return netProfit / netRevenue * 100
}
func ROE(netProfit, netEquity float64) float64 {
	return netProfit / netEquity * 100
}
func DebtRatio(netDebt, netEquity float64) float64 {
	return netDebt / netEquity * 100
}
func DividendYield(dividend, stockPrice float64) float64 {
	return dividend / stockPrice * 100
}

func GoodPER(pricePerEarning float64) bool {
	return pricePerEarning < PER_MARK
}
func GoodPBV(pricePerAsset float64) bool {
	return pricePerAsset < PBV_MARK
}
func GoodProfitMargin(margin float64) bool {
	return margin > PROFIT_MARGIN_MARK
}
func GoodROE(roe float64) bool {
	return roe > ROE_MARK
}
func GoodDebitRatio(debit float64) bool {
	return debit < DEBIT_RATIO_MARK
}
func GoodDividendYield(dy float64) bool {
	return dy > DIVIDEND_YELD_MARK
}
