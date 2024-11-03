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

	expected := map[string]stock.StockIndicator{
		stock.PER_NAME: {
			Name:  stock.PER_NAME,
			Label: stock.PER_LABEL,
			Mark:  stock.PER_MARK,
			Value: 0.9801580900145184,
			Good:  true,
		},
		stock.PBV_NAME: {
			Name:  stock.PBV_NAME,
			Label: stock.PBV_LABEL,
			Mark:  stock.PBV_MARK,
			Value: 0.23246738340283887,
			Good:  true,
		},
		stock.PROFIT_MARGIN_NAME: {
			Name:  stock.PROFIT_MARGIN_NAME,
			Label: stock.PROFIT_MARGIN_LABEL,
			Mark:  stock.PROFIT_MARGIN_MARK,
			Value: 23.194821267076016,
			Good:  true,
		},
		stock.ROE_NAME: {
			Name:  stock.ROE_NAME,
			Label: stock.ROE_LABEL,
			Mark:  stock.ROE_MARK,
			Value: 23.718666342175712,
			Good:  true,
		},
		stock.DEBIT_RATIO_NAME: {
			Name:  stock.DEBIT_RATIO_NAME,
			Label: stock.DEBIT_RATIO_LABEL,
			Mark:  stock.DEBIT_RATIO_MARK,
			Value: 23.24653200292042,
			Good:  true,
		},
	}

	if !reflect.DeepEqual(indicators, expected) {
		t.Fatalf("Analyse returned unexpected indicators. Expected %v, but got %v", expected, indicators)
	}
}
