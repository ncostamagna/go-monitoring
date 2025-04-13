package product

import (
	"errors"
	"fmt"
)

var ErrNameRequired = errors.New("name is required")

type ErrNotFound struct {
	ProductID int
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("product '%d' doesn't exist", e.ProductID)
}
