package fund

import (
	"errors"
	"math"
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
	Benchmark           string
	AdministrationFee   float64
	PerformanceFee      float64
	DividendYieldAnnual float64
	Shares              int
	Reports             []string
}

const PBV_MARK = 1.3
const PBV_NAME = "pbv"
const PBV_LABEL = "P/VPC (Preço / Valor Patrimonial da Cota)"

const DIVIDEND_YELD_MARK = 0.5
const DIVIDEND_YELD_NAME = "dividendYield"
const DIVIDEND_YELD_LABEL = "Dividend Yield do Período (Rendimentos por Cota / Preço da Cota)"

const ADMINISTRATION_FEE_MARK = 1.50
const ADMINISTRATION_FEE_NAME = "administrationFee"
const ADMINISTRATION_FEE_LABEL = "Taxa de Administração"

var ErrFundIDInvalid = errors.New("fund ID is invalid")
var ErrFundAdminstratorInvalid = errors.New("fund administrator is invalid")
var ErrFundLastIncomeInvalid = errors.New("fund last income is invalid")
var ErrFundNetEquityInvalid = errors.New("fund net equity is invalid")
var ErrFundPriceInvalid = errors.New("fund price is invalid")
var ErrFundAdministrationFeeInvalid = errors.New("fund administration fee is invalid")
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
	if entity.AdministrationFee == 0 {
		return ErrFundAdministrationFeeInvalid
	}
	if entity.Shares == 0 {
		return ErrFundSharesInvalid
	}

	return nil
}

func (entity *FundEntity) GetPBV() domain.Indicator {
	eps := math.Floor((entity.NetEquity/float64(entity.Shares))*100) / 100
	pbv := PBV(entity.Price, eps)

	return domain.Indicator{
		Name:  PBV_NAME,
		Label: PBV_LABEL,
		Mark:  PBV_MARK,
		Value: pbv,
		Good:  GoodPBV(pbv),
	}
}
func (entity *FundEntity) GetDividenYield() domain.Indicator {
	dym := DividendYieldMonth(entity.LastIncome, entity.Price)
	return domain.Indicator{
		Name:  DIVIDEND_YELD_NAME,
		Label: DIVIDEND_YELD_LABEL,
		Mark:  DIVIDEND_YELD_MARK,
		Value: dym,
		Good:  GoodDividendYield(dym),
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
	matched, err := regexp.MatchString("^[a-zA-Z]{4}(11|12|13|14|15)$", ID)

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
func GoodDividendYield(dy float64) bool {
	return dy > DIVIDEND_YELD_MARK
}
func GoodAdministrationFee(fee float64) bool {
	return fee < ADMINISTRATION_FEE_MARK
}
