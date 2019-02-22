package models

// DetailOrder struct
type DetailOrder struct {
	OrderID  uint
	Product  Product
	Quantity int
}

// GetSubTotal detailed product
func (detail *DetailOrder) GetSubTotal() float64 {
	return (float64(detail.Quantity) * detail.Product.Price)
}
