package goibcp

var endpoints = map[string]string{
	//session endpoints
	"sessionStatus":         "/v1/portal/iserver/auth/status",
	"sessionReauthenticate": "/v1/portal/iserver/reauthenticate?force=true",
	"sessionTickle":         "/v1/portal/tickle",
	"sessionLogout":         "/v1/portal/logout",
	"sessionValidateSSO":    "/v1/portal/sso/validate",
	//Contracts
	"contractFuturesBySymbols": "/v1/portal/trsrv/futures",
	//Order Endpoints
	"orderPlace": "/v1/portal/iserver/account/{accountId}/order",
	"ordersLive": "/v1/portal/iserver/account/orders",
	//accountSelected
	"accountIserver": "/v1/portal/iserver/accounts",
	//portfolio
	"portfolioAccounts":  "/v1/portal/portfolio/accounts",
	"portfolioPositions": "/v1/portal/portfolio/{accountId}/positions/{pageId}",
}
