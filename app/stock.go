package app

import "github.com/chrissgon/goinvest/domain"

type StockApp struct {
	searchRepo domain.StockSearchRepo
}

func NewStockApp(searchRepo domain.StockSearchRepo) *StockApp {
	return &StockApp{searchRepo}
}

func (app *StockApp) Search(ID string) (*domain.StockEntity, error){
	err := domain.CheckStockID(ID)

	if err != nil {
		return nil, err
	}
	
	return app.searchRepo.Run(ID)
}

func (app *StockApp) Analyse(stock *domain.StockEntity) {
	
}
