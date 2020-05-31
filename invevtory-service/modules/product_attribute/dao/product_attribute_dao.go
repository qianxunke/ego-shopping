package dao

import (
	"fmt"
	productAttributeProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute"
	"sync"
)

var (
	dao *daoIml
	m   sync.Mutex
)

type daoIml struct {
}

type Dao interface {
	FindById(id int64) (product *productAttributeProto.ProductAttribute, err error)

	Insert(product *productAttributeProto.ProductAttribute) (err error)

	SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *productAttributeProto.Out_GetProductAttributes, err error)

	Delete(ids []int64) (err error)

	Update(id int64, reqMap map[string]interface{}) (err error)

	GetProductAttributeList(product_attribute_category_id int64, type_value int64) (rsp *productAttributeProto.Out_GetProductAttributeList, err error)
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
