package domain

import (
	"errors"
	"testing"
)

func TestCheckStockID(t *testing.T) {
	have := CheckStockID("PETR4")

	if have != nil {
		t.Fatalf("CheckStockID should return nil because ID is valid")
	}

	have = CheckStockID("PETR111")

	if !errors.Is(have, ErrStockInvalidID) {
		t.Fatalf("CheckStockID should return an error because ID is invalid")
	}
}

func TestValuePerShare(t *testing.T) {
	have := ValuePerShare(31643000000.00, 13044496930)
	expected := 2.42

	if have != expected {
		t.Fatalf("ValuePerShare should return %v but got %v", expected, have)
	}
}

func TestPER(t *testing.T) {
	vps := ValuePerShare(31643000000.00, 13044496930)
	have := PER(26.74, vps)
	expected := 11.049586776859504

	if have != expected {
		t.Fatalf("PER should return %v but got %v", expected, have)
	}
}

func TestPBV(t *testing.T) {
	vps := ValuePerShare(303619000000.00, 13044496930)
	have := PBV(26.74, vps)
	expected := 1.1491190373871938

	if have != expected {
		t.Fatalf("PBV should return %v but got %v", expected, have)
	}
}

func TestProfitMargin(t *testing.T) {
	have := ProfitMargin(31643000000.00, 302337000000.00)
	expected := 10.46613547134489

	if have != expected {
		t.Fatalf("ProfitMargin should return %v but got %v", expected, have)
	}
}

func TestROE(t *testing.T) {
	have := ROE(31643000000.00, 303619000000.00)
	expected := 10.421943290769024

	if have != expected {
		t.Fatalf("ProfitMargin should return %v but got %v", expected, have)
	}
}

func TestDebitRatio(t *testing.T) {
	have := DebtRatio(108643000000.00, 303619000000.00)
	expected := 35.782674997282776

	if have != expected {
		t.Fatalf("DebtRatio should return %v but got %v", expected, have)
	}
}
func TestDividendYield(t *testing.T) {
	have := DividendYield(0.87, 26.74)
	expected := 3.2535527299925207

	if have != expected {
		t.Fatalf("DividendYield should return %v but got %v", expected, have)
	}
}

func TestGoodPER(t *testing.T) {
	vps := ValuePerShare(31643000000.00, 13044496930)
	price := PER(26.74, vps)

	if !GoodPER(price) {
		t.Fatalf("GoodPER should return true, but got false")
	}

	vps = ValuePerShare(29460000.00, 87941972)
	price = PER(19.07, vps)

	if GoodPER(price) {
		t.Fatalf("GoodPER should return false, but got true")
	}
}

func TestGoodPBV(t *testing.T) {
	vps := ValuePerShare(303619000000.00, 13044496930)
	price := PBV(26.74, vps)

	if !GoodPBV(price) {
		t.Fatalf("GoodPBV should return true, but got false")
	}

	vps = ValuePerShare(780032000.00, 87941972)
	price = PBV(19.07, vps)

	if GoodPBV(price) {
		t.Fatalf("GoodPBV should return false, but got true")
	}
}

func TestGoodProfitMargin(t *testing.T) {
	margin := ProfitMargin(31643000000.00, 302337000000.00)

	if !GoodProfitMargin(margin) {
		t.Fatalf("GoodProfitMargin should return true, but got false")
	}

	margin = ProfitMargin(29460000.00, 487220000.00)

	if GoodProfitMargin(margin) {
		t.Fatalf("GoodProfitMargin should return false, but got true")
	}
}

func TestGoodROE(t *testing.T) {
	roe := ROE(31643000000.00, 303619000000.00)

	if !GoodROE(roe) {
		t.Fatalf("GoodROE should return true, but got false")
	}

	roe = ROE(106850000.00, 4072017000.00)

	if GoodROE(roe) {
		t.Fatalf("GoodROE should return false, but got true")
	}
}
func TestGoodDebitRatio(t *testing.T) {
	debit := DebtRatio(108643000000.00, 303619000000.00)

	if !GoodDebitRatio(debit) {
		t.Fatalf("GoodDebitRatio should return true, but got false")
	}

	debit = DebtRatio(2341132000.00, 1743722000.00)

	if GoodDebitRatio(debit) {
		t.Fatalf("GoodDebitRatio should return false, but got true")
	}
}
func TestGoodDividendYield(t *testing.T) {
	dy := DividendYield(0.87, 26.74)

	if !GoodDividendYield(dy) {
		t.Fatalf("GoodDividendYield should return true, but got false")
	}

	dy = DividendYield(0.01, 7.85)

	if GoodDividendYield(dy) {
		t.Fatalf("GoodDividendYield should return false, but got true")
	}
}
