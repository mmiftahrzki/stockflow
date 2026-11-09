package order

type item struct {
	SKU string
	Qty uint
}

type commandCreate struct {
	Customer string
	Items    []item
}

type commandCancel struct {
}

func (cmd *commandCreate) Validate() error {
	return nil
}
