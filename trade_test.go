package goibcp

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"
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
	err = ib.GetFutresContractBySymbols([]string{"MNQ", "ES"}, &contracts)
	if err != nil {
		t.Error("error getting contracts")
	}
	fmt.Printf("All Contracts Found: \n %+v \n", contracts)
	mnq, err := GetCurrentContract(contracts, "MNQ")
	if err != nil {
		t.Error("Could not find contract")
	}
	//Trade the first MNQ Contract
	mnqOrder := IBOrder{Conid: mnq.Conid, SecType: strconv.Itoa(mnq.Conid) + ":FUT", COID: "TEST03", OrderType: "MKT", ListingExchange: "SMART", Side: "BUY", Tif: "DAY", Referrer: "GOIBCP", Quantity: 1}
	var iserverAccount IBTradeAccount
	ib.GetTradeAccount(&iserverAccount)
	fmt.Println("Selected account:", iserverAccount.SelectedAccount)
	orderReply, err := ib.PlaceOrder(mnqOrder)
	if err != nil {
		t.Error("Error while placing order", err.Error())
	}
	fmt.Printf("OrderReply: %+v \n", orderReply)
}

func GetCurrentContract(conList IBFutContractList, ticker string) (IBFutContract, error) {
	today := time.Now()
	for _, k := range conList[ticker] {
		expDate, err := time.Parse("20060102", fmt.Sprint(k.ExpirationDate))
		if err != nil {
			fmt.Printf("Error reading contract expiry date")
			return k, err
		}
		td := expDate.AddDate(0, 0, -7)
		if td.After(today) {
			return k, nil
		}
		fmt.Printf("%+v \n", k)

	}
	return IBFutContract{}, errors.New("Error gettign contract")
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
