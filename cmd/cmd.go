package cmd

import (
	"github.com/lucasgomide/menyoo-api/types"
)

func NewCmdProduct(store types.Store) *CmdProduct {
	return &CmdProduct{store}
}

func NewCmdOrder(store types.Store) *CmdOrder {
	return &CmdOrder{store}
}

func NewCmdProductOrder(store types.Store) *CmdProductOrder {
	return &CmdProductOrder{store}
}
