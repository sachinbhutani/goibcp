package goibcp

//IBClient - Represents the IB API client which executes all API functions
type IBClient struct {
	UserID          int      `json:"userId"`
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

//IBPortfolioAccounts - Accounts
type IBPortfolioAccounts []struct {
	ID             string `json:"id"`
	AccountID      string `json:"accountId"`
	AccountVan     string `json:"accountVan"`
	AccountTitle   string `json:"accountTitle"`
	DisplayName    string `json:"displayName"`
	AccountAlias   string `json:"accountAlias"`
	AccountStatus  int64  `json:"accountStatus"`
	Currency       string `json:"currency"`
	Type           string `json:"type"`
	TradingType    string `json:"tradingType"`
	Faclient       bool   `json:"faclient"`
	ClearingStatus string `json:"clearingStatus"`
	Parent         string `json:"parent"`
	Desc           string `json:"desc"`
	Covestor       bool   `json:"covestor"`
}

//IBPositions - current open positions
type IBPortfolioPositions []struct {
	AcctID        string   `json:"acctId"`
	Conid         int      `json:"conid"`
	ContractDesc  string   `json:"contractDesc"`
	AssetClass    string   `json:"assetClass"`
	Position      float64  `json:"position"`
	MktPrice      float64  `json:"mktPrice"`
	MktValue      float64  `json:"mktValue"`
	Currency      string   `json:"currency"`
	AvgCost       float64  `json:"avgCost"`
	AvgPrice      float64  `json:"avgPrice"`
	RealizedPnl   float64  `json:"realizedPnl"`
	UnrealizedPnl float64  `json:"unrealizedPnl"`
	Exchs         string   `json:"exchs"`
	Expiry        string   `json:"expiry"`
	PutOrCall     string   `json:"putOrCall"`
	Multiplier    int      `json:"multiplier"`
	Strike        float64  `json:"strike"`
	ExerciseStyle string   `json:"exerciseStyle"`
	UndConid      int      `json:"undConid"`
	ConExchMap    []string `json:"conExchMap"`
	Model         string   `json:"model"`
}

//IBTickle - reply recieved from server when tickled
type IBTickle struct {
	SsoExpires int    `json:"ssoExpires"`
	Collission bool   `json:"collission"`
	UserID     int    `json:"userId"`
	Session    string `json:"session"`
	Iserver    struct {
		Error      string `json:"error"`
		Tickle     bool   `json:"tickle"`
		AuthStatus struct {
			Authenticated bool   `json:"authenticated"`
			Competing     bool   `json:"competing"`
			Connected     bool   `json:"connected"`
			Message       string `json:"message"`
			MAC           string `json:"MAC"`
		} `json:"authStatus"`
	} `json:"iserver"`
}

//IBTrades - array of trades from trades endpoint
type IBTrades []struct {
	ExecutionID          string  `json:"execution_id"`
	Symbol               string  `json:"symbol"`
	Side                 string  `json:"side"`
	OrderDescription     string  `json:"order_description"`
	TradeTime            string  `json:"trade_time"`
	TradeTimeR           int64   `json:"trade_time_r"`
	Size                 float64 `json:"size"`
	Price                string  `json:"price"`
	Submitter            string  `json:"submitter"`
	Exchange             string  `json:"exchange"`
	Comission            string  `json:"comission"`
	NetAmount            float64 `json:"net_amount"`
	Account              string  `json:"account"`
	CompanyName          string  `json:"company_name"`
	ContractDescription1 string  `json:"contract_description_1"`
	SecType              string  `json:"sec_type"`
	Conidex              string  `json:"conidex"`
	Position             string  `json:"position"`
	ClearingID           string  `json:"clearing_id"`
	ClearingName         string  `json:"clearing_name"`
	OrderRef             string  `json:"order_ref"`
}

//IBLogout - struct for information recieved with logout endpoint
type IBLogout struct {
	Confirmed bool `json:"confirmed"`
}

//IBAccountLedger - account ledger in base currency
type IBAccountLedger struct {
	BASE struct {
		Commoditymarketvalue      float64 `json:"commoditymarketvalue"`
		Futuremarketvalue         float64 `json:"futuremarketvalue"`
		Settledcash               float64 `json:"settledcash"`
		Exchangerate              int64   `json:"exchangerate"`
		Sessionid                 int64   `json:"sessionid"`
		Cashbalance               float64 `json:"cashbalance"`
		Corporatebondsmarketvalue float64 `json:"corporatebondsmarketvalue"`
		Warrantsmarketvalue       float64 `json:"warrantsmarketvalue"`
		Netliquidationvalue       float64 `json:"netliquidationvalue"`
		Interest                  int64   `json:"interest"`
		Unrealizedpnl             float64 `json:"unrealizedpnl"`
		Stockmarketvalue          float64 `json:"stockmarketvalue"`
		Moneyfunds                float64 `json:"moneyfunds"`
		Currency                  string  `json:"currency"`
		Realizedpnl               float64 `json:"realizedpnl"`
		Funds                     float64 `json:"funds"`
		Acctcode                  string  `json:"acctcode"`
		Issueroptionsmarketvalue  float64 `json:"issueroptionsmarketvalue"`
		Key                       string  `json:"key"`
		Timestamp                 int64   `json:"timestamp"`
		Severity                  int64   `json:"severity"`
		Stockoptionmarketvalue    float64 `json:"stockoptionmarketvalue"`
		Tbondsmarketvalue         float64 `json:"tbondsmarketvalue"`
		Futureoptionmarketvalue   float64 `json:"futureoptionmarketvalue"`
		Cashbalancefxsegment      float64 `json:"cashbalancefxsegment"`
		Secondkey                 string  `json:"secondkey"`
		Tbillsmarketvalue         float64 `json:"tbillsmarketvalue"`
		Dividends                 float64 `json:"dividends"`
	} `json:"BASE"`
}
