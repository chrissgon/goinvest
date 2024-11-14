package fund

import (
	"errors"
	"regexp"

	"github.com/chrissgon/goinvest/domain"
)

type FundEntity struct {
	ID                  string
	Name                string
	Administrator       string
	LastIncome          float64
	NetEquity           float64
	Price               float64
	AdministrationFee   float64
	PerformanceFee      float64
	DividendYieldAnnual float64
	Shares              int
	Benchmark           string
	Reports             []string
}

const PBV_MARK = 1.3
const PBV_NAME = "pbv"
const PBV_LABEL = "P/VPC (Preço / Valor Patrimonial da Cota)"

const DIVIDEND_YELD_MARK = 0.5
const DIVIDEND_YELD_MONTH_NAME = "dividendYieldMonth"
const DIVIDEND_YELD_MONTH_LABEL = "Dividend Yield do Período (Rendimentos por Cota / Preço da Cota)"

const ADMINISTRATION_FEE_MARK = 1.50
const ADMINISTRATION_FEE_NAME = "administrationFee"
const ADMINISTRATION_FEE_LABEL = "Taxa de Administração"

var ErrFundIDInvalid = errors.New("fund ID is invalid")
var ErrFundAdminstratorInvalid = errors.New("fund administrator is invalid")
var ErrFundLastIncomeInvalid = errors.New("fund last income is invalid")
var ErrFundNetEquityInvalid = errors.New("fund net equity is invalid")
var ErrFundPriceInvalid = errors.New("fund price is invalid")
var ErrFundDividendYieldInvalid = errors.New("fund dividend yield is invalid")
var ErrFundSharesInvalid = errors.New("fund shares is invalid")

func (entity *FundEntity) IsValid() error {
	err := CheckFundID(entity.ID)

	if err != nil {
		return err
	}
	if entity.Administrator == "" {
		return ErrFundAdminstratorInvalid
	}
	if entity.NetEquity == 0 {
		return ErrFundNetEquityInvalid
	}
	if entity.Price == 0 {
		return ErrFundPriceInvalid
	}
	if entity.Shares == 0 {
		return ErrFundSharesInvalid
	}

	return nil
}

func (entity *FundEntity) GetPBV() domain.Indicator {
	vps := domain.ValuePerShare(entity.NetEquity, entity.Shares)
	pbv := PBV(entity.Price, vps)

	return domain.Indicator{
		Name:  PBV_NAME,
		Label: PBV_LABEL,
		Mark:  PBV_MARK,
		Value: pbv,
		Good:  GoodPBV(pbv),
	}
}
func (entity *FundEntity) GetDividenYieldMonth() domain.Indicator {
	dym := DividendYieldMonth(entity.LastIncome, entity.Price)
	return domain.Indicator{
		Name:  DIVIDEND_YELD_MONTH_NAME,
		Label: DIVIDEND_YELD_MONTH_LABEL,
		Mark:  DIVIDEND_YELD_MARK,
		Value: dym,
		Good:  GoodDividendYieldMonth(dym),
	}
}
func (entity *FundEntity) GetAdministrationFee() domain.Indicator {
	return domain.Indicator{
		Name:  ADMINISTRATION_FEE_NAME,
		Label: ADMINISTRATION_FEE_LABEL,
		Mark:  ADMINISTRATION_FEE_MARK,
		Value: entity.AdministrationFee,
		Good:  GoodAdministrationFee(entity.AdministrationFee),
	}
}

func CheckFundID(ID string) error {
	matched, err := regexp.MatchString("^[a-zA-Z]{4}(11)$", ID)

	if err != nil {
		return err
	}

	if !matched {
		return ErrFundIDInvalid
	}

	return nil
}

func PBV(fundPrice, eps float64) float64 {
	return fundPrice / eps
}
func DividendYieldMonth(lastIncome, fundPrice float64) float64 {
	return lastIncome / fundPrice * 100
}

func GoodPBV(pricePerAsset float64) bool {
	return pricePerAsset < PBV_MARK
}
func GoodDividendYieldMonth(dym float64) bool {
	return dym > DIVIDEND_YELD_MARK
}
func GoodAdministrationFee(fee float64) bool {
	return fee < ADMINISTRATION_FEE_MARK
}
