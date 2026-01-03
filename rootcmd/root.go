package root

import (
	"bybit/cmd"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "Bybit",
	Short: "Bybit-CLI",
}

func init() {
	Cmd.AddCommand(cmd.MarketTickerCmd)
	Cmd.AddCommand(cmd.AccountCmd)
	Cmd.AddCommand(cmd.PlaceOrderCmd)
	Cmd.AddCommand(cmd.CancelOrderCmd)
	Cmd.AddCommand(cmd.ShowOrderCmd)
	Cmd.AddCommand(cmd.CancelAllCmd)
	Cmd.AddCommand(cmd.OrderHistoryCmd)

	//PlaceOrder
	cmd.PlaceOrderCmd.Flags().StringVar(&cmd.BuySymbol, "symbol", "", "trading pair (BTCUSDT)")
	cmd.PlaceOrderCmd.Flags().StringVar(&cmd.BuyQty, "qty", "", "order quantity")
	cmd.PlaceOrderCmd.Flags().StringVar(&cmd.BuyPrice, "price", "", "orderprice")
	cmd.PlaceOrderCmd.Flags().StringVar(&cmd.BuyOrderType, "type", "limit", "limit/market")
	cmd.PlaceOrderCmd.Flags().StringVar(&cmd.Side, "side", "", "buy/sell")
	cmd.PlaceOrderCmd.Flags().StringVar(&cmd.Category, "category", "", "spot/inverse/linear/option")

	//CancelOrder
	cmd.CancelOrderCmd.Flags().StringVar(&cmd.CancelSymbol, "symbol", "", "symbol (symbol)")
	cmd.CancelOrderCmd.Flags().StringVar(&cmd.CancelOrderId, "id", "", "id")
	cmd.CancelOrderCmd.Flags().StringVar(&cmd.CancelCategory, "category", "", "spot/inverse/linear/option")

	//Show order history
	cmd.OrderHistoryCmd.Flags().StringVar(&cmd.Limit,
		"limit",
		"",
		"--limit (num)")

	//CancelALl
	cmd.CancelAllCmd.Flags().StringVar(&cmd.CancelAllCategory,
		"category",
		"",
		"--category spot/inverse/linear/option")
}
