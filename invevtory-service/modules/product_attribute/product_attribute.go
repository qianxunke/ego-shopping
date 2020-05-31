package product_attribute

import (
	"inventory-service/modules/product_attribute/dao"
	"inventory-service/modules/product_attribute/service"
)

func Init() {
	dao.Init()
	service.Init()
}
