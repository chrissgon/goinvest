package controller

import (
	"reflect"
	"testing"

	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain"
)

func TestController_StockSearch(t *testing.T) {
	controller := newStockControllerMock()

	stock, err := controller.Search("VALE3")

	if err != nil {
		t.Fatalf("should not return an error because ID stock is valid")
	}

	if !reflect.DeepEqual(*stock, domain.StockEntityMockVALE3) {
		t.Fatalf("Search should return a stock")
	}
}

func newStockControllerMock() *StockController {
	stockSearchRepo = domain.NewStockSearchRepoMock()
	stockApp = app.NewStockApp(stockSearchRepo)
	return &StockController{}
}
