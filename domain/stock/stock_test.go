package stock

import (
	"errors"
	"testing"
)

func TestDomain_StockIsValid(t *testing.T) {
	stock := StockEntityMockVALE3
	stock.ID = ""

	err := stock.IsValid()

	if !errors.Is(err, ErrStockIDInvalid) {
		t.Fatalf("stock IsValid should return the error: %v, but got %v", ErrStockIDInvalid, err)
	}

	stock = StockEntityMockVALE3
	stock.Company = ""

	err = stock.IsValid()

	if !errors.Is(err, ErrStockCompanyInvalid) {
		t.Fatalf("stock IsValid should return the error: %v, but got %v", ErrStockCompanyInvalid, err)
	}

	stock = StockEntityMockVALE3
	stock.NetProfit = 0

	err = stock.IsValid()

	if !errors.Is(err, ErrStockNetProfitInvalid) {
		t.Fatalf("stock IsValid should return the error: %v, but got %v", ErrStockNetProfitInvalid, err)
	}

	stock = StockEntityMockVALE3
	stock.NetRevenue = 0

	err = stock.IsValid()

	if !errors.Is(err, ErrStockNetRevenueInvalid) {
		t.Fatalf("stock IsValid should return the error: %v, but got %v", ErrStockNetRevenueInvalid, err)
	}

	stock = StockEntityMockVALE3
	stock.NetEquity = 0

	err = stock.IsValid()

	if !errors.Is(err, ErrStockNetEquityInvalid) {
		t.Fatalf("stock IsValid should return the error: %v, but got %v", ErrStockNetEquityInvalid, err)
	}

	stock = StockEntityMockVALE3
	stock.NetDebt = 0

	err = stock.IsValid()

	if !errors.Is(err, ErrStockNetDebtInvalid) {
		t.Fatalf("stock IsValid should return the error: %v, but got %v", ErrStockNetDebtInvalid, err)
	}

	stock = StockEntityMockVALE3
	stock.Price = 0

	err = stock.IsValid()

	if !errors.Is(err, ErrStockNetPriceInvalid) {
		t.Fatalf("stock IsValid should return the error: %v, but got %v", ErrStockNetPriceInvalid, err)
	}

	stock = StockEntityMockVALE3
	stock.Shares = 0

	err = stock.IsValid()

	if !errors.Is(err, ErrStockNetSharesInvalid) {
		t.Fatalf("stock IsValid should return the error: %v, but got %v", ErrStockNetSharesInvalid, err)
	}

	stock = StockEntityMockVALE3

	err = stock.IsValid()

	if err != nil {
		t.Fatalf("stock IsValid should return nil because is valid, but got %v", err)
	}
}

func TestDomain_StockGetPER(t *testing.T) {
	stock := StockEntityMockVALE3

	per := stock.GetPER()

	if per.Name != PER_NAME {
		t.Fatalf("GetPER should return a struct with Name %v, but got %v", PER_NAME, per.Name)
	}

	if per.Mark != PER_MARK {
		t.Fatalf("GetPER should return a struct with Mark %v, but got %v", PER_MARK, per.Mark)
	}

	value := 0.9801580900145184
	if per.Value != value {
		t.Fatalf("GetPER should return a struct with Value %v, but got %v", value, per.Value)
	}

	if !per.Good {
		t.Fatalf("GetPER Good property should be true, but got false")
	}

	stock = StockEntityMockYDUQ3

	per = stock.GetPER()

	value = 32
	if per.Value != value {
		t.Fatalf("GetPER should return a struct with Value %v, but got %v", value, per.Value)
	}

	if per.Good {
		t.Fatalf("GetPER Good property should be false, but got true")
	}
}

