package goibcp

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_Futures_Trade(t *testing.T) {
	ib, err := Connect(&Config{CPURL: URL, LogLevel: 2})
	if err != nil {
		t.Error("Not able to connect")
		return
	}
	//verify connection
	if ib.IsAuthenticated != true {
		t.Error("Connection not authenticated")
		return
	}
	var contracts IBFutContractList
	err = ib.GetFutresContractBySymbols([]string{"ES", "MNQ"}, &contracts)
	if err != nil {
		t.Error("error getting contracts")
	}
	//TODO: Replace with asset function , assert the symbol and underlying contract id as date and current contract can change
	//TODO: also asset the expiration date is not 0
	mnq := contracts["MNQ"][0]
	fmt.Printf("First MNQ contrat: %+v \n", mnq)
	//Trade the first MNQ Contract
	mnqOrder := IBOrder{Conid: mnq.Conid, SecType: strconv.Itoa(mnq.Conid) + ":FUT", COID: "TESTS03", OrderType: "MKT", ListingExchange: "SMART", Side: "SLL", Tif: "DAY", Referrer: "GOIBCP", Quantity: 1}
	var iserverAccount IBTradeAccount
	ib.GetTradeAccount(&iserverAccount)
	fmt.Println("Selected account:", iserverAccount.SelectedAccount)
	orderReply, err := ib.PlaceOrder(mnqOrder)
	if err != nil {
		t.Error("Error while placing order", err.Error())
	}
	fmt.Printf("OrderReply: %+v \n", orderReply)
}

func Test_Get_Live_Orders(t *testing.T) {
	var liveOrders IBLiveOrders
	ib, err := Connect(&Config{CPURL: URL, LogLevel: 0})
	if err != nil {
		t.Error("Not able to connect")
		return
	}
	ib.GetLiveOrders(&liveOrders)
	fmt.Printf("%+v", liveOrders.Orders)
}

func Test_Get_Open_Positions(t *testing.T) {
	var openPositions IBPortfolioPositions
	ib, err := Connect(&Config{CPURL: URL, LogLevel: 2})
	if err != nil {
		t.Error("Not able to connect")
		return
	}
	ib.GetPortfolioPositions(&openPositions, 0)
	fmt.Printf("%+v", openPositions)
}

func Test_Get_Trades_List(t *testing.T) {
	var trades IBTrades
	var iserverAccount IBTradeAccount
	ib, err := Connect(&Config{CPURL: URL, LogLevel: 2})
	if err != nil {
		t.Error("Not able to connect")
		return
	}
	ib.GetTradeAccount(&iserverAccount)
	fmt.Println("Selected account:", iserverAccount.SelectedAccount)
	ib.GetTrades(&trades)
	fmt.Printf("%+v", trades)
}

func Test_Get_Account_Ledger(t *testing.T) {
	var ledger IBAccountLedger
	ib, err := Connect(&Config{CPURL: URL, LogLevel: 2})
	if err != nil {
		t.Error("Not able to connect")
		return
	}
	ib.GetAccountLedger(&ledger)
	fmt.Printf("%+v", ledger)
}
