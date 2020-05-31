package product_sku_stock

import (
	"inventory-service/modules/product_sku_stock/dao"
	"inventory-service/modules/product_sku_stock/service"
)

func Init() {
	dao.Init()
	service.Init()
}
