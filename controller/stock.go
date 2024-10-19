package controller

import (
	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain"
	"github.com/chrissgon/goinvest/internal/vinosinvest"
)

type StockController struct {}

var stockSearchRepo = vinosinvest.NewVisnoInvest()
var stockApp = app.NewStockApp(stockSearchRepo)

func (StockController) Search(ID string) (*domain.StockEntity, error){
	return stockApp.Search(ID)
}