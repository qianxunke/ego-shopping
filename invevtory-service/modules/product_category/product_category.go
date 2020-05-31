package product_category

import (
	"inventory-service/modules/product_category/dao"
	"inventory-service/modules/product_category/service"
)

func Init() {
	dao.Init()
	service.Init()
}
