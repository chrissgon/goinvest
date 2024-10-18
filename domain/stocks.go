package domain

import (
	"math"
)

const PER_MARK = 15
const PBV_MARK = 2
const PROFIT_MARGIN_MARK = 10
const ROE_MARK = 10
const DEBIT_RATIO_MARK = 70
const DIVIDEND_YELD_MARK = 2

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
func DividendYield(dividend, sharePrice float64) float64{
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
