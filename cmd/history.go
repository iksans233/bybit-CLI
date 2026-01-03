package cmd

import (
	"bybit/auth"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var Limit string

var OrderHistoryCmd = &cobra.Command{
	Use: "orderhistory",
	Run: func(cmd *cobra.Command, Args []string) {
		OrderHistory(Limit)
	},
}

func OrderHistory(Limit string) {
	type History struct {
		RetMsg string `json:"retMsg"`
		Result struct {
			Category string `json:"category"`
			Nested   []struct {
				Symbol    string `json:"symbol"`
				Ordertype string `json:"orderType"`
				Qty       string `json:"qty"`
				Orderid   string `json:"orderId"`
			} `json:"list"`
		} `json:"result"`
	}

	path := "/v5/order/history"
	query := fmt.Sprintf("category=spot&limit=%s", Limit)

	body, err := auth.GetAuth(path, query)
	if err != nil {
		fmt.Printf("Can't fetch Order History!: %v\n", err)
		return
	}

	var orderhistory History

	err = json.Unmarshal(body, &orderhistory)
	if err != nil {
		fmt.Printf("Err response OrderHistory: %v", err)
		return
	}

	for _, result := range orderhistory.Result.Nested {
		fmt.Printf("Symbol: %s\nOrdertype: %s\nqty: %s\nOrderId: %s\n____________\n",
			result.Symbol,
			result.Ordertype,
			result.Qty,
			result.Orderid)
	}
}
