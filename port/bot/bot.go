package bot

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/chrissgon/goinvest/controller"
	"github.com/chrissgon/goinvest/domain"
	"github.com/chrissgon/goinvest/domain/fund"
	"github.com/chrissgon/lowbot"
	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
)

func StartBot() {
	lowbot.DEBUG = true

	lowbot.SetCustomActions(lowbot.ActionsMap{
		"SearchStock": func(interaction *lowbot.Interaction) (*lowbot.Interaction, bool) {
			stockID := interaction.Parameters.Text
			stockController := controller.StockController{}
			stockEntity, err := stockController.Search(stockID)

			if err != nil {
				in := lowbot.NewInteractionMessageText("Não foi possível encontrar a ação informada.")
				in.SetTo(interaction.To)
				in.SetFrom(interaction.From)
				return in, false
			}

			// indicators, err := stockController.Analyse(stockEntity)

			if err != nil {
				in := lowbot.NewInteractionMessageText("Ocorreu um erro ao gerar os indicadores.\n Por favor, tente novamente mais tarde.")
				in.SetTo(interaction.To)
				in.SetFrom(interaction.From)
				return in, false
			}

			text := fmt.Sprintf(`
🏢 %s - %s

🔹 Empresa (%s)
🔹 Lucro Líquido: %s
🔹 Receita Líquida: %s
🔹 Patrimônio Líquido: %s
🔹 Despesa Líquida: %s
🔹 Total de Ações: %s

📊 Indicadores e Análise


⚠️ Go Invest pode cometer erros. Verifique informações importantes.

https://visnoinvest.com.br/stocks/PETR4/%s/`,
				stockEntity.ID,
				formatFloat64ToString(stockEntity.Price),
				stockEntity.Company,
				formatFloat64ToString(stockEntity.NetProfit),
				formatFloat64ToString(stockEntity.NetRevenue),
				formatFloat64ToString(stockEntity.NetEquity),
				formatFloat64ToString(stockEntity.NetDebt),
				humanize.Comma(int64(stockEntity.Shares)),
				// getIndicatorText(indicators[stock.PER_NAME]),
				// getIndicatorText(indicators[stock.PBV_NAME]),
				// getIndicatorText(indicators[stock.PROFIT_MARGIN_NAME]),
				// getIndicatorText(indicators[stock.ROE_NAME]),
				// getIndicatorText(indicators[stock.DEBIT_RATIO_NAME]),
				// getIndicatorText(indicators[stock.DIVIDEND_YELD_NAME]),
				stockEntity.ID,
			)
			// sb := strings.Builder{}

			// sb.WriteString(fmt.Sprintf("🏢 %v - %v \n\n", stockEntity.ID, formatFloat64ToString(stockEntity.Price)))

			// sb.WriteString(fmt.Sprintf("Empresa (%v) \n", stockEntity.Company))
			// sb.WriteString(fmt.Sprintf("\nLucro Líquido \n%v \n", formatFloat64ToString(stockEntity.NetProfit)))
			// sb.WriteString(fmt.Sprintf("\nReceita Líquida \n%v \n", formatFloat64ToString(stockEntity.NetRevenue)))
			// sb.WriteString(fmt.Sprintf("\nPatrimônio Líquido \n%v \n", formatFloat64ToString(stockEntity.NetEquity)))
			// sb.WriteString(fmt.Sprintf("\nDespesa Líquida \n%v \n", formatFloat64ToString(stockEntity.NetDebt)))
			// sb.WriteString(fmt.Sprintf("\nTotal de Ações \n%v \n \n", stockEntity.Shares))

			// sb.WriteString("📈 Indicadores\n\n")

			// getIndicatorText(&sb, indicators[stock.PER_NAME])
			// getIndicatorText(&sb, indicators[stock.PBV_NAME])
			// getIndicatorText(&sb, indicators[stock.PROFIT_MARGIN_NAME])
			// getIndicatorText(&sb, indicators[stock.ROE_NAME])
			// getIndicatorText(&sb, indicators[stock.DEBIT_RATIO_NAME])
			// getIndicatorText(&sb, indicators[stock.DIVIDEND_YELD_NAME])

			in := lowbot.NewInteractionMessageText(text)
			in.SetTo(interaction.To)
			in.SetFrom(interaction.From)

			return in, true
		},
		"SearchFund": func(interaction *lowbot.Interaction) (*lowbot.Interaction, bool) {
			fundID := interaction.Parameters.Text
			fundController := controller.FundController{}
			fundEntity, err := fundController.Search(fundID)

			if err != nil {
				in := lowbot.NewInteractionMessageText("Não foi possível encontrar o fundo informado.")
				in.SetTo(interaction.To)
				in.SetFrom(interaction.From)
				return in, false
			}

			indicators, err := fundController.Analyse(fundEntity)

			if err != nil {
				in := lowbot.NewInteractionMessageText("Ocorreu um erro ao gerar os indicadores.\n Por favor, tente novamente mais tarde.")
				in.SetTo(interaction.To)
				in.SetFrom(interaction.From)
				return in, false
			}

			text := fmt.Sprintf(`
🏢 %s - %s

🔹 Administrador: %s
🔹 Patrimônio Líquido: %s
🔹 Último Rendimento: %s por cota
🔹 Dividend Yield (Últimos 12 meses): %v%%
🔹 Taxa de Administração: %v%% a.a
🔹 Total de Cotas: %s

📊 Indicadores e Análise

%s
%s
%s

⚠️ Go Invest pode cometer erros. Verifique informações importantes.

https://fiis.com.br/%s/`,
				fundEntity.ID,
				formatFloat64ToString(fundEntity.Price),
				fundEntity.Administrator,
				formatFloat64ToString(fundEntity.NetEquity),
				formatFloat64ToString(fundEntity.LastIncome),
				fundEntity.DividendYieldAnnual,
				fundEntity.AdministrationFee,
				humanize.Comma(int64(fundEntity.Shares)),
				getIndicatorText(indicators[fund.PBV_NAME]),
				getIndicatorText(indicators[fund.DIVIDEND_YELD_MONTH_NAME]),
				getIndicatorText(indicators[fund.ADMINISTRATION_FEE_NAME]),
				fundEntity.ID,
			)

			in := lowbot.NewInteractionMessageText(text)
			in.SetTo(interaction.To)
			in.SetFrom(interaction.From)

			return in, true
		},
	})

	channel, err := lowbot.NewTelegramChannel(os.Getenv("TELEGRAM_TOKEN"))

	if err != nil {
		panic(err)
	}

	persist, err := lowbot.NewMemoryFlowPersist()

	if err != nil {
		panic(err)
	}

	flow, err := lowbot.NewFlow("./port/bot/flow.yaml")

	if err != nil {
		panic(err)
	}

	consumer := lowbot.NewJourneyConsumer(flow, persist)

	bot := lowbot.NewBot(consumer, map[uuid.UUID]lowbot.IChannel{
		channel.GetChannel().ChannelID: channel,
	})

	bot.Start()

	// keep the process running
	sc := make(chan os.Signal, 1)
	<-sc
}

