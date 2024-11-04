package controller

import (
	"reflect"
	"testing"

	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain/fund"
)

func TestController_FundSearch(t *testing.T) {
	controller := newFundControllerMock()

	fundEntity, err := controller.Search("MXRF11")

	if err != nil {
		t.Fatalf("should not return an error because ID stock is valid")
	}

	if !reflect.DeepEqual(fundEntity, fund.FundEntityMockMXRF11) {
		t.Fatalf("should return a stock")
	}
}
func TestController_FundAnalyse(t *testing.T) {
	controller := newFundControllerMock()

	_, err := controller.Analyse(fund.FundEntity{})

	if err == nil {
		t.Fatalf("should return an error because fund entity is not valid")
	}

	indicators, err := controller.Analyse(fund.FundEntityMockMXRF11)

	if err != nil {
		t.Fatalf("should not return an error because fund entity is valid")
	}

	if indicators == nil {
		t.Fatalf("should return the indicators")
	}

	if !reflect.DeepEqual(indicators, fund.FundIndicatorsMockMXRF11) {
		t.Fatalf("returned unexpected indicators. Expected %v, but got %v", fund.FundIndicatorsMockMXRF11, indicators)
	}
}

func newFundControllerMock() *FundController {
	fundSearchRepo = fund.NewFundSearchRepoMock()
	fundApp = app.NewFundApp(fundSearchRepo)
	return &FundController{}
}
