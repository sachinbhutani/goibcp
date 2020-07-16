package goibcp

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

//Version - the version of go-ib-cp
const Version = "0.0.1"

//ERROR, WARNING or INFO constants for Log Levels
const (
	ERROR   = 0
	WARNING = 1
	INFO    = 2
)

//Config to connect to CP Web gateway
//LogInfo 0=> Log Errors only , 1=> log warnings, 2=> log information (default)
type Config struct {
	CPURL      string
	LogLevel   int
	AutoTickle bool
}

//Settings - Default settings if no setting are provided to the Connect() function.
var Settings = &Config{CPURL: "https://localhost:5000", LogLevel: 2}

//Client - IB Client which can be used to call all api functions
var Client IBClient

//User - IBUser
var User IBUser
var rClient = resty.New()

//Connect to CP Web gateway.
//Returns a ib api client if successful or an error if connection is not established.
//If a session is established, auto-trickle function would be triggered to keep the session active using trciker api every minute.
func Connect(userSettings ...*Config) (*IBClient, error) {
	//Overwrite default settings if settings are provided.
	if len(userSettings) > 0 {
		if userSettings[0].CPURL != "" {
			Settings.CPURL = userSettings[0].CPURL
		}
		if userSettings[0].LogLevel != 2 {
			Settings.LogLevel = userSettings[0].LogLevel
		}
	}

	//ValidateSSO
	err := Client.GetEndpoint("sessionValidateSSO", &User)
	if err != nil {
		logMsg(ERROR, "Connect", "Failed to validate SSO", err)
		return &Client, err
	}
	//Get client authentication status, if client is not authenticate, attemp to re-authenticate 1 time.
	for i := 0; i < 2; i++ {
		err = Client.SessionStatus()
		if err != nil {
			logMsg(ERROR, "Connect", "Failed to validate SSO", err)
			return &Client, err
		}
		// if status is not connected, return error.
		if Client.IsConnected == false {
			logMsg(ERROR, "Connect", "Not connected to gateway, please login to CP web gateway again")
			return &Client, errors.New("Not connected to gateway, please login to CP web gateway again")
		}
		// if status is connected, but not authenticated, try to reauthenticate once.
		if Client.IsAuthenticated == false {
			err = Client.PostEndpoint("sessionReauthenticate", &IBClient{})
			if err != nil {
				logMsg(ERROR, "Connect", "Not able to re-authenticate with the gateway..quitting now")
				return &Client, err
			}
			time.Sleep(3 * time.Second)
			continue
		} else {
			//TODO: trigger auto tickle
			return &Client, nil
		}
	}
	fmt.Printf("GOIBCP Client: %+v", Client)
	return &Client, nil
}

//trickle the server every minute to keep the session alive
//this function is not exported but called internally every minute
func (c *IBClient) trickle() {
	fmt.Println("Trickling...")

}

//Logout the IB client
func (c *IBClient) Logout() error {
	return c.GetEndpoint("sessionLogout", &c)
}

//PlaceOrder - posts and order
func (c *IBClient) PlaceOrder(order IBOrder) (IBOrderReply, error) {
	//Get Trading Account
	var orderReply IBOrderReply
	selectedAccount, err := c.GetSelectedAccount()
	if err != nil || selectedAccount == "" {
		logMsg(ERROR, "PlaceOrder", "Not able to find selected Trade account", err)
		return nil, err
	}
	epURL := Settings.CPURL + endpoints["orderPlace"]
	req := rClient.R().SetPathParams(map[string]string{"accountId": selectedAccount}).SetHeader("Content-Type", "application/json")
	req = req.SetBody(order).SetResult(&orderReply)
	resp, err := req.Post(epURL)
	if err != nil {
		logMsg(ERROR, "PlaceOrder", "Failed to post", err)
		return nil, err
	}
	logMsg(INFO, "PlaceOrder", resp.String())
	return orderReply, nil
}

//GetLiveOrders - Update live order to the order list struct
func (c *IBClient) GetLiveOrders(liveOrders *IBLiveOrders) error {
	return c.GetEndpoint("ordersLive", liveOrders)
}

//GetTradeAccount Get TradeAccount Information for the current trade account
func (c *IBClient) GetTradeAccount(res interface{}) error {
	return c.GetEndpoint("accountIserver", res)
}

//GetSelectedAccount - Get selected trade account ID , returns accountID as string or an error
func (c *IBClient) GetSelectedAccount() (string, error) {
	var tradeAccount IBTradeAccount
	err := c.GetEndpoint("accountIserver", &tradeAccount)
	if err != nil {
		logMsg(ERROR, "GetSelectedAccount", "Could not get Iserver trade account info", err)
		return "", err
	}
	return tradeAccount.SelectedAccount, nil
}