func getIndicatorText(indicator domain.Indicator) string {

	valuePrefix := ""
	markPrefix := ""
	symbol := "❌"

	if indicator.Good {
		symbol = "✅"
	}
	if indicator.Name == fund.PBV_NAME {
		markPrefix = ""
	}
	if indicator.Name == fund.DIVIDEND_YELD_MONTH_NAME {
		valuePrefix = "% a.m"
		markPrefix = "%"
	}
	if indicator.Name == fund.ADMINISTRATION_FEE_NAME {
		valuePrefix = "% a.a"
		markPrefix = "%"
	}

	return fmt.Sprintf(`%s %s
	- Valor Atual: %v%s
	- Referência Ideal: %s %v%s
`,
		symbol,
		indicator.Label,
		toFixed(indicator.Value, 2),
		valuePrefix,
		indicator.Operator,
		toFixed(float64(indicator.Mark), 2),
		markPrefix,
	)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func formatFloat64ToString(value float64) string {
	// Define thresholds for K, M, B
	thousand := 1000.0
	million := 1000000.0
	billion := 1000000000.0

	var result float64
	var suffix string

	switch {
	case value >= billion:
		result = value / billion
		suffix = "B"
	case value >= million:
		result = value / million
		suffix = "M"
	case value >= thousand:
		result = value / thousand
		suffix = "K"
	default:
		result = value
		suffix = ""
	}

	// Round to two decimal places
	result = math.Round(result*100) / 100

	// Convert to string with two decimal places
	strValue := strconv.FormatFloat(result, 'f', 2, 64)

	// Replace dot with comma for decimal separator
	strValue = strings.Replace(strValue, ".", ",", 1)

	// Add thousand separators
	parts := strings.Split(strValue, ",")
	integerPart := parts[0]
	decimalPart := parts[1]

	var formattedInteger string
	for i, r := range reverse(integerPart) {
		if i > 0 && i%3 == 0 {
			formattedInteger = "." + formattedInteger
		}
		formattedInteger = string(r) + formattedInteger
	}

	// Combine all parts
	return fmt.Sprintf("R$ %s,%s %s", formattedInteger, decimalPart, suffix)
}

// Helper function to reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