func TestDomain_StockGetPBV(t *testing.T) {
	stock := StockEntityMockVALE3

	pbv := stock.GetPBV()

	if pbv.Name != PBV_NAME {
		t.Fatalf("GetPBV should return a struct with Name %v, but got %v", PBV_NAME, pbv.Name)
	}

	if pbv.Mark != PBV_MARK {
		t.Fatalf("GetPBV should return a struct with Mark %v, but got %v", PBV_MARK, pbv.Mark)
	}

	value := 0.23246738340283887
	if pbv.Value != value {
		t.Fatalf("GetPBV should return a struct with Value %v, but got %v", value, pbv.Value)
	}

	if !pbv.Good {
		t.Fatalf("GetPBV Good property should be true, but got false")
	}

	stock = StockEntityMockYDUQ3

	pbv = stock.GetPBV()

	value = 2.1333333333333333
	if pbv.Value != value {
		t.Fatalf("GetPBV should return a struct with Value %v, but got %v", value, pbv.Value)
	}

	if pbv.Good {
		t.Fatalf("GetPBV Good property should be false, but got true")
	}
}

func TestDomain_StockGetProfitMargin(t *testing.T) {
	stock := StockEntityMockVALE3

	margin := stock.GetProfitMargin()

	if margin.Name != PROFIT_MARGIN_NAME {
		t.Fatalf("GetProfitMargin should return a struct with Name %v, but got %v", PROFIT_MARGIN_NAME, margin.Name)
	}

	if margin.Mark != PROFIT_MARGIN_MARK {
		t.Fatalf("GetProfitMargin should return a struct with Mark %v, but got %v", PROFIT_MARGIN_MARK, margin.Mark)
	}

	value := 23.194821267076016
	if margin.Value != value {
		t.Fatalf("GetProfitMargin should return a struct with Value %v, but got %v", value, margin.Value)
	}

	if !margin.Good {
		t.Fatalf("GetProfitMargin Good property should be true, but got false")
	}

	stock = StockEntityMockYDUQ3

	margin = stock.GetProfitMargin()

	value = 2.7570621468926553
	if margin.Value != value {
		t.Fatalf("GetProfitMargin should return a struct with Value %v, but got %v", value, margin.Value)
	}

	if margin.Good {
		t.Fatalf("GetProfitMargin Good property should be false, but got true")
	}
}

func TestDomain_StockGetROE(t *testing.T) {
	stock := StockEntityMockVALE3

	roe := stock.GetROE()

	if roe.Name != ROE_NAME {
		t.Fatalf("GetROE should return a struct with Name %v, but got %v", ROE_NAME, roe.Name)
	}

	if roe.Mark != ROE_MARK {
		t.Fatalf("GetROE should return a struct with Mark %v, but got %v", ROE_MARK, roe.Mark)
	}

	value := 23.718666342175712
	if roe.Value != value {
		t.Fatalf("GetROE should return a struct with Value %v, but got %v", value, roe.Value)
	}

	if !roe.Good {
		t.Fatalf("GetROE Good property should be true, but got false")
	}

	stock = StockEntityMockYDUQ3

	roe = stock.GetROE()

	value = 6.841121495327103
	if roe.Value != value {
		t.Fatalf("GetROE should return a struct with Value %v, but got %v", value, roe.Value)
	}

	if roe.Good {
		t.Fatalf("GetROE Good property should be false, but got true")
	}
}

func TestDomain_StockGetDebtRatio(t *testing.T) {
	stock := StockEntityMockVALE3

	debit := stock.GetDebtRatio()

	if debit.Name != DEBIT_RATIO_NAME {
		t.Fatalf("GetDebtRatio should return a struct with Name %v, but got %v", DEBIT_RATIO_NAME, debit.Name)
	}

	if debit.Mark != DEBIT_RATIO_MARK {
		t.Fatalf("GetDebtRatio should return a struct with Mark %v, but got %v", DEBIT_RATIO_MARK, debit.Mark)
	}

	value := 23.24653200292042
	if debit.Value != value {
		t.Fatalf("GetDebtRatio should return a struct with Value %v, but got %v", value, debit.Value)
	}

	if !debit.Good {
		t.Fatalf("GetDebtRatio Good property should be true, but got false")
	}

	stock = StockEntityMockYDUQ3

	debit = stock.GetDebtRatio()

	value = 213.0841121495327
	if debit.Value != value {
		t.Fatalf("GetDebtRatio should return a struct with Value %v, but got %v", value, debit.Value)
	}

	if debit.Good {
		t.Fatalf("GetDebtRatio Good property should be false, but got true")
	}
}

