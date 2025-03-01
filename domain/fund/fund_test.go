package fund

import (
	"errors"
	"testing"

	"github.com/chrissgon/goinvest/domain"
)

func TestDomain_FundIsValid(t *testing.T) {
	fund := FundEntityMockMXRF11
	fund.ID = ""

	err := fund.IsValid()

	if !errors.Is(err, ErrFundIDInvalid) {
		t.Fatalf("fund IsValid should return the error: %v, but got %v", ErrFundIDInvalid, err)
	}

	fund = FundEntityMockMXRF11
	fund.Administrator = ""

	err = fund.IsValid()

	if !errors.Is(err, ErrFundAdminstratorInvalid) {
		t.Fatalf("fund IsValid should return the error: %v, but got %v", ErrFundAdminstratorInvalid, err)
	}

	fund = FundEntityMockMXRF11
	fund.NetEquity = 0

	err = fund.IsValid()

	if !errors.Is(err, ErrFundNetEquityInvalid) {
		t.Fatalf("fund IsValid should return the error: %v, but got %v", ErrFundNetEquityInvalid, err)
	}

	fund = FundEntityMockMXRF11
	fund.Price = 0

	err = fund.IsValid()

	if !errors.Is(err, ErrFundPriceInvalid) {
		t.Fatalf("fund IsValid should return the error: %v, but got %v", ErrFundPriceInvalid, err)
	}

	fund = FundEntityMockMXRF11
	fund.Shares = 0

	err = fund.IsValid()

	if !errors.Is(err, ErrFundSharesInvalid) {
		t.Fatalf("fund IsValid should return the error: %v, but got %v", ErrFundPriceInvalid, err)
	}

	fund = FundEntityMockMXRF11

	err = fund.IsValid()

	if err != nil {
		t.Fatalf("fund IsValid should return nil because is valid, but got %v", err)
	}
}

func TestDomain_FundGetPBV(t *testing.T) {
	fund := FundEntityMockMXRF11

	pbv := fund.GetPBV()

	if pbv.Name != PBV_NAME {
		t.Fatalf("GetPBV should return a struct with Name %v, but got %v", PBV_NAME, pbv.Name)
	}

	if pbv.Mark != PBV_MARK {
		t.Fatalf("GetPBV should return a struct with Mark %v, but got %v", PBV_MARK, pbv.Mark)
	}

	value := 1.0567951318458417
	if pbv.Value != value {
		t.Fatalf("GetPBV should return a struct with Value %v, but got %v", value, pbv.Value)
	}

	if !pbv.Good {
		t.Fatalf("GetPBV Good property should be true, but got false")
	}

	fund = FundEntityMockIRDM11

	pbv = fund.GetPBV()

	value = 1.374493927125506
	if pbv.Value != value {
		t.Fatalf("GetPBV should return a struct with Value %v, but got %v", value, pbv.Value)
	}

	if pbv.Good {
		t.Fatalf("GetPBV Good property should be false, but got true")
	}
}

func TestDomain_FundGetDividenYieldMonth(t *testing.T) {
	fund := FundEntityMockMXRF11

	dividend := fund.GetDividenYieldMonth()

	if dividend.Name != DIVIDEND_YIELD_MONTH_NAME {
		t.Fatalf("GetDividenYieldMonth should return a struct with Name %v, but got %v", DIVIDEND_YIELD_MONTH_NAME, dividend.Name)
	}

	if dividend.Mark != DIVIDEND_YIELD_MARK {
		t.Fatalf("GetDividenYieldMonth should return a struct with Mark %v, but got %v", DIVIDEND_YIELD_MARK, dividend.Mark)
	}

	value := 0.9596928982725529
	if dividend.Value != value {
		t.Fatalf("GetDividenYieldMonth should return a struct with Value %v, but got %v", value, dividend.Value)
	}

	if !dividend.Good {
		t.Fatalf("GetDividenYieldMonth Good property should be true, but got false")
	}

	fund = FundEntityMockIRDM11

	dividend = fund.GetDividenYieldMonth()

	value = 0.49091801669121254
	if dividend.Value != value {
		t.Fatalf("GetDividenYieldMonth should return a struct with Value %v, but got %v", value, dividend.Value)
	}

	if dividend.Good {
		t.Fatalf("GetDividenYieldMonth Good property should be false, but got true")
	}
}

