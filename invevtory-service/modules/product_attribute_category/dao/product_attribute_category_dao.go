package dao

import (
	"fmt"
	productAttributeCategoryProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute_category"
	"sync"
)

var (
	dao *daoIml
	m   sync.Mutex
)

type daoIml struct {
}

type Dao interface {
	FindById(id int64) (product *productAttributeCategoryProto.ProductAttributeCategory, err error)

	Insert(product *productAttributeCategoryProto.ProductAttributeCategory) (err error)

	SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *productAttributeCategoryProto.Out_GetProductAttributeCategorys, err error)

	Delete(ids []int64) (err error)

	Update(id int64, reqMap map[string]interface{}) (err error)
}

func GetDao() (Dao, error) {
	if dao == nil {
		return nil, fmt.Errorf("[GetDao] Dao 未初始化")
	}
	return dao, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()
	if dao != nil {
		return
	}
	dao = &daoIml{}
}
