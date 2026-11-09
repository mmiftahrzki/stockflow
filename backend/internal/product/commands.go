package product

import "strings"

type commandUpdate struct{}
type commandDelete struct{}
type commandCreateStock struct{}
type commandCreate struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func (cmd *commandCreate) Validate() error {
	if strings.TrimSpace(cmd.SKU) == "" {
		return errProductSKUEmpty
	}

	if strings.TrimSpace(cmd.Name) == "" {
		return errProductNameEmpty
	}

	return nil
}
