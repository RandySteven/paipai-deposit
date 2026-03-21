package rest_handler

import (
	"net/http"
)

// Auth implements [rest_interfaces.IDepositRest].
func (*Deposits) Auth(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
