package vinosinvest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/chrissgon/goinvest/domain"
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

func NewVisnoInvest() domain.StockSearchRepo {
	return &VisnoInvest{}
}

func (v *VisnoInvest) Run(ID string) (*domain.StockEntity, error) {
	res, err := http.Get(url + ID)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, domain.ErrStockNotFound
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	data := visnoInvestResponse{}
	json.Unmarshal(body, &data)
	fmt.Println(data.Metadata.Price)

	price, err := convertStringToFloat64(data.Metadata.Price)

	if err != nil {
		return nil, err
	}

	netProfit, err := convertStringToFloat64(data.MetricGroups[1].Metrics[6].Value)

	if err != nil {
		return nil, err
	}

	netRevenue, err := convertStringToFloat64(data.MetricGroups[1].Metrics[4].Value)

	if err != nil {
		return nil, err
	}

	netEquity, err := convertStringToFloat64(data.MetricGroups[1].Metrics[10].Value)

	if err != nil {
		return nil, err
	}

	netDebt, err := convertStringToFloat64(data.MetricGroups[1].Metrics[17].Value)

	if err != nil {
		return nil, err
	}

	marketCap, err := convertStringToFloat64(data.MetricGroups[1].Metrics[17].Value)

	if err != nil {
		return nil, err
	}

	return &domain.StockEntity{
		Price:      price,
		Company:    data.Metadata.Company,
		NetProfit:  netProfit,
		NetRevenue: netRevenue,
		NetEquity:  netEquity,
		NetDebt:    netDebt,
		Shares:     int(marketCap / price),
	}, nil
}

func convertStringToFloat64(s string) (float64, error) {
	// Remove currency symbol and any whitespace
	s = strings.TrimSpace(strings.TrimPrefix(s, "R$"))

	// Remove thousands separators
	s = strings.ReplaceAll(s, ".", "")

	// Replace comma with dot for decimal point
	s = strings.ReplaceAll(s, ",", ".")

	// Extract the numeric part
	re := regexp.MustCompile(`[\d.]+`)
	numStr := re.FindString(s)

	// Convert to float64
	value, err := strconv.ParseFloat(numStr, 64)

	if err != nil {
		return 0, err
	}

	// Handle suffix (M for million, B for billion, etc.)
	if strings.Contains(s, "K") {
		value *= 1e3
	} else if strings.Contains(s, "M") {
		value *= 1e6
	} else if strings.Contains(s, "B") {
		value *= 1e9
	}

	return value, nil
}
