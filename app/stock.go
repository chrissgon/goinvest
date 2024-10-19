package app

import "github.com/chrissgon/goinvest/domain"

type StockApp struct {
	searchRepo domain.StockSearchRepo
}

func NewStockApp(searchRepo domain.StockSearchRepo) *StockApp {
	return &StockApp{searchRepo}
}

func (app *StockApp) Search(ID string) (*domain.StockEntity, error) {
	err := domain.CheckStockID(ID)

	if err != nil {
		return nil, err
	}

	stock, err := app.searchRepo.Run(ID)
	stock.ID = ID

	return stock, err
}

func (app *StockApp) Analyse(stock *domain.StockEntity) (map[string]*domain.StockIndicator, error) {
	err := stock.IsValid()

	if err != nil {
		return nil, err
	}

	indicators := map[string]*domain.StockIndicator{}

	per := stock.GetPER()
	indicators[per.Name] = per

	pbv := stock.GetPBV()
	indicators[pbv.Name] = pbv

	margin := stock.GetProfitMargin()
	indicators[margin.Name] = margin

	roe := stock.GetROE()
	indicators[roe.Name] = roe

	debt := stock.GetDebtRatio()
	indicators[debt.Name] = debt

	// dividend := stock.GetDividenYeld()
	// indicators[dividend.Name] = dividend

	return indicators, nil
}
