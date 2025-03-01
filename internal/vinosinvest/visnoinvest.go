package vinosinvest

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/chrissgon/goinvest/domain/stock"
	"github.com/chrissgon/goinvest/internal"
)

const url = "https://api.visnoinvest.com.br/stocks/details/"

type VisnoInvest struct{}

type visnoInvestResponse struct {
	Metadata struct {
		Company string
		Price   string
	} `json:"metadata"`
	MetricGroups []struct {
		Key     string
		Metrics []struct {
			Key   string
			Value string
		}
	} `json:"metric_groups"`
}

func NewVisnoInvest() stock.StockSearchRepo {
	return &VisnoInvest{}
}

func (v *VisnoInvest) Run(ID string) (stock.StockEntity, error) {
	ID = strings.ToUpper(ID)

	res, err := http.Get(url + ID)

	if err != nil {
		return stock.StockEntity{}, err
	}

	if res.StatusCode != http.StatusOK {
		return stock.StockEntity{}, stock.ErrStockNotFound
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return stock.StockEntity{}, err
	}

	data := visnoInvestResponse{}
	json.Unmarshal(body, &data)

	price, err := internal.ConvertStringToFloat64(data.Metadata.Price)

	if err != nil {
		return stock.StockEntity{}, err
	}

	netProfit, err := internal.ConvertStringToFloat64(data.MetricGroups[3].Metrics[4].Value)

	if err != nil {
		return stock.StockEntity{}, err
	}

	netRevenue, err := internal.ConvertStringToFloat64(data.MetricGroups[3].Metrics[0].Value)

	if err != nil {
		return stock.StockEntity{}, err
	}

	netEquity, err := internal.ConvertStringToFloat64(data.MetricGroups[4].Metrics[1].Value)

	if err != nil {
		return stock.StockEntity{}, err
	}

	netDebt, err := internal.ConvertStringToFloat64(data.MetricGroups[4].Metrics[3].Value)

	if err != nil {
		return stock.StockEntity{}, err
	}

	dividendYield, err := internal.ConvertStringToFloat64(data.MetricGroups[1].Metrics[13].Value)

	if err != nil {
		return stock.StockEntity{}, err
	}

	lpa, err := internal.ConvertStringToFloat64(data.MetricGroups[1].Metrics[7].Value)

	if err != nil {
		return stock.StockEntity{}, err
	}

	return stock.StockEntity{
		ID:            ID,
		Price:         price,
		Company:       data.Metadata.Company,
		NetProfit:     netProfit,
		NetRevenue:    netRevenue,
		NetEquity:     netEquity,
		NetDebt:       netDebt,
		Shares:        int(netProfit / lpa),
		DividendYield: dividendYield,
		CreatedAt:     time.Now(),
	}, nil
}
