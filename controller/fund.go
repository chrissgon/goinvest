package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/chrissgon/goinvest/app"
	"github.com/chrissgon/goinvest/domain"
	"github.com/chrissgon/goinvest/domain/fund"
	"github.com/chrissgon/goinvest/infra"
	"github.com/chrissgon/goinvest/internal/fiis"
)

type FundController struct{}

var fundPersist = infra.NewPersistMemory[fund.FundEntity]()
var fundSearchRepo = fiis.NewFiis()
var fundApp = app.NewFundApp(fundSearchRepo)

func (FundController) Search(ID string) (fund.FundEntity, error) {
	ID = strings.ToUpper(ID)

	fundCached := fundPersist.Get(ID)

	cachedIsValid := fundCached.ID == ID && time.Since(fundCached.CreatedAt) < 24*time.Hour

	if cachedIsValid {
		fmt.Println("cached valid")
		return fundCached, nil
	}
	
	fundGot, err := fundApp.Search(ID)
	
	if err == nil {
		fmt.Println("add persist")
		fundPersist.Add(fundGot.ID, fundGot)
	}

	return fundGot, err
}

func (FundController) Analyse(fundEntity fund.FundEntity) (map[string]domain.Indicator, error) {
	return fundApp.Analyse(fundEntity)
}
