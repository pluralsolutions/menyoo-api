package cmd

import (
	"../schema"
	"../types"
)

type CmdEvaluation struct {
	types.Store
}

func (cmd CmdEvaluation) CreateEvaluation(evaluation schema.Evaluation) (schema.Evaluation, error) {
	var err error
	if _, err = cmd.Store.FindProductByUser(evaluation.ProductOrderID, evaluation.UserID); err != nil {
		return evaluation, err
	}

	err = cmd.Store.CreateEvaluation(&evaluation)
	return evaluation, err
}
