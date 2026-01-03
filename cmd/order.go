package cmd

import (
	"bybit/auth"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	//BUY
	BuySymbol    string
	BuyPrice     string
	BuyQty       string
	BuyOrderType string
	Category     string
	Side         string

	//CancelOrder
	CancelCategory string
	CancelOrderId  string
	CancelSymbol   string

	//cancelAll
	CancelAllCategory string
)

var PlaceOrderCmd = &cobra.Command{
	Use: "buy",
	Run: func(cmd *cobra.Command, Args []string) {
		if BuySymbol == "" || BuyQty == "" || BuyPrice == "" || BuyOrderType == "" {
			fmt.Println("Require value to fill!")
			return
		}
		PlaceOrder(BuySymbol, BuyPrice, BuyQty, BuyOrderType, Category, Side)
	},
}

var CancelOrderCmd = &cobra.Command{
	Use: "cancel",
	Run: func(cmd *cobra.Command, Args []string) {
		if CancelSymbol == "" || CancelOrderId == "" {
			fmt.Println("Require value to fill!")
			return
		}
		CancelOrder(CancelCategory, CancelSymbol, CancelOrderId)
	},
}

var ShowOrderCmd = &cobra.Command{
	Use: "showorder",
	Run: func(cmd *cobra.Command, Args []string) {
		ShowOrders()
	},
}

var CancelAllCmd = &cobra.Command{
	Use: "cancelall",
	Run: func(Cmd *cobra.Command, Args []string) {
		CancelAllOrders(CancelAllCategory)
	},
}

func PlaceOrder(buySymbol, buyPrice, buyQty, buyOrderType, category, side string) {
	type OrderResponse struct {
		RetCode int    `json:"retCode"`
		RetMsg  string `json:"retMsg"`
		Result  struct {
			OrderId string `json:"orderId"`
		} `json:"result"`
	}

	path := "/v5/order/create"
	data := map[string]interface{}{
		"category":  category,
		"symbol":    buySymbol,
		"side":      side,
		"orderType": buyOrderType,
		"qty":       buyQty,
		"price":     buyPrice,
	}

	body, err := auth.PostAuth(path, data)
	if err != nil {
		fmt.Printf("Error fetch Place Order: %v", err)
		return
	}

	var response OrderResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("err Unmarshall: %v", err)
		return
	}

	if response.RetCode != 0 {
		fmt.Printf("Order Failed: %s", err)
		return
	}

	fmt.Println("Order placed successfully!")
	fmt.Printf("Category: %s\n", category)
	fmt.Printf("Symbol: %s\n", buySymbol)
	fmt.Printf("Side: %s\n", side)
	fmt.Printf("OrderType: %s\n", buyOrderType)
	fmt.Printf("Qty: %s\n", buyQty)
	fmt.Printf("Price: $%s\n", buyPrice)

}

func CancelOrder(CancelCategory, CancelSymbol, CancelOrderId string) {
	type Cancel struct {
		RetMsg string `json:"retMsg"`
		Result struct {
			OrderId string `json:"orderId"`
		} `json:"result"`
	}

	path := "/v5/order/cancel"
	data := map[string]interface{}{
		"category": CancelCategory,
		"symbol":   CancelSymbol,
		"orderId":  CancelOrderId,
	}
	body, err := auth.PostAuth(path, data)
	if err != nil {
		fmt.Printf("Error when fetch CancelOrder: %v", err)
		return
	}
	var cancel Cancel

	err = json.Unmarshal(body, &cancel)
	if err != nil {
		fmt.Printf("Err resp cancel order: %v", err)
		return
	}

	if cancel.RetMsg != "OK" {
		fmt.Println("\nOrder does not exist")
		return
	}

	fmt.Printf("category: %s\nsymbol: %s\norderId: %s\n",
		CancelCategory,
		CancelSymbol,
		CancelOrderId)

	if cancel.RetMsg == "OK" {
		fmt.Println("Done cancel!")
	}

}

func ShowOrders() {
	type Orders struct {
		Result struct {
			List []struct {
				Symbol    string `json:"symbol"`
				OrderType string `json:"orderType"`
				LinkId    string `json:"orderLinkId"`
				OrderId   string `json:"orderId"`
			} `json:"list"`
		} `json:"result"`
	}

	path := "/v5/order/realtime"
	query := fmt.Sprintf("category=spot&limit=5")

	result, err := auth.GetAuth(path, query)
	if err != nil {
		fmt.Printf("Error when fetch ShowOrders: %v", err)
		return
	}

	var order Orders

	err = json.Unmarshal(result, &order)
	if err != nil {
		fmt.Printf("Error resp ShowOrders: %v", err)
		return
	}

	hasOrders := false
	for _, Show := range order.Result.List {
		fmt.Printf("Symbol: %s | OrderType: %s | LinkId: %s | OrderId: %s\n",
			Show.Symbol,
			Show.OrderType,
			Show.LinkId,
			Show.OrderId)
		hasOrders = true
	}
	if !hasOrders {
		fmt.Println("Order is empty")
	}
}

func CancelAllOrders(CancelAllCategory string) {
	type JSON struct {
		Msg    string `json:"retMsg"`
		Result struct {
			List []struct {
				OrderID string `json:"orderId"`
			} `json:"list"`
		} `json:"result"`
	}

	path := "/v5/order/cancel-all"
	data := map[string]interface{}{
		"category": CancelAllCategory,
	}

	body, err := auth.PostAuth(path, data)
	if err != nil {
		fmt.Printf("!! Error cancel all: %v\n", err)
		return
	}

	var Orders JSON

	err = json.Unmarshal(body, &Orders)

	if Orders.Msg != "OK" {
		fmt.Println("Should type categories!\nUsage: --category spot/inverse/option/linear")
		return
	}

	fmt.Printf("Msg: %s\n", Orders.Msg)

	orders := false
	for _, result := range Orders.Result.List {
		fmt.Printf("OrderId: %s\n",
			result.OrderID)
		orders = true
	}

	if orders == true {
		fmt.Println("All order has been cancelled")
	}

	if !orders {
		fmt.Println("There's no order to cancel")
	}

}