func TestDomain_StockGetDividenYeld(t *testing.T) {
	stock := StockEntityMockVALE3

	dividend := stock.GetDividenYeld()

	if dividend.Name != DIVIDEND_YELD_NAME {
		t.Fatalf("GetDividenYeld should return a struct with Name %v, but got %v", DIVIDEND_YELD_NAME, dividend.Name)
	}

	if dividend.Mark != DIVIDEND_YELD_MARK {
		t.Fatalf("GetDividenYeld should return a struct with Mark %v, but got %v", DIVIDEND_YELD_MARK, dividend.Mark)
	}

	value := 11.790651744568795
	if dividend.Value != value {
		t.Fatalf("GetDividenYeld should return a struct with Value %v, but got %v", value, dividend.Value)
	}

	if !dividend.Good {
		t.Fatalf("GetDividenYeld Good property should be true, but got false")
	}

	stock = StockEntityMockYDUQ3

	dividend = stock.GetDividenYeld()

	value = 1.7041015625
	if dividend.Value != value {
		t.Fatalf("GetDividenYeld should return a struct with Value %v, but got %v", value, dividend.Value)
	}

	if dividend.Good {
		t.Fatalf("GetDividenYeld Good property should be false, but got true")
	}
}

func TestDomain_StockCheckStockID(t *testing.T) {
	have := CheckStockID("VALE3")

	if have != nil {
		t.Fatalf("CheckStockID should return nil because ID is valid")
	}

	have = CheckStockID("VALE333")

	if !errors.Is(have, ErrStockIDInvalid) {
		t.Fatalf("CheckStockID should return an error because ID is invalid")
	}
}

func TestDomain_StockValuePerShare(t *testing.T) {
	have := ValuePerShare(31643000000.00, 13044496930)
	expected := 2.42

	if have != expected {
		t.Fatalf("ValuePerShare should return %v but got %v", expected, have)
	}
}

func TestDomain_StockPER(t *testing.T) {
	vps := ValuePerShare(StockEntityMockVALE3.NetProfit, StockEntityMockVALE3.Shares)
	have := PER(StockEntityMockVALE3.Price, vps)
	expected := 0.9801580900145184

	if have != expected {
		t.Fatalf("PER should return %v but got %v", expected, have)
	}
}

func TestDomain_StockPBV(t *testing.T) {
	vps := ValuePerShare(StockEntityMockVALE3.NetEquity, StockEntityMockVALE3.Shares)
	have := PBV(StockEntityMockVALE3.Price, vps)
	expected := 0.23246738340283887

	if have != expected {
		t.Fatalf("PBV should return %v but got %v", expected, have)
	}
}

func TestDomain_StockProfitMargin(t *testing.T) {
	have := ProfitMargin(StockEntityMockVALE3.NetProfit, StockEntityMockVALE3.NetRevenue)
	expected := 23.194821267076016

	if have != expected {
		t.Fatalf("ProfitMargin should return %v but got %v", expected, have)
	}
}

func TestDomain_StockROE(t *testing.T) {
	have := ROE(StockEntityMockVALE3.NetProfit, StockEntityMockVALE3.NetEquity)
	expected := 23.718666342175712

	if have != expected {
		t.Fatalf("ProfitMargin should return %v but got %v", expected, have)
	}
}

