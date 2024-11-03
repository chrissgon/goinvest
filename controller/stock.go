package controller

import (
	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain/stock"
	"github.com/chrissgon/goinvest/internal/vinosinvest"
)

type StockController struct{}

var stockSearchRepo = vinosinvest.NewVisnoInvest()
var stockApp = app.NewStockApp(stockSearchRepo)

func (StockController) Search(ID string) (stock.StockEntity, error) {
	return stockApp.Search(ID)
}

func (StockController) Analyse(stockEntity stock.StockEntity) (map[string]*stock.StockIndicator, error) {
	return stockApp.Analyse(stockEntity)
}
