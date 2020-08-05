# goibcp
Go lang wrapper for interactive brokers Client Portal (CP) web gateway

## Install
    go get -u github.com/sachinbhutani/goibcp

## Usage Example 
    TODO

# Coverage 
### Session
- [X] Tickle
- [x] Logout
- [x] Validate SSO
- [x] Authentication Status
- [X] Reauthenticate
### IBCust
- [ ] Entity Info
### Portfolio Analyst
- [ ] Account Performance 
- [ ] Account Balance Summary
### Account 
- [X] Portfolio Accounts
- [ ] List of Subaccounts 
- [ ] Account Information
- [ ] Account Summary 
- [ ] Account Ledger
- [X] Brokerage Accounts
- [ ] Update Selected Account
- [ ] PnL for selected Account 
### Portfolio 
- [X] Portfolio Accounts
- [ ] List of Subaccounts 
- [ ] Account Information
- [ ] Account Allocation
- [X] Portfolio Positions
- [ ] Position by conid per account
- [ ] Invalidate Cache
- [ ] Account Summary
- [ ] Account Ledger
- [ ] Position by conid 
### Trades
- [ ] List of Trades
### Order 
- [X] Live Order 
- [X] Place Order 
- [X] Place Order (Support Bracket Order) 
- [X] Place Order Reply
- [ ] Preview Order 
- [ ] Modify Order 
- [ ] Delete Order 
### Market Data 
- [ ] Market Data 
- [ ] Market Data Cancel (Single) 
- [ ] Market Data Cancel (All)
- [ ] Market Data History
### Contract
- [ ] Contract Info 
- [ ] Search by Symbol Name
- [ ] Get strikes for options/warrants
- [ ] Get available conids of future/option/warrant/cash/CFD
- [ ] Secdef by Conid
- [ ] Security Futures by Symbol
- [ ] Security Futures by Symbol
### Scanner 
- [ ] get lists of available scanners
- [ ] get lists of available scanners
### PnL
- [ ] PnL for the selected account


# FAQs

## Could not connect to IB CP Gateway, Ensure the CP gateway is running and logged in before connecting
- Ensure the client portal gateway is running at the configured host and destination
- Ensure you have logged in to the portal 
- default link for the portal will be https://localhost:5000/.
- Follow the getting started instructions at https://interactivebrokers.github.io/cpwebapi/index.html

## X509 Certificate error 

Please refer to the interactive brokers CP Web gateway FAQq in common question section at 
https://interactivebrokers.github.io/cpwebapi/faq.html

> Since the gateway is running on your premises the certificate needs to be created/self-signed by you, or officially signed by a 3rd party. The gateway is similar to another webserver such as Tomcat which doesn’t provide a certificate along with the release.

https://www.sslshopper.com/article-how-to-create-a-self-signed-certificate-using-java-keytool.html