func TestDomain_StockDebitRatio(t *testing.T) {
	have := DebtRatio(StockEntityMockVALE3.NetDebt, StockEntityMockVALE3.NetEquity)
	expected := 23.24653200292042

	if have != expected {
		t.Fatalf("DebtRatio should return %v but got %v", expected, have)
	}
}
func TestDomain_StockDividendYield(t *testing.T) {
	have := DividendYield(0.87, 26.74)
	expected := 3.2535527299925207

	if have != expected {
		t.Fatalf("DividendYield should return %v but got %v", expected, have)
	}
}

func TestDomain_StockGoodPER(t *testing.T) {
	vps := ValuePerShare(StockEntityMockVALE3.NetProfit, StockEntityMockVALE3.Shares)
	per := PER(StockEntityMockVALE3.Price, vps)

	if !GoodPER(per) {
		t.Fatalf("GoodPER should return true, but got false")
	}

	vps = ValuePerShare(StockEntityMockYDUQ3.NetProfit, StockEntityMockYDUQ3.Shares)
	per = PER(StockEntityMockYDUQ3.Price, vps)

	if GoodPER(per) {
		t.Fatalf("GoodPER should return false, but got true")
	}
}

func TestDomain_StockGoodPBV(t *testing.T) {
	vps := ValuePerShare(StockEntityMockVALE3.NetEquity, StockEntityMockVALE3.Shares)
	pbv := PBV(StockEntityMockVALE3.Price, vps)

	if !GoodPBV(pbv) {
		t.Fatalf("GoodPBV should return true, but got false")
	}

	vps = ValuePerShare(StockEntityMockYDUQ3.NetEquity, StockEntityMockYDUQ3.Shares)
	pbv = PBV(StockEntityMockYDUQ3.Price, vps)

	if GoodPBV(pbv) {
		t.Fatalf("GoodPBV should return false, but got true")
	}
}

func TestDomain_StockGoodProfitMargin(t *testing.T) {
	margin := ProfitMargin(StockEntityMockVALE3.NetProfit, StockEntityMockVALE3.NetRevenue)

	if !GoodProfitMargin(margin) {
		t.Fatalf("GoodProfitMargin should return true, but got false")
	}

	margin = ProfitMargin(StockEntityMockYDUQ3.NetProfit, StockEntityMockYDUQ3.NetRevenue)

	if GoodProfitMargin(margin) {
		t.Fatalf("GoodProfitMargin should return false, but got true")
	}
}

func TestDomain_StockGoodROE(t *testing.T) {
	roe := ROE(StockEntityMockVALE3.NetProfit, StockEntityMockVALE3.NetEquity)

	if !GoodROE(roe) {
		t.Fatalf("GoodROE should return true, but got false")
	}

	roe = ROE(StockEntityMockYDUQ3.NetProfit, StockEntityMockYDUQ3.NetEquity)

	if GoodROE(roe) {
		t.Fatalf("GoodROE should return false, but got true")
	}
}
func TestDomain_StockGoodDebitRatio(t *testing.T) {
	debit := DebtRatio(StockEntityMockVALE3.NetDebt, StockEntityMockVALE3.NetEquity)

	if !GoodDebitRatio(debit) {
		t.Fatalf("GoodDebitRatio should return true, but got false")
	}

	debit = DebtRatio(StockEntityMockYDUQ3.NetDebt, StockEntityMockYDUQ3.NetEquity)

	if GoodDebitRatio(debit) {
		t.Fatalf("GoodDebitRatio should return false, but got true")
	}
}
func TestDomain_StockGoodDividendYield(t *testing.T) {
	dy := DividendYield(0.87, 26.74)

	if !GoodDividendYield(dy) {
		t.Fatalf("GoodDividendYield should return true, but got false")
	}

	dy = DividendYield(0.01, 7.85)

	if GoodDividendYield(dy) {
		t.Fatalf("GoodDividendYield should return false, but got true")
	}
}
