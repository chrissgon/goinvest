package bot

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/chrissgon/goinvest/controller"
	"github.com/chrissgon/lowbot"
)

func StartBot() {
	lowbot.SetCustomActions(lowbot.ActionsMap{
		"Search": func(flow *lowbot.Flow, interaction *lowbot.Interaction, channel lowbot.IChannel) (bool, error) {
			// step := flow.CurrentStep
			fmt.Println("test")
			stockID := flow.GetLastResponseText()
			fmt.Println(stockID)

			stockController := controller.StockController{}
			stock, err := stockController.Search(stockID)

			if err != nil {
				in := lowbot.NewInteractionMessageText(channel, interaction.Destination, interaction.Sender, "Infelizmente ocorreu um erro")
				err := channel.SendText(in)
				panic(err)
			}

			indicators, err := stockController.Analyse(stock)

			if err != nil {
				in := lowbot.NewInteractionMessageText(channel, interaction.Destination, interaction.Sender, "Infelizmente ocorreu um erro")
				err := channel.SendText(in)
				panic(err)
			}

			sb := strings.Builder{}

			sb.WriteString(fmt.Sprintf("🏢 %v - %v \n\n", stock.ID, formatFloat64ToString(stock.Price)))

			sb.WriteString(fmt.Sprintf("Empresa - %v \n", stock.Company))
			sb.WriteString(fmt.Sprintf("Lucro Líquido - %v \n", formatFloat64ToString(stock.NetProfit)))
			sb.WriteString(fmt.Sprintf("Receita Líquida - %v \n", formatFloat64ToString(stock.NetRevenue)))
			sb.WriteString(fmt.Sprintf("Patrimônio Líquido - %v \n", formatFloat64ToString(stock.NetEquity)))
			sb.WriteString(fmt.Sprintf("Despesa Líquida - %v \n", formatFloat64ToString(stock.NetDebt)))
			sb.WriteString(fmt.Sprintf("Total de Ações - %v \n \n", stock.Shares))

			sb.WriteString("📈 Indicadores\n\n")

			for _, indicator := range indicators {
				sb.WriteString(fmt.Sprintf("%v (%v) - Margem Baseada (%v)", indicator.Label, aroundFloor(indicator.Value), indicator.Mark))

				if indicator.Good {
					sb.WriteString(" ✅")
				} else {
					sb.WriteString(" ❌")
				}

				sb.WriteString("\n\n")
			}

			in := lowbot.NewInteractionMessageText(channel, interaction.Destination, interaction.Sender, sb.String())
			channel.SendText(in)

			return true, nil
		},
		// func(flow *lowbot.Flow, interaction *lowbot.Interaction, channel lowbot.IChannel) (bool, error) {
		// 	// step := flow.CurrentStep
		// 	// template := lowbot.ParseTemplate(step.Parameters.Texts)
		// 	// templateWithUsername := fmt.Sprintf(template, flow.GetLastResponseText())
		// 	// in := lowbot.NewInteractionMessageText(channel, interaction.Destination, interaction.Sender, templateWithUsername)
		// 	// err := channel.SendText(in)
		// 	// return true, err
		// }
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

	lowbot.StartConsumer(consumer, []lowbot.IChannel{channel})
}

func aroundFloor(value float64) float64 {
	return math.Floor((value)*100) / 100
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