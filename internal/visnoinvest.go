package internal

import (
	"encoding/json"
	"fmt"
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
	err := domain.CheckStockID(ID)

	if err != nil {
		return nil, err
	}

	// res, err := http.Get(url + ID)

	// if err != nil {
	// 	return nil, err
	// }

	// if res.StatusCode != http.StatusOK {
	// 	return nil, domain.ErrStockNotFound
	// }

	// body, err := io.ReadAll(res.Body)

	// if err != nil {
	// 	return nil, err
	// }

	body := []byte(`
	{
  "metric_groups": [
    {
      "key": "general",
      "text": "Informações gerais",
      "metrics": [
        {
          "key": "company",
          "text": "Empresa",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general", "extra"],
          "table": {
            "show_by_default": false,
            "show_in_filters_by_column": false
          },
          "value": "ISA CTEEP - Transmissão Paulista"
        },
        {
          "key": "activity",
          "text": "Atividade",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": false,
          "value": "Transmissão de Energia Elétrica"
        },
        {
          "text": "Setor",
          "key": "sector",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": false
          },
          "value": "Utilidade Pública"
        },
        {
          "text": "Subsetor",
          "key": "subsector",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": false
          },
          "value": "Energia Elétrica"
        },
        {
          "text": "Segmento",
          "key": "segment",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": false
          },
          "value": "Energia Elétrica"
        },
        {
          "text": "Site",
          "key": "website",
          "format": "website_url",
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": false,
          "value": "isacteep.com.br"
        },
        {
          "text": "Situação",
          "key": "status_cvm",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": false,
          "value": "Fase Operacional"
        },
        {
          "text": "Anos na bolsa",
          "key": "years_listed",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": 25
        },
        {
          "text": "Código CVM",
          "key": "cd_cvm",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general", "extra"],
          "table": false,
          "value": 18376
        },
        {
          "text": "CNPJ",
          "key": "cnpj",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": false,
          "value": "02.998.611/0001-04"
        },
        {
          "text": "Dt. último preço",
          "key": "dt_last_price",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": false,
          "value": "17/10/2024"
        },
        {
          "text": "Dt. último DFP/ITR",
          "key": "dt_last_dfp_itr",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": false,
          "value": "30/06/2024"
        },
        {
          "text": "Free float",
          "key": "free_float",
          "format": "float_percentage_ptbr",
          "bank": true,
          "insurance": true,
          "groups": ["general"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_percentage_ptbr",
            "to_value_function": "ptbr_percentage_to_float"
          },
          "value": "64,18 %"
        },
        {
          "text": "Outros tickers",
          "key": "all_tickers",
          "format": false,
          "bank": true,
          "insurance": true,
          "groups": ["general", "extra"],
          "table": false,
          "value": "TRPL3,TRPL4"
        }
      ]
    },
    {
      "key": "financial",
      "text": "Informações financeiras",
      "metrics": [
        {
          "text": "Preço",
          "key": "price",
          "format": "human_readable_currency_ptbr",
          "enable_history": false,
          "history_granularity": "daily",
          "bank": true,
          "insurance": true,
          "description": "Último preço da ação no mercado.",
          "groups": ["financial", "extra"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 32,30"
        },
        {
          "text": "Market cap.",
          "key": "market_cap",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": true,
          "insurance": true,
          "description": "Market cap. é valor de mercado da empresa. <br/> Ele é calculado multiplicando-se o preço da ação pelo número total de ações em circulação.",
          "groups": ["financial"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 18,13 B"
        },
        {
          "text": "Volume diário",
          "key": "trading_volume",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": true,
          "insurance": true,
          "description": "Volume diário é o valor total das negociações de um único dia. <br/> Ele é calculado utilizando-se a média dos últimos três meses.",
          "groups": ["financial"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 79,01 K"
        },
        {
          "text": "Enterprise value (EV)",
          "key": "enterprise_value",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": false,
          "insurance": false,
          "description": "<i> Enterprise value </i> é o valor total da empresa. <br/> Ele é calculado somando-se o market cap. com a dívida líquida.",
          "groups": ["financial"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 26,96 B"
        },
        {
          "text": "Receita líquida",
          "key": "net_revenue",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "Receita líquida é o total de dinheiro que a empresa gerou através de suas atividades. <br/> Ela é calculada subtraindo-se determinados impostos, descontos, abatimentos e devoluções da receita bruta.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 6,81 B"
        },
        {
          "text": "Lucro bruto",
          "key": "gross_profit",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "Lucro bruto é o lucro total da empresa. <br/> Ele é calculado subtraindo-se os custos variáveis, ou seja, custo dos bens e serviços, da receita líquida.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 3,59 B"
        },
        {
          "text": "Lucro líquido",
          "key": "net_profit",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "Lucro líquido é o lucro da empresa após todos os custos. <br/> Ele é calculado subtraindo-se os custos fixos, como salários, aluguéis e depreciação de ativos, do lucro bruto da empresa.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 2,67 B"
        },
        {
          "text": "LPA",
          "key": "lpa",
          "format": "human_readable_currency_ptbr",
          "enable_history": false,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "LPA ou Lucro líquido por ação é o lucro da empresa atribuível a cada ação. <br/> Ele é calculado dividindo-se o lucro líquido da empresa pelo número total de ações.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 4,05"
        },
        {
          "text": "EBIT",
          "key": "ebit",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "EBIT é o lucro da empresa antes de juros e impostos. <br/> Ele é calculado somando-se os juros e impostos pagos ao lucro líquido da empresa.",
          "groups": ["financial"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 3,74 B"
        },
        {
          "text": "EBITDA",
          "key": "ebitda",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "EBITDA é o lucro da empresa antes de juros, impostos, depreciação e amortização. <br/> Ele é calculado somando-se os juros, impostos, depreciação e amortização pagos ao lucro líquido da empresa.",
          "groups": ["financial"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 3,77 B"
        },
        {
          "text": "Patrimônio líquido",
          "key": "net_worth",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "Patrimônio líquido é o valor que os acionistas possuem na empresa. <br/> Ele é calculado subtraindo-se os passivos dos ativos da empresa.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 18,60 B"
        },
        {
          "text": "VPA",
          "key": "vpa",
          "format": "human_readable_currency_ptbr",
          "enable_history": false,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "VPA ou Patrimônio líquido por ação é o valor contábil da empresa atribuível a cada ação. <br/> Ele é calculado dividindo-se o patrimônio líquido da empresa pelo número total de ações.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 28,24"
        },
        {
          "text": "Ativos",
          "key": "assets",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "Ativos são tudo o que a empresa possui de valor, como imóveis, máquinas, equipamentos, dinheiro em caixa e contas a receber.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 38,56 B"
        },
        {
          "text": "Dividend Yield",
          "key": "dy",
          "format": "float_percentage_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": true,
          "insurance": true,
          "description": "Dividend Yield é a proporção entre os dividendos pagos aos acionistas e o preço ação. <br/> Ele é calculado dividindo-se a soma dos dividendos pagos nos últimos 12 meses pelo preço atual da ação.",
          "groups": ["financial", "extra"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_percentage_ptbr",
            "to_value_function": "ptbr_percentage_to_float"
          },
          "value": "6,82 %"
        },
        {
          "text": "Dívida de curto prazo",
          "key": "short_term_debt",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Dívida de curto prazo são as obrigações financeiras da empresa que vencem em um prazo inferior a 12 meses.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 1,16 B"
        },
        {
          "text": "Dívida de longo prazo",
          "key": "long_term_debt",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Dívida de longo prazo são as obrigações financeiras da empresa que vencem em um prazo superior a 12 meses.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 10,29 B"
        },
        {
          "text": "Dívida bruta",
          "key": "gross_debt",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Dívida bruta é o valor total de todas as obrigações financeiras da empresa. <br/> Ela é calculada somando-se a dívida de curto prazo com a dívida de longo prazo.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 11,45 B"
        },
        {
          "text": "Dívida líquida",
          "key": "net_debt",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Dívida líquida é o valor remanescente das obrigações financeiras da empresa após deduzir os recursos disponíveis.<br/> Ela é calculada subtraindo-se o valor disponível em caixa e equivalentes de caixa da dívida bruta.",
          "groups": ["financial"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 8,83 B"
        }
      ]
    },
    {
      "key": "price",
      "text": "Indicadores de preço",
      "metrics": [
        {
          "text": "EV/EBIT",
          "key": "ev_ebit_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": false,
          "insurance": false,
          "description": "EV/EBIT é um indicador que relaciona o valor de uma empresa com seu lucro operacional. <br/> Ele é calculado dividindo-se o <i> enterprise value </i> pelo EBIT da empresa.",
          "groups": ["price"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "7,20"
        },
        {
          "text": "EV/EBITDA",
          "key": "ev_ebitda_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": false,
          "insurance": false,
          "description": "EV/EBITDA é um indicador que relaciona o valor de uma empresa com ao seu lucro antes de juros, impostos, depreciação e amortização. <br/> Ele é calculado dividindo-se o <i> enterprise value </i>pelo EBITDA da empresa.",
          "groups": ["price"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "7,14"
        },
        {
          "text": "EV/Receita Líquida",
          "key": "ev_net_revenue_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": false,
          "insurance": false,
          "description": "EV/Receita Líquida é um indicador que relaciona o valor de uma empresa com sua receita líquida. <br/> Ele é calculado dividindo-se o <i> enterprise value </i> pela Receita Líquida da empresa.",
          "groups": ["price"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "3,96"
        },
        {
          "text": "EV/Ativos",
          "key": "ev_assets_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": false,
          "insurance": false,
          "description": "EV/Ativos é um indicador que relaciona o valor de uma empresa com seu ativos. <br/> Ele é calculado dividindo-se o <i> enterprise value </i> pelo ativos da empresa.",
          "groups": ["price"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "0,70"
        },
        {
          "text": "Preço/VPA",
          "key": "price_net_worth_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": true,
          "insurance": true,
          "description": "Preço/VPA é um indicador que relaciona o preço atual da ação com o valor patrimonial por ação da empresa. <br/> Ele é calculado dividindo-se o preço da ação pelo valor patrimonial por ação.",
          "groups": ["price", "extra"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "1,14"
        },
        {
          "text": "Preço/Lucro",
          "key": "price_net_profit_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": true,
          "insurance": true,
          "description": "Preço/Lucro é um indicador que relaciona o preço atual da ação com o lucro por ação da empresa. <br/> Ele é calculado dividindo-se o preço da ação pelo lucro por ação.",
          "groups": ["price", "extra"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "7,97"
        },
        {
          "text": "Preço/Receita Líquida",
          "key": "price_net_revenue_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": true,
          "insurance": true,
          "description": "Preço/Receita Líquida é um indicador que relaciona o preço atual da ação com a receita líquida por ação da empresa. <br/> Ele é calculado dividindo-se o preço da ação pela receita líquida por ação.",
          "groups": ["price"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "3,12"
        },
        {
          "text": "Preço/Ativos",
          "key": "price_assets_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": true,
          "insurance": true,
          "description": "Preço/Ativos é um indicador que relaciona o preço atual da ação com o ativos por ação da empresa. <br/> Ele é calculado dividindo-se o preço da ação pelo ativos por ação.",
          "groups": ["price"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "0,55"
        },
        {
          "text": "Preço/EBIT",
          "key": "price_ebit_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "daily",
          "bank": false,
          "insurance": false,
          "description": "Preço/EBIT é um indicador que relaciona o preço atual da ação com o lucro operacional por ação da empresa. <br/> Ele é calculado dividindo-se o preço da ação pelo EBIT por ação.",
          "groups": ["price"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "5,69"
        }
      ]
    },
    {
      "key": "profit",
      "text": "Indicadores de rentabilidade",
      "metrics": [
        {
          "text": "Margem bruta (LTM)",
          "key": "gross_margin",
          "format": "float_percentage_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "Margem bruta é um indicador que representa a porcentagem de lucro bruto em relação à receita líquida da empresa. <br/> Ele é calculado calculado dividindo-se o lucro bruto pela receita líquida.",
          "groups": ["profit"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_percentage_ptbr",
            "to_value_function": "ptbr_percentage_to_float"
          },
          "value": "52,74 %"
        },
        {
          "text": "Margem líquida (LTM)",
          "key": "net_margin",
          "format": "float_percentage_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "Margem líquida é um indicador que representa a porcentagem de lucro líquido em relação à receita líquida da empresa. <br/> Ele é calculado dividindo-se o lucro líquido pela receita líquida.",
          "groups": ["profit"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_percentage_ptbr",
            "to_value_function": "ptbr_percentage_to_float"
          },
          "value": "39,21 %"
        },
        {
          "text": "Margem EBIT (LTM)",
          "key": "ebit_margin",
          "format": "float_percentage_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Margem EBIT é um indicador que representa a porcentagem de lucro operacional em relação à receita líquida da empresa. <br/> Ele é calculado dividindo-se o EBIT pela receita líquida.",
          "groups": ["profit"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_percentage_ptbr",
            "to_value_function": "ptbr_percentage_to_float"
          },
          "value": "54,96 %"
        },
        {
          "text": "ROIC (LTM)",
          "key": "roic",
          "format": "float_percentage_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "ROIC (<i>Return on Invested Capital </i>) é um indicador que representa a porcentagem de retorno sobre o capital investido na empresa. <br/> Ele é calculado dividindo-se o lucro operacional pela soma do patrimônio líquido com os empréstimos e financiamentos.",
          "groups": ["profit"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_percentage_ptbr",
            "to_value_function": "ptbr_percentage_to_float"
          },
          "value": "12,25 %"
        },
        {
          "text": "ROE (LTM)",
          "key": "roe",
          "format": "float_percentage_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "ROE (Return on Equity) é um indicador que representa a porcentagem de retorno sobre o patrimônio líquido da empresa. <br/> Ele é calculado dividindo-se o lucro líquido pelo patrimônio líquido.",
          "groups": ["profit"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_percentage_ptbr",
            "to_value_function": "ptbr_percentage_to_float"
          },
          "value": "14,36 %"
        },
        {
          "text": "ROA (LTM)",
          "key": "roa",
          "format": "float_percentage_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "ROA (Return on Assets) é um indicador que representa a porcentagem de retorno sobre os ativos da empresa. <br/> Ele é calculado dividindo-se o lucro líquido pelos ativos.",
          "groups": ["profit"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_percentage_ptbr",
            "to_value_function": "ptbr_percentage_to_float"
          },
          "value": "6,93 %"
        }
      ]
    },
    {
      "key": "debt",
      "text": "Indicadores de endividamento",
      "metrics": [
        {
          "text": "Dív. líquida/EBIT",
          "key": "net_debt_ebit_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Dívida líquida/EBIT é um indicador que avalia a capacidade de uma empresa pagar sua dívida com base em seu EBIT. <br/> Ele é calculado dividindo-se a dívida líquida pelo EBIT.",
          "groups": ["debt"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "2,36"
        },
        {
          "text": "Dív. líquida/EBITDA",
          "key": "net_debt_ebitda_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Dívida líquida/EBITDA é um indicador que avalia a capacidade de uma empresa pagar sua dívida com base em seu EBITDA. <br/> Ele é calculado dividindo-se a dívida líquida pelo EBITDA.",
          "groups": ["debt"],
          "table": {
            "show_by_default": true,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "2,34"
        },
        {
          "text": "Dív. líquida/Patr. líquido",
          "key": "net_debt_net_worth_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Dívida líquida/Patrimônio líquido é um indicador que mostra a proporção da dívida líquida de uma empresa em relação ao seu patrimônio líquido. <br/> Ele é calculado dividindo-se a dívida líquida pelo patrimônio líquido.",
          "groups": ["debt"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "0,47"
        },
        {
          "text": "Dív. bruta/Patr. líquido",
          "key": "gross_debt_net_worth_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "Dívida bruta/Patrimônio líquido é um indicador que mostra a proporção da dívida bruta de uma empresa em relação ao seu patrimônio líquido. <br/> Ele é calculado dividindo-se a dívida bruta pelo patrimônio líquido",
          "groups": ["debt"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "0,62"
        },
        {
          "text": "Patr. líquido/Ativos",
          "key": "net_worth_assets_ratio",
          "format": "float_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "Patrimônio líquido/Ativos é um indicador que mostra a proporção dos fundos próprios de uma empresa em relação ao total de ativos. <br/> Ele é calculado dividindo-se o patrimônio líquido da empresa pelo total de ativos.",
          "groups": ["debt"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_number_ptbr",
            "to_value_function": "ptbr_number_to_float"
          },
          "value": "0,48"
        }
      ]
    },
    {
      "key": "target_price",
      "text": "Preço-alvo",
      "metrics": [
        {
          "text": "Fórmula de Graham",
          "key": "target_price_graham",
          "format": "human_readable_currency_ptbr",
          "enable_history": false,
          "history_granularity": "quarterly",
          "bank": true,
          "insurance": true,
          "description": "A Fórmula de Graham, concebida pelo economista Benjamin Graham, é uma abordagem para determinar o maior valor que um investidor deveria pagar por uma ação.<br/> O preço-alvo é dado pela fórmula √(22,5 x LPA x VPA).",
          "groups": ["target_price"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 50,75"
        }
      ]
    },
    {
      "key": "cash_flow",
      "text": "Fluxo de caixa",
      "metrics": [
        {
          "text": "Fluxo de caixa operacional",
          "key": "operating_cash_flow",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "",
          "groups": ["cash_flow"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 440,86 M"
        },
        {
          "text": "Fluxo de caixa de investimentos",
          "key": "investing_cash_flow",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "",
          "groups": ["cash_flow"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ -1,42 B"
        },
        {
          "text": "Fluxo de caixa de financiamentos",
          "key": "financing_cash_flow",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "",
          "groups": ["cash_flow"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 325,95 M"
        },
        {
          "text": "CAPEX",
          "key": "capex",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "",
          "groups": ["cash_flow"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ -34,43 M"
        },
        {
          "text": "Fluxo de caixa livre",
          "key": "free_cash_flow",
          "format": "human_readable_currency_ptbr",
          "enable_history": true,
          "history_granularity": "quarterly",
          "bank": false,
          "insurance": false,
          "description": "",
          "groups": ["cash_flow"],
          "table": {
            "show_by_default": false,
            "show_in_filter_by_column": true,
            "sort_function_name": "sort_human_readable_currency_ptbr",
            "to_value_function": "human_readable_currency_to_float"
          },
          "value": "R$ 406,42 M"
        }
      ]
    }
  ],
  "metadata": {
    "company": "ISA CTEEP - Transmissão Paulista",
    "cd_cvm": 18376,
    "all_tickers": "TRPL3,TRPL4",
    "price": "R$ 32,30",
    "dy": "6,82 %",
    "price_net_worth_ratio": "1,14",
    "price_net_profit_ratio": "7,97",
    "ticker": "TRPL3",
    "dt_last_update": "17/10/2024",
    "is_bank": false,
    "is_insurance": false,
    "judicial_recovery": false,
    "most_traded": false,
    "pronoun": "a",
    "history_text": "\n        <div>\n            <p>\n                Nos útlimos 12 meses, a ISA CTEEP - Transmissão Paulista reportou uma <b>receita líquida de R$ 6,81 B</b> e um <b>lucro líquido de R$ 2,67 B</b>, resultando em uma <b>margem líquida de <span style=\"white-space:nowrap;\">39,21 %</span></b>. Com um <b>patrimônio líquido</b> de <b>R$ 18,60 B</b>, a ISA CTEEP - Transmissão Paulista alcançou um <b>ROE de <span style=\"white-space:nowrap;\">14,36 %</span></b>.\n            </p>\n        </div>\n\n        <div>\n            <p>\n                Em 17/10/2024, a ISA CTEEP - Transmissão Paulista estava sendo negociada com um <b>P/L de 7,97</b>, com um <b>P/VPA de 1,14</b> e com um <b><i>Dividend Yield</i> (DY) de <span style=\"white-space:nowrap;\">6,82 %</span></b>.\n            </p>\n        </div>\n        "
  }
}

	`)

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
	fmt.Println(int(marketCap / price), marketCap / price)

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
	if strings.Contains(s, "M") {
		value *= 1e6
	} else if strings.Contains(s, "B") {
		value *= 1e9
	}

	return value, nil
}
