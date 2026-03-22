package rest_handler

import (
	"context"
	"log"
	"net/http"

	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/enums"
	"github.com/RandySteven/paipai-deposit/utils"
	"github.com/google/uuid"
)

// CreateAccount implements [rest_interfaces.IDepositRest].
func (d *Deposits) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.New().String()
		request = &requests.CreateAccountRequest{}
		dataKey = "data"
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
	)

	if err := utils.BindJSON(r, &request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, "failed to bind request", nil, nil, err)
		return
	}

	log.Println("request create account", request)
	response, appError := d.Usecases.CreateAccount(ctx, request)
	if appError != nil {
		utils.ResponseHandler(w, http.StatusInternalServerError, appError.Error(), nil, nil, appError)
		return
	}

	utils.ResponseHandler(w, http.StatusOK, "success", &dataKey, response, nil)
}
