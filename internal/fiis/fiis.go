package fiis

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/chrissgon/goinvest/ai"
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

		response, err := ai.Ask(`
		Me informe as taxas que voce pode identificar no texto indicado, bem como seus valores. Aqui está suas regras:

		1. Extraia as informações sobre taxas do texto indicado, e me apresente como um JSON.
		2. Deixe o JSON o mais simples possível, apenas com o nome da taxa e o valor numérico.
		3. Não me responda nada além do JSON.
		4. Caso não encontre nenhuma taxa retorne um JSON vazio.
		5. Não adicione a taxa ao JSON caso o valor seja zero ou null.
		7. A principal taxa a ser encontrada é 'Taxa de Administração'.
		8. Responda no formato do JSON abaixo:

		{
			NOME DA TAXA: VALOR NUMERICO DA TAXA,
		}

		Texto: ` + e.Text)

		if err != nil {
			log.Println(err)
			return
		}

		taxes := map[string]float64{}
		err = json.Unmarshal([]byte(response), &taxes)

		if err != nil {
			log.Println(err)
			return
		}

		fundEntity.Taxes = taxes
		fundEntity.AdministrationFee = taxes["Taxa de Administração"]
	})

	c.OnHTML(".updatesContent", func(e *colly.HTMLElement) {
	})

	c.OnError(func(_ *colly.Response, e error) {
		err = e
	})

	err = c.Visit(URL + ID)

	if err != nil {
		return fund.FundEntity{}, err
	}

	fundEntity.ID = ID
	fundEntity.CreatedAt = time.Now()

	return fundEntity, err
}

func getFundName(text string) string {
	nameRegex := regexp.MustCompile(`(?i)([a-zA-Z]{2,}.*?)\s\(RZAT11\)`)

	nameResp := nameRegex.FindAllStringSubmatch(text, -1)

	if len(nameResp) > 0 && len(nameResp[0]) > 1 {
		return nameResp[0][1]
	}

	return ""
}
