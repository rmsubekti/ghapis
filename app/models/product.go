package models

import (
	"github.com/rmsubekti/ghapis/app/utils"

	"github.com/jinzhu/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Size        string  `json:"size"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

// Validate product payload
func (product *Product) Validate() (map[string]interface{}, bool) {
	if product.Code == "" {
		return utils.Message(false, "Product code is required"), false
	}
	if product.Name == "" {
		return utils.Message(false, "Product name is required"), false
	}
	if product.Price <= 0 {
		return utils.Message(false, "Product price is required"), false
	}
	if product.Stock <= 0 {
		return utils.Message(false, "Product stock is required"), false
	}

	//All the required parameters are present
	return utils.Message(true, "success"), true
}

// Create new Product
func (product *Product) Create() map[string]interface{} {
	if resp, ok := product.Validate(); !ok {
		return resp
	}
	GetDB().Create(product)
	response := utils.Message(true, "Product has been created")
	response["product"] = product
	return response
}

// GetProduct by id
func GetProduct(ID string) *Product {
	product := &Product{}
	err := GetDB().Table("products").Where("id = ?", ID).First(product).Error
	if err != nil {
		return nil
	}
	return product
}

// GetProducts show all product
func GetProducts() []Product {
	products := []Product{}
	if err := GetDB().Find(&products).Error; err != nil {
		return nil
	}
	return products
}
