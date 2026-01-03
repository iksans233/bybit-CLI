package cmd

import (
	"bybit/auth"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var MarketTickerCmd = &cobra.Command{
	Use:   "market [symbol]",
	Short: "Market ticker Price",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, Args []string) {
		symbol := Args[0]
		MarketTicker(symbol)
		if symbol == "" {
			fmt.Println("!!: Symbol cannot be empty")
			return
		}
	},
}

func MarketTicker(symbol string) {
	type JSON struct {
		Result struct {
			List []struct {
				Symbol    string `json:"symbol"`
				LastPrice string `json:"lastPrice"`
			} `json:"list"`
		} `json:"result"`
	}

	path := "/v5/market/tickers"
	query := fmt.Sprintf("category=spot&symbol=%s", symbol)

	if symbol == "" {
		fmt.Println("!!: Symbol cannot be empty")
		return
	}

	body, err := auth.GetAuth(path, query)
	if err != nil {
		fmt.Printf("!! Failed to fetch ticker: %v\n", err)
		return
	}

	var ticker JSON

	err = json.Unmarshal(body, &ticker)
	if err != nil {
		fmt.Printf("!! Failed to parse response: %v\n", err)
		return
	}

	for _, list := range ticker.Result.List {
		fmt.Printf("Symbol: %s\n",
			list.Symbol)
		fmt.Printf("Price: $%s\n",
			list.LastPrice)
	}
}
