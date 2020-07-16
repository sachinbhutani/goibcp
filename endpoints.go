package goibcp

var endpoints = map[string]string{
	"sessionStatus":         "/v1/portal/iserver/auth/status",
	"sessionReauthenticate": "/v1/portal/iserver/reauthenticate",
	"sessionTickle ":        "/tickle",
	"sessionLogout":         "/v1/portal/logout",
	"sessionValidateSSO":    "/v1/portal/sso/validate",
	//Contracts
	"contractFuturesBySymbols": "/v1/portal/trsrv/futures",
	//Order Endpoints
	"orderPlace": "/v1/portal/iserver/account/{accountId}/order",
	"ordersLive": "/v1/portal/iserver/account/orders",
	//accountSelected
	"accountIserver": "/v1/portal/iserver/accounts",
}
