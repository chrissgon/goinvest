package controller

import (
	"time"

	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain"
	"github.com/chrissgon/goinvest/domain/fund"
	"github.com/chrissgon/goinvest/infra"
	"github.com/chrissgon/goinvest/internal/fiis"
)

type FundController struct{}

var persist = infra.NewPersistMemory[fund.FundEntity]()
var fundSearchRepo = fiis.NewFiis()
var fundApp = app.NewFundApp(fundSearchRepo)

func (FundController) Search(ID string) (fund.FundEntity, error) {
	fundCached := persist.Get(ID)

	if fundCached.ID != ID || fundCached.CreatedAt.After(fundCached.CreatedAt.Add(24*time.Hour)) {
		fundGot, err := fundApp.Search(ID)

		if err == nil {
			persist.Add(fundGot.ID, fundGot)
		}

		return fundGot, err
	}

	return fundCached, nil
}

func (FundController) Analyse(fundEntity fund.FundEntity) (map[string]domain.Indicator, error) {
	return fundApp.Analyse(fundEntity)
}
