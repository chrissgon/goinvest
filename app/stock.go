package app

import "github.com/chrissgon/goinvest/domain/stock"

type StockApp struct {
	searchRepo stock.StockSearchRepo
}

func NewStockApp(searchRepo stock.StockSearchRepo) *StockApp {
	return &StockApp{searchRepo}
}

func (app *StockApp) Search(ID string) (stock.StockEntity, error) {
	err := stock.CheckStockID(ID)

	if err != nil {
		return stock.StockEntity{}, err
	}

	stockEntity, err := app.searchRepo.Run(ID)

	stockEntity.ID = ID

	return stockEntity, err
}

func (app *StockApp) Analyse(stockEntity stock.StockEntity) (map[string]*stock.StockIndicator, error) {
	err := stockEntity.IsValid()

	if err != nil {
		return nil, err
	}

	indicators := map[string]*stock.StockIndicator{}

	per := stockEntity.GetPER()
	indicators[per.Name] = per

	pbv := stockEntity.GetPBV()
	indicators[pbv.Name] = pbv

	margin := stockEntity.GetProfitMargin()
	indicators[margin.Name] = margin

	roe := stockEntity.GetROE()
	indicators[roe.Name] = roe

	debt := stockEntity.GetDebtRatio()
	indicators[debt.Name] = debt

	// dividend := stockEntity.GetDividenYeld()
	// indicators[dividend.Name] = dividend

	return indicators, nil
}
