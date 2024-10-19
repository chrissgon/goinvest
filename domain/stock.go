package domain

import (
	"errors"
	"math"
	"regexp"
)

type StockEntity struct {
	ID         string
	Company    string
	NetProfit  float64
	NetRevenue float64
	NetEquity  float64
	NetDebt    float64
	Price      float64
	Dividend   float64
	Shares     int
}

type StockIndicator struct {
	name string
	mark int
	value float64
	good bool
}

type StockSearchRepo interface {
	Run(ID string) (*StockEntity, error)
}

const PER_MARK = 15
const PBV_MARK = 2
const PROFIT_MARGIN_MARK = 10
const ROE_MARK = 10
const DEBIT_RATIO_MARK = 70
const DIVIDEND_YELD_MARK = 2

var ErrStockInvalidID = errors.New("invalid stock id")
var ErrStockNotFound = errors.New("stock not found")

func CheckStockID(ID string) error {
	matched, err := regexp.MatchString("^[a-zA-Z]{4}[0-9]{1,2}$", ID)

	if err != nil {
		return err
	}

	if !matched {
		return ErrStockInvalidID
	}

	return nil
}
func ValuePerShare(value float64, shares int) float64 {
	return math.Floor((value/float64(shares))*100) / 100
}
func PER(sharePrice, vps float64) float64 {
	return sharePrice / vps
}
func PBV(sharePrice, vps float64) float64 {
	return sharePrice / vps
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
func DividendYield(dividend, sharePrice float64) float64 {
	return dividend / sharePrice * 100
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
