package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rmsubekti/ghapis/app/utils"
)

// Order struct
type Order struct {
	gorm.Model
	AccountID    uint
	DetailOrders []DetailOrder `gorm:"foreignkey:OrderID;association_foreignkey:Refer"`
	Total        float64
}

// Create order
func (order *Order) Create() map[string]interface{} {
	var total float64
	for _, v := range order.DetailOrders {
		total += v.GetSubTotal()
	}
	order.Total = total
	GetDB().Create(order)
	response := utils.Message(true, "Order has been created")
	response["order"] = order
	return response
}
