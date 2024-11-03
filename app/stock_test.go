package app

import (
	"errors"
	"reflect"
	"testing"

	"github.com/chrissgon/goinvest/domain/stock"
)

func TestApp_StockSearch(t *testing.T) {
	app := NewStockApp(stock.NewStockSearchRepoMock())

	_, err := app.Search("")

	if !errors.Is(err, stock.ErrStockIDInvalid) {
		t.Fatalf("Search should return an error because ID is invalid")
	}

	stockEntity, _ := app.Search("VALE3")

	if !reflect.DeepEqual(stockEntity, stock.StockEntityMockVALE3) {
		t.Fatalf("Search should return a stock")
	}
}

func TestApp_StockAnalyse(t *testing.T) {
	app := NewStockApp(stock.NewStockSearchRepoMock())

	_, err := app.Analyse(stock.StockEntity{})

	if err == nil {
		t.Fatalf("Analyse should return an error because stock entity is not valid")
	}

	indicators, err := app.Analyse(stock.StockEntityMockVALE3)

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
