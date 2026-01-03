package cmd

import (
	"bybit/auth"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account [symbol]",
	Short: "Check account info",
	Run: func(cmd *cobra.Command, Args []string) {
		Coin := Args[0]
		AccountInfo(Coin)
	},
}

type JSON struct {
	Result struct {
		List []struct {
			Coin []struct {
				WalletBalance string `json:"walletBalance"`
			} `json:"coin"`
		} `json:"list"`
	} `json:"result"`
}

func AccountInfo(COIN string) {
	path := "/v5/account/wallet-balance"
	query := fmt.Sprintf("accountType=UNIFIED&coin=%s", COIN)

	body, err := auth.GetAuth(path, query)
	if err != nil {
		fmt.Printf("Failed to fetch AccountInfo: %v\n", err)
		return
	}

	var data JSON

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Can't read data: %v", err)
		return
	}

	Result := false
	for _, result := range data.Result.List {
		for _, Wbalance := range result.Coin {
			fmt.Printf("\n%s: %s\n", COIN, Wbalance.WalletBalance)
			Result = true
		}
	}
	if !Result {
		fmt.Printf("\nNo data returned for asset: %s\n", COIN)
	}

}
