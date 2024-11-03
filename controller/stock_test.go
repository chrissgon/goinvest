package controller

import (
	"reflect"
	"testing"

	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain/stock"
)

func TestController_StockSearch(t *testing.T) {
	controller := newStockControllerMock()

	stockEntity, err := controller.Search("VALE3")

	if err != nil {
		t.Fatalf("should not return an error because ID stock is valid")
	}

	if !reflect.DeepEqual(stockEntity, stock.StockEntityMockVALE3) {
		t.Fatalf("Search should return a stock")
	}
}

func TestController_StockAnalyse(t *testing.T) {
	controller := newStockControllerMock()

	_, err := controller.Analyse(stock.StockEntity{})

	if err == nil {
		t.Fatalf("Analyse should return an error because stock entity is not valid")
	}

	indicators, err := controller.Analyse(stock.StockEntityMockVALE3)

	if err != nil {
		t.Fatalf("Analyse should not return an error because stock entity is valid")
	}

	if indicators == nil {
		t.Fatalf("Analyse should return the indicators")
	}

	if !reflect.DeepEqual(indicators, stock.StockIndicatorsMockVALE3) {
		t.Fatalf("Analyse returned unexpected indicators. Expected %v, but got %v", stock.StockIndicatorsMockVALE3, indicators)
	}
}

func newStockControllerMock() *StockController {
	stockSearchRepo = stock.NewStockSearchRepoMock()
	stockApp = app.NewStockApp(stockSearchRepo)
	return &StockController{}
}
