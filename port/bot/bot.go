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
	"github.com/chrissgon/goinvest/domain/stock"
	"github.com/chrissgon/lowbot"
	"github.com/google/uuid"
)

func StartBot() {
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

			indicators, err := stockController.Analyse(stockEntity)

			if err != nil {
				in := lowbot.NewInteractionMessageText("Ocorreu um erro ao gerar os indicadores.\n Por favor, tente novamente mais tarde.")
				in.SetTo(interaction.To)
				in.SetFrom(interaction.From)
				return in, false
			}

			sb := strings.Builder{}

			sb.WriteString(fmt.Sprintf("🏢 %v - %v \n\n", stockEntity.ID, formatFloat64ToString(stockEntity.Price)))

			sb.WriteString(fmt.Sprintf("Empresa (%v) \n", stockEntity.Company))
			sb.WriteString(fmt.Sprintf("\nLucro Líquido \n%v \n", formatFloat64ToString(stockEntity.NetProfit)))
			sb.WriteString(fmt.Sprintf("\nReceita Líquida \n%v \n", formatFloat64ToString(stockEntity.NetRevenue)))
			sb.WriteString(fmt.Sprintf("\nPatrimônio Líquido \n%v \n", formatFloat64ToString(stockEntity.NetEquity)))
			sb.WriteString(fmt.Sprintf("\nDespesa Líquida \n%v \n", formatFloat64ToString(stockEntity.NetDebt)))
			sb.WriteString(fmt.Sprintf("\nTotal de Ações \n%v \n \n", stockEntity.Shares))

			sb.WriteString("📈 Indicadores\n\n")

			getIndicatorText(&sb, indicators[stock.PER_NAME])
			getIndicatorText(&sb, indicators[stock.PBV_NAME])
			getIndicatorText(&sb, indicators[stock.PROFIT_MARGIN_NAME])
			getIndicatorText(&sb, indicators[stock.ROE_NAME])
			getIndicatorText(&sb, indicators[stock.DEBIT_RATIO_NAME])
			getIndicatorText(&sb, indicators[stock.DIVIDEND_YELD_NAME])

			in := lowbot.NewInteractionMessageText(sb.String())
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

			sb := strings.Builder{}

			sb.WriteString(fmt.Sprintf("🏢 %v - %v \n\n", fundEntity.ID, formatFloat64ToString(fundEntity.Price)))

			sb.WriteString(fmt.Sprintf("Administrador (%v) \n", fundEntity.Administrator))
			sb.WriteString(fmt.Sprintf("\nPatrimônio Líquido \n%v \n", formatFloat64ToString(fundEntity.NetEquity)))
			sb.WriteString(fmt.Sprintf("\nÚltimo Rendimento \n%v \n", formatFloat64ToString(fundEntity.LastIncome)))
			sb.WriteString(fmt.Sprintf("\nTaxa de Administração \n%v \n", formatFloat64ToString(fundEntity.AdministrationFee)))
			sb.WriteString(fmt.Sprintf("\nTaxa de Performance \n%v \n", formatFloat64ToString(fundEntity.PerformanceFee)))
			sb.WriteString(fmt.Sprintf("\nDividend Yield (Último 12 meses) \n%v \n", formatFloat64ToString(fundEntity.DividendYieldAnnual)))
			sb.WriteString(fmt.Sprintf("\nTotal de Cotas \n%v \n \n", fundEntity.Shares))

			sb.WriteString("📈 Indicadores\n\n")

			getIndicatorText(&sb, indicators[fund.PBV_NAME])
			getIndicatorText(&sb, indicators[fund.DIVIDEND_YELD_NAME])
			getIndicatorText(&sb, indicators[fund.ADMINISTRATION_FEE_NAME])

			in := lowbot.NewInteractionMessageText(sb.String())
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

func getIndicatorText(sb *strings.Builder, indicator domain.Indicator) {
	if indicator.Good {
		sb.WriteString("✅ ")
	} else {
		sb.WriteString("❌ ")
	}
	sb.WriteString(fmt.Sprintf("%v \n", indicator.Label))
	sb.WriteString(fmt.Sprintf("Valor - %v \n", toFixed(indicator.Value, 2)))
	sb.WriteString(fmt.Sprintf("Marca Sugerida - %v \n", toFixed(float64(indicator.Mark), 2)))

	sb.WriteString("\n")
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
