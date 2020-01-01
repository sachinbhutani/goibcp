package goibcp

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func logMsg(level int, fn string, msg string, err ...error) {
	if level <= Settings.LogLevel {
		switch level {
		case 0:
			fmt.Println("ERROR in function", fn, ":", msg, err)
		case 1:
			fmt.Println("ERROR in function", fn, ":", msg, err)
		case 2:
			fmt.Println("GOIBCP", fn, ":", msg)
		case 3:
			fmt.Println("GOIBCP Debug:", fn, msg)
		}
	}
}

//GetEndpoint - Get an endpoint from IBCP and return a map
func (c *IBClient) GetEndpoint(endp string, res interface{}, qs ...string) error {
	epURL := Settings.CPURL + endpoints[endp]
	var req *resty.Request
	if len(qs) > 0 {
		req = rClient.R().SetResult(res).SetQueryString(qs[0])
	} else {
		req = rClient.R().SetResult(res).SetHeader("Accept", "application/json")
	}
	resp, err := req.Get(epURL)
	if err != nil {
		logMsg(ERROR, endp, "Failed to get", err)
		return err
	}
	logMsg(DEBUG, endp, resp.String())
	return nil
}

//PostEndpoint - Post to an endpoint from IBCP
func (c *IBClient) PostEndpoint(endp string, res interface{}) error {
	epURL := Settings.CPURL + endpoints[endp]
	resp, err := rClient.R().SetResult(res).SetHeader("Content-Type", "application/json").Post(epURL)
	if err != nil {
		logMsg(ERROR, endp, "Failed to post", err)
		return err
	}
	logMsg(DEBUG, endp, resp.String())
	return nil
}

//SessionStatus - Returns session status
func (c *IBClient) SessionStatus() error {
	statusURL := Settings.CPURL + endpoints["sessionStatus"]
	_, err := rClient.R().SetResult(c).Get(statusURL)
	if err != nil {
		logMsg(ERROR, "SessionStatus", "Error getting session status", err)
		return err
	}
	logMsg(INFO, "SessionStatus:", fmt.Sprintf("%+v", c))
	return nil
}

//AutoTickle - Keeps the sesssion alive by tickeling the server every minute unless an error is encountered or session expires
func AutoTickle(c *IBClient) error {
	var reply IBUser
	var err error
	for {
		time.Sleep(60 * time.Second)
		err = c.GetEndpoint("sessionValidateSSO", &reply)
		logMsg(INFO, "AutoTickle", fmt.Sprintf("%+v", reply))
		if err != nil {
			break
		}
		if reply.Expires == 0 {
			err = errors.New("Session Expired")
			break
		}
	}
	return err
}

//TODO: create Helper methods to place simple market and limit orders for stocks and futures