func TestDomain_FundGetAdministrationFee(t *testing.T) {
	fund := FundEntityMockMXRF11

	fee := fund.GetAdministrationFee()

	if fee.Name != ADMINISTRATION_FEE_NAME {
		t.Fatalf("GetAdministrationFee should return a struct with Name %v, but got %v", ADMINISTRATION_FEE_NAME, fee.Name)
	}

	if fee.Mark != ADMINISTRATION_FEE_MARK {
		t.Fatalf("GetAdministrationFee should return a struct with Mark %v, but got %v", ADMINISTRATION_FEE_MARK, fee.Mark)
	}

	value := 0.9
	if fee.Value != value {
		t.Fatalf("GetAdministrationFee should return a struct with Value %v, but got %v", value, fee.Value)
	}

	if !fee.Good {
		t.Fatalf("GetAdministrationFee Good property should be true, but got false")
	}

	fund = FundEntityMockIRDM11

	fee = fund.GetAdministrationFee()

	value = 1.7
	if fee.Value != value {
		t.Fatalf("GetAdministrationFee should return a struct with Value %v, but got %v", value, fee.Value)
	}

	if fee.Good {
		t.Fatalf("GetAdministrationFee Good property should be false, but got true")
	}
}

func TestDomain_FundCheckFundID(t *testing.T) {
	have := CheckFundID("MXRF11")

	if have != nil {
		t.Fatalf("CheckFundID should return nil because ID is valid")
	}

	have = CheckFundID("MXRF111")

	if !errors.Is(have, ErrFundIDInvalid) {
		t.Fatalf("CheckFundID should return an error because ID is invalid")
	}
}

func TestDomain_FundPBV(t *testing.T) {
	vps := domain.ValuePerShare(FundEntityMockMXRF11.NetEquity, FundEntityMockMXRF11.Shares)
	have := PBV(FundEntityMockMXRF11.Price, vps)
	expected := 1.0567951318458417

	if have != expected {
		t.Fatalf("PBV should return %v but got %v", expected, have)
	}
}

func TestDomain_FundDividenYieldMonth(t *testing.T) {
	have := DividendYieldMonth(FundEntityMockMXRF11.LastIncome, FundEntityMockMXRF11.Price)
	expected := 0.9596928982725529

	if have != expected {
		t.Fatalf("DividendYieldMonth should return %v but got %v", expected, have)
	}
}

func TestDomain_FundGoodPBV(t *testing.T) {
	vps := domain.ValuePerShare(FundEntityMockMXRF11.NetEquity, FundEntityMockMXRF11.Shares)
	pbv := PBV(FundEntityMockMXRF11.Price, vps)

	if !GoodPBV(pbv) {
		t.Fatalf("GoodPBV should return true, but got false")
	}

	vps = domain.ValuePerShare(FundEntityMockIRDM11.NetEquity, FundEntityMockIRDM11.Shares)
	pbv = PBV(FundEntityMockIRDM11.Price, vps)

	if GoodPBV(pbv) {
		t.Fatalf("GoodPBV should return false, but got true")
	}
}

func TestDomain_FundGoodDividendYieldMonth(t *testing.T) {
	dym := DividendYieldMonth(FundEntityMockMXRF11.LastIncome, FundEntityMockMXRF11.Price)

	if !GoodDividendYieldMonth(dym) {
		t.Fatalf("GoodDividendYieldMonth should return true, but got false")
	}

	dym = DividendYieldMonth(FundEntityMockIRDM11.LastIncome, FundEntityMockIRDM11.Price)

	if GoodDividendYieldMonth(dym) {
		t.Fatalf("GoodDividendYieldMonth should return false, but got true")
	}
}

func TestDomain_FundGoodAdministrationFee(t *testing.T) {
	if !GoodAdministrationFee(FundEntityMockMXRF11.AdministrationFee) {
		t.Fatalf("GoodAdministrationFee should return true, but got false")
	}

	if GoodAdministrationFee(FundEntityMockIRDM11.AdministrationFee) {
		t.Fatalf("GoodAdministrationFee should return false, but got true")
	}
}
