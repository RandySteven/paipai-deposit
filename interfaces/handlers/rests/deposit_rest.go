package rest_interfaces

import "net/http"

type IDepositRest interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	AccountInquiry(w http.ResponseWriter, r *http.Request)
	BalanceInquiry(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
	Capture(w http.ResponseWriter, r *http.Request)
	TransactionHistory(w http.ResponseWriter, r *http.Request)
	TransactionDetail(w http.ResponseWriter, r *http.Request)
}
