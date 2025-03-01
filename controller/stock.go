package controller

import (
	"strings"
	"time"

	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain"
	"github.com/chrissgon/goinvest/domain/stock"
	"github.com/chrissgon/goinvest/infra"
	"github.com/chrissgon/goinvest/internal/vinosinvest"
)

type StockController struct{}

var stockPersist = infra.NewPersistMemory[stock.StockEntity]()
var stockSearchRepo = vinosinvest.NewVisnoInvest()
var stockApp = app.NewStockApp(stockSearchRepo)

func (StockController) Search(ID string) (stock.StockEntity, error) {
	ID = strings.ToUpper(ID)

	stockCached := stockPersist.Get(ID)

	cachedIsValid := stockCached.ID == ID && time.Since(stockCached.CreatedAt) < 24*time.Hour

	if cachedIsValid {
		return stockCached, nil
	}

	stockGot, err := stockApp.Search(ID)

	if err != nil {
		stockPersist.Add(stockGot.ID, stockGot)
	}

	return stockGot, err
}

func (StockController) Analyse(stockEntity stock.StockEntity) (map[string]domain.Indicator, error) {
	return stockApp.Analyse(stockEntity)
}
