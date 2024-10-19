package controller

import (
	"reflect"
	"testing"

	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/mock"
)

func TestController_StockSearch(t *testing.T) {
	controller := newStockControllerMock()

	stock, err := controller.Search("PETR4")

	if err != nil {
		t.Fatalf("should not return an error because ID stock is valid")
	}

	if !reflect.DeepEqual(*stock, mock.StockEntityMockPETR4) {
		t.Fatalf("Search should return a stock")
	}
}

func newStockControllerMock() *StockController {
	stockSearchRepo = mock.NewStockSearchRepoMock()
	stockApp = app.NewStockApp(stockSearchRepo)
	return &StockController{}
}
