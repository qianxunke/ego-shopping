package brand

import (
	"inventory-service/modules/brand/dao"
	"inventory-service/modules/brand/service"
)

func Init() {
	dao.Init()
	service.Init()
}
