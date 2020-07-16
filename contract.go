package goibcp

import (
	"strings"
)

//GetFutresContractBySymbols - Gte list of futures contract by symbols Functions related to fetching IB contract informaion
func (c *IBClient) GetFutresContractBySymbols(symbols []string, res interface{}) error {
	// epURL := Settings.CPURL + endpoints["contractFuturesBySymbols"]
	qs := "symbols=" + strings.Join(symbols, ",")
	return c.GetEndpoint("contractFuturesBySymbols", res, qs)
	// req := rClient.R().SetResult(res).SetQueryString(qs)
	// _, err := req.Get(epURL)
	// if err != nil {
	// 	logMsg(ERROR, "GetFutresContractBySymbols", "Failed to get", err)
	// 	return err
	// }
	// logMsg(INFO, "GetFutresContractBySymbols", fmt.Sprintf("%+v", res))
	// return nil
}
