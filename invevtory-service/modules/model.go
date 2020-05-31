package model

import (
	"inventory-service/modules/brand"
	"inventory-service/modules/product"
	"inventory-service/modules/product_attribute"
	"inventory-service/modules/product_attribute_category"
	"inventory-service/modules/product_category"
	"inventory-service/modules/product_sku_stock"
)

func Init() {
	product.Init()
	brand.Init()
	product_category.Init()
	product_attribute_category.Init()
	product_attribute.Init()
	product_sku_stock.Init()
}
