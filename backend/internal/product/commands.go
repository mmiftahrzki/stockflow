package product

import "strings"

type createCommand struct {
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func (cmd *createCommand) Validate() error {
	if strings.TrimSpace(cmd.SKU) == "" {
		return errProductSKUEmpty
	}

	if strings.TrimSpace(cmd.Name) == "" {
		return errProductNameEmpty
	}

	return nil
}
