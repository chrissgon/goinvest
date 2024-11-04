package app

import (
	"strings"

	"github.com/chrissgon/goinvest/domain"
	"github.com/chrissgon/goinvest/domain/fund"
)

type FundApp struct {
	searchRepo fund.FundSearchRepo
}

func NewFundApp(searchRepo fund.FundSearchRepo) FundApp {
	return FundApp{searchRepo}
}

func (app *FundApp) Search(ID string) (fund.FundEntity, error) {
	ID = strings.ToUpper(ID)

	err := fund.CheckFundID(ID)

	if err != nil {
		return fund.FundEntity{}, err
	}

	fundEntity, err := app.searchRepo.Run(ID)

	fundEntity.ID = ID

	return fundEntity, err
}

func (app *FundApp) Analyse(fundEntity fund.FundEntity) (map[string]domain.Indicator, error) {
	err := fundEntity.IsValid()

	if err != nil {
		return nil, err
	}

	indicators := map[string]domain.Indicator{}

	pbv := fundEntity.GetPBV()
	indicators[pbv.Name] = pbv

	dym := fundEntity.GetDividenYieldMonth()
	indicators[dym.Name] = dym

	administrationFee := fundEntity.GetAdministrationFee()
	indicators[administrationFee.Name] = administrationFee

	return indicators, nil
}
