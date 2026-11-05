package product

import "errors"

var errProductSKUAlreadyExists error = errors.New("SKU number already exists")
var errProductSKUEmpty error = errors.New("SKU is required")
var errProductNameEmpty error = errors.New("Name is required")
