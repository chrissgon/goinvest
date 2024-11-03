package controller

import (
	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain"
	"github.com/chrissgon/goinvest/domain/fund"
	"github.com/chrissgon/goinvest/internal/fiis"
)

type FundController struct{}

var fundSearchRepo = fiis.NewFiis()
var fundApp = app.NewFundApp(fundSearchRepo)

func (FundController) Search(ID string) (fund.FundEntity, error) {
	return fundApp.Search(ID)
}

func (FundController) Analyse(fundEntity fund.FundEntity) (map[string]domain.Indicator, error) {
	return fundApp.Analyse(fundEntity)
}
