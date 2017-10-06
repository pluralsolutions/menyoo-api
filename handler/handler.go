package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/plural-solutions/menyoo-api/types"
)

type errorBadRequest struct {
	Err string `json:"error"`
}

func NewProductsHandler(cmd types.ProductCmd) *ProductHandler {
	return &ProductHandler{cmd}
}

func NewOrdersHandler(cmd types.OrderCmd) *OrderHandler {
	return &OrderHandler{cmd}
}

func NewProductOrdersHandler(cmd types.ProductOrderCmd) *ProductOrderHandler {
	return &ProductOrderHandler{cmd}
}

func NewEvaluationsHandler(cmd types.EvaluationCmd) *EvaluationHandler {
	return &EvaluationHandler{cmd}
}

func badRequest(w http.ResponseWriter, err error) {
	body := errorBadRequest{err.Error()}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&body)
}

func UnauthorizedRequest(
	w http.ResponseWriter,
) {
	w.WriteHeader(http.StatusUnauthorized)
}

func renderSuccess(
	w http.ResponseWriter,
	status int,
	body interface{},
) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func missingParamsError() error {
	return errors.New("You missing some parameters")
}
