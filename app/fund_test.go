package app

import (
	"errors"
	"reflect"
	"testing"

	"github.com/chrissgon/goinvest/domain/fund"
)

func TestApp_FundSearch(t *testing.T) {
	app := NewFundApp(fund.NewFundSearchRepoMock())

	_, err := app.Search("")

	if !errors.Is(err, fund.ErrFundIDInvalid) {
		t.Fatalf("Search should return an error because ID is invalid")
	}

	fundEntity, _ := app.Search("MXRF11")

	if !reflect.DeepEqual(fundEntity, fund.FundEntityMockMXRF11) {
		t.Fatalf("Search should return a fund")
	}
}

func TestApp_FundAnalyse(t *testing.T) {
	app := NewFundApp(fund.NewFundSearchRepoMock())

	_, err := app.Analyse(fund.FundEntity{})

	if err == nil {
		t.Fatalf("Analyse should return an error because fund entity is not valid")
	}

	indicators, err := app.Analyse(fund.FundEntityMockMXRF11)

	if err != nil {
		t.Fatalf("Analyse should not return an error because fund entity is valid")
	}

	if indicators == nil {
		t.Fatalf("Analyse should return the indicators")
	}

	if !reflect.DeepEqual(indicators, fund.FundIndicatorsMockMXRF11) {
		t.Fatalf("Analyse returned unexpected indicators. Expected %v, but got %v", fund.FundIndicatorsMockMXRF11, indicators)
	}
}
