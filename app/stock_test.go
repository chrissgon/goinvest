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

	if !reflect.DeepEqual(*stockEntity, stock.StockEntityMockVALE3) {
		t.Fatalf("Search should return a stock")
	}
}

// func TestApp_StockAnalyse(t *testing.T) {
// 	app := NewStockApp(stock.NewStockSearchRepoMock())

// 	_, err := app.Analyse()

// 	if !errors.Is(err, stock.ErrStockIDInvalid) {
// 		t.Fatalf("Search should return an error because ID is invalid")
// 	}

// 	stockEntity, _ := app.Search("VALE3")

// 	if !reflect.DeepEqual(*stockEntity, stock.StockEntityMockVALE3) {
// 		t.Fatalf("Search should return a stock")
// 	}
// }
