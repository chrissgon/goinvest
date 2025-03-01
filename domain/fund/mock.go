package fund

import "github.com/chrissgon/goinvest/domain"

type fundSearchRepoMock struct{}

var FundEntityMockMXRF11 = FundEntity{
	ID:                  "MXRF11",
	Name:                "Maxi Renda Fundo de Investimento Imobiliário - FII",
	Administrator:       "BTG PACTUAL SERVIÇOS FINANCEIROS S/A DTVM",
	LastIncome:          0.10,
	NetEquity:           3300000000,
	Price:               10.42,
	AdministrationFee:   0.90,
	DividendYieldAnnual: 12.76,
	Shares:              334655892,
	Benchmark:           "",
	Reports:             []string{},
}

var FundEntityMockIRDM11 = FundEntity{
	ID:                  "IRDM11",
	Name:                "Iridium Recebíveis Imobiliários",
	Administrator:       "BTG PACTUAL SERVIÇOS FINANCEIROS S/A DTVM",
	LastIncome:          0.5,
	NetEquity:           2700000000,
	Price:               101.85,
	AdministrationFee:   1.7,
	DividendYieldAnnual: 13.49,
	Shares:              36433827,
	Benchmark:           "",
	Reports:             []string{},
}

var FundIndicatorsMockMXRF11 = map[string]domain.Indicator{
	PBV_NAME:{
		Name: PBV_NAME,
		Label: PBV_LABEL,
		Mark: PBV_MARK,
		Value: 1.0567951318458417,
		Good: true,
	},
	DIVIDEND_YELD_MONTH_NAME: {
		Name: DIVIDEND_YELD_MONTH_NAME,
		Label: DIVIDEND_YELD_MONTH_LABEL,
		Mark: DIVIDEND_YELD_MARK,
		Value: 0.9596928982725529,
		Good: true,
	},
	ADMINISTRATION_FEE_NAME: {
		Name: ADMINISTRATION_FEE_NAME,
		Label: ADMINISTRATION_FEE_LABEL,
		Mark: ADMINISTRATION_FEE_MARK,
		Value: 0.9,
		Good: true,
	},
}

func NewFundSearchRepoMock() FundSearchRepo {
	return &fundSearchRepoMock{}
}

func (v *fundSearchRepoMock) Run(ID string) (FundEntity, error) {
	return FundEntityMockMXRF11, nil
}
