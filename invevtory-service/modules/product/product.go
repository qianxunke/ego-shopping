package product

import (
	"inventory-service/modules/product/dao"
	"inventory-service/modules/product/service"
)

func Init() {
	dao.Init()
	service.Init()
}
