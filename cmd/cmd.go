package cmd

import (
	"github.com/lucasgomide/menyoo-api/types"
)

func NewCmdProduct(store types.Store) *CmdProduct {
	return &CmdProduct{store}
}
