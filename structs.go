package goibcp

//IBClient - Represents the IB API client which executes all API functions
type IBClient struct {
	IsConnected     bool     `json:"connected"`
	IsAuthenticated bool     `json:"authenticated"`
	IsCompeting     bool     `json:"competing"`
	Fail            string   `json:"fail"`
	Message         string   `json:"message"`
	Prompts         []string `json:"prompts"`
}

//IBUser -Represents the IB user currently logged in
type IBUser struct {
	PaperUsername      string `json:"PAPER_USER_NAME"`
	LoginType          int    `json:"loginType"`
	IsPendingApplicant bool   `json:"IS_PENDING_APPLICANT"`
	SFEnabled          bool   `json:"SF_ENABLED"`
	AuthTime           int64  `json:"AUTH_TIME"`
	Credential         string `json:"CREDENTIAL"`
	SFConfig           string `json:"SF_CONFIG"`
	Username           string `json:"USER_NAME"`
	CredentialType     int    `json:"CREDENTIAL_TYPE"`
	Result             bool   `json:"RESULT"`
	IsFreeTrial        bool   `json:"IS_FREE_TRIAL"`
	IP                 string `json:"IP"`
	UserID             int    `json:"USER_ID"`
	Expires            int64  `json:"EXPIRES"`
	Token              string `json:"TOKEN"`
	IsGw               bool   `json:"isGw"`
}

//IBFutContract - Contracts returned for symbols searched
type IBFutContract struct {
	Symbol          string `json:"symbol"`
	Conid           int    `json:"conid"`
	UnderlyingConid int    `json:"underlyingConid"`
	ExpirationDate  int    `json:"expirationDate"`
	Ltd             int    `json:"ltd"`
}

//IBFutContractList - A list of futures contract
type IBFutContractList map[string][]IBFutContract

//IBOrder - struct to prepare an order
type IBOrder struct {
	AcctID          string `json:"acctId"`
	Conid           int    `json:"conid"`
	SecType         string `json:"secType"`
	COID            string `json:"cOID"`
	ParentID        string `json:"parentId"`
	OrderType       string `json:"orderType"`
	ListingExchange string `json:"listingExchange"`
	OutsideRTH      bool   `json:"outsideRTH"`
	Price           int    `json:"price"`
	Side            string `json:"side"`
	Ticker          string `json:"ticker"`
	Tif             string `json:"tif"`
	Referrer        string `json:"referrer"`
	Quantity        int    `json:"quantity"`
	UseAdaptive     bool   `json:"useAdaptive"`
}

//IBOrderReply - Reply struct to order information
type IBOrderReply []struct {
	ID           string   `json:"id"`
	Message      []string `json:"message"`
	OrderID      string   `json:"order_id"`
	LocalOrderID string   `json:"local_order_id"`
	OrderStatus  string   `json:"order_status"`
}

//IBTradeAccount - Gets information about the current trading account
type IBTradeAccount struct {
	SelectedAccount    string `json:"selectedAccount"`
	TradingPermissions struct {
	} `json:"tradingPermissions"`
	AllowFeatures struct {
		ShowGFIS               bool `json:"showGFIS"`
		AllowFXConv            bool `json:"allowFXConv"`
		AllowTypeAhead         bool `json:"allowTypeAhead"`
		SnapshotRefreshTimeout int  `json:"snapshotRefreshTimeout"`
		LiteUser               bool `json:"liteUser"`
		ShowWebNews            bool `json:"showWebNews"`
		Research               bool `json:"research"`
		DebugPnl               bool `json:"debugPnl"`
		ShowTaxOpt             bool `json:"showTaxOpt"`
	} `json:"allowFeatures"`
	ServerInfo struct {
		ServerName    string `json:"serverName"`
		ServerVersion string `json:"serverVersion"`
	} `json:"serverInfo"`
}

//IBLiveOrders - List of live orders
type IBLiveOrders struct {
	Orders []struct {
		Acct               string  `json:"acct"`
		Exchange           string  `json:"exchange"`
		Conid              int     `json:"conid"`
		OrderID            int     `json:"orderId"`
		CashCcy            string  `json:"cashCcy"`
		SizeAndFills       string  `json:"sizeAndFills"`
		OrderDesc          string  `json:"orderDesc"`
		Description1       string  `json:"description1"`
		Ticker             string  `json:"ticker"`
		SecType            string  `json:"secType"`
		ListingExchange    string  `json:"listingExchange"`
		RemainingQuantity  float64 `json:"remainingQuantity"`
		FilledQuantity     float64 `json:"filledQuantity"`
		CompanyName        string  `json:"companyName"`
		Status             string  `json:"status"`
		AvgPrice           string  `json:"avgPrice"`
		OrigOrderType      string  `json:"origOrderType"`
		SupportsTaxOpt     string  `json:"supportsTaxOpt"`
		LastExecutionTime  string  `json:"lastExecutionTime"`
		LastExecutionTimeR int64   `json:"lastExecutionTime_r"`
		OrderType          string  `json:"orderType"`
		OrderRef           string  `json:"order_ref"`
		Side               string  `json:"side"`
		TimeInForce        string  `json:"timeInForce"`
		BgColor            string  `json:"bgColor"`
		FgColor            string  `json:"fgColor"`
	} `json:"orders"`
}
