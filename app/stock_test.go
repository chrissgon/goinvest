package app

import (
	"errors"
	"reflect"
	"testing"

	"github.com/chrissgon/goinvest/domain"
)

func TestApp_StockSearch(t *testing.T) {
	app := NewStockApp(domain.NewStockSearchRepoMock())

	_, err := app.Search("")

	if !errors.Is(err, domain.ErrStockIDInvalid) {
		t.Fatalf("Search should return an error because ID is invalid")
	}

	stock, _ := app.Search("VALE3")

	if !reflect.DeepEqual(*stock, domain.StockEntityMockVALE3) {
		t.Fatalf("Search should return a stock")
	}
}
