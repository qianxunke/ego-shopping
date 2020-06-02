package product_es

import (
	"ego-goods-search-service/modules/product_es/dao"
	"ego-goods-search-service/modules/product_es/service"
)

func Init() {
	dao.Init()
	service.Init()
}
