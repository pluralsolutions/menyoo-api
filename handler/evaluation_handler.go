package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"../schema"
	"../types"
)

type EvaluationHandler struct {
	types.EvaluationCmd
}

func (cmd EvaluationHandler) Create(w http.ResponseWriter, r *http.Request) {
	uID := r.Header.Get("uid")
	if uID == "" {
		UnauthorizedRequest(w)
		return
	}

	muxParams := mux.Vars(r)

	restaurantID, err := strconv.Atoi(muxParams["restaurant_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	productOrderID, err := strconv.Atoi(muxParams["product_order_id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	if restaurantID <= 0 || productOrderID <= 0 {
		badRequest(w, missingParamsError())
		return
	}

	body := schema.Evaluation{ProductOrderID: productOrderID, UserID: uID}

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		badRequest(w, err)
		return
	}

	evaluation, err := cmd.EvaluationCmd.CreateEvaluation(body)

	if err != nil {
		badRequest(w, err)
		return
	}

	renderSuccess(w, http.StatusCreated, evaluation)
}
