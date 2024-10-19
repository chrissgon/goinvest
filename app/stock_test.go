package app

import (
	"errors"
	"reflect"
	"testing"

	"github.com/chrissgon/goinvest/domain"
	"github.com/chrissgon/goinvest/mock"
)

func TestApp_StockSearch(t *testing.T) {
	app := NewStockApp(mock.NewStockSearchRepoMock())

	_, err := app.Search("")

	if !errors.Is(err, domain.ErrStockInvalidID) {
		t.Fatalf("Search should return an error because ID is invalid")
	}

	stock, _ := app.Search("PETR4")

	if !reflect.DeepEqual(*stock, mock.StockEntityMockPETR4) {
		t.Fatalf("Search should return a stock")
	}
}
