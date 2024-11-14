package fiis

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/chrissgon/goinvest/domain/fund"
	"github.com/chrissgon/goinvest/internal"
	"github.com/gocolly/colly"
)

type Fiis struct{}

const URL = "https://fiis.com.br/"

func NewFiis() fund.FundSearchRepo {
	return &Fiis{}
}

func (f *Fiis) Run(ID string) (fund.FundEntity, error) {
	var err error
	fundEntity := fund.FundEntity{}

	ID = strings.ToUpper(ID)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.OnHTML(".indicators__box", func(e *colly.HTMLElement) {
		title := strings.ToLower(internal.Normalization(e.ChildText("p:nth-child(2)")))
		value := e.ChildText("p>b")

		matched, _ := regexp.MatchString(`(?i)(dividend.*yield|yield.*dividend)`, title)
		if matched {
			fundEntity.DividendYieldAnnual, _ = internal.ConvertStringToFloat64(value)
			return
		}

		matched, _ = regexp.MatchString(`(?i)(patrimonio.*liquido|liquido.*patrimonio)`, title)
		if matched {
			fundEntity.NetEquity, _ = internal.ConvertStringToFloat64(value)
			return
		}

		matched, _ = regexp.MatchString(`(?i)(ultimo.*rendimento|rendimento.*ultimo)`, title)
		if matched {
			fundEntity.LastIncome, _ = internal.ConvertStringToFloat64(value)
			return
		}
	})

	c.OnHTML(".moreInfo.wrapper p", func(e *colly.HTMLElement) {
		title := strings.ToLower(internal.Normalization(e.ChildText("span")))
		value := e.ChildText("b")

		matched, _ := regexp.MatchString(`(?i)(numero.*cotas|cotas.*numero)`, title)
		if matched {
			sharesFloat64, _ := internal.ConvertStringToFloat64(value)
			fundEntity.Shares = int(sharesFloat64)
		}
	})

	c.OnHTML(".item.quotation", func(e *colly.HTMLElement) {
		fundEntity.Price, _ = internal.ConvertStringToFloat64(e.ChildText(".value"))
	})

	c.OnHTML(".informations__adm__name", func(e *colly.HTMLElement) {
		fundEntity.Administrator = e.Text
	})

	c.OnHTML(".newsContent__article", func(e *colly.HTMLElement) {
		fundEntity.Name = getFundName(e.Text)
		fundEntity.AdministrationFee, fundEntity.PerformanceFee = getTaxes(e.Text)
	})

	c.OnHTML(".updatesContent", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnError(func(_ *colly.Response, e error) {
		err = e
	})

	err = c.Visit(URL + ID)

	if err != nil {
		return fund.FundEntity{}, err
	}

	fundEntity.ID = ID

	return fundEntity, err
}

func getTaxes(text string) (admFee float64, prfFee float64) {
	admRegex := regexp.MustCompile(`(?i)taxa de (administração).*?(\d+[,\.]?\d*)`)
	prfRegex := regexp.MustCompile(`(?i)taxa de (performance).*?(\d+[,\.]?\d*)`)

	admResp := admRegex.FindAllStringSubmatch(text, -1)
	prfResp := prfRegex.FindAllStringSubmatch(text, -1)

	if len(admResp) > 0 && len(admResp[0]) > 2 {
		admFee, _ = internal.ConvertTaxStringToFloat64(admResp[0][2])
	}
	if len(prfResp) > 0 && len(prfResp[0]) > 2 {
		prfFee, _ = internal.ConvertTaxStringToFloat64(prfResp[0][2])
	}

	return
}

func getFundName(text string) string {
	nameRegex := regexp.MustCompile(`(?i)([a-zA-Z]{2,}.*?)\s\(RZAT11\)`)

	nameResp := nameRegex.FindAllStringSubmatch(text, -1)

	if len(nameResp) > 0 && len(nameResp[0]) > 1 {
		return nameResp[0][1]
	}

	return ""
}
