package product_attribute_category

import (
	"inventory-service/modules/product_attribute_category/dao"
	"inventory-service/modules/product_attribute_category/service"
)

func Init() {
	dao.Init()
	service.Init()
}
