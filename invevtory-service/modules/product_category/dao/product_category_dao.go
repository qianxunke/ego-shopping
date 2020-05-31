package dao

import (
	"fmt"
	productCategoryProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_category"
	"sync"
)

var (
	dao *daoIml
	m   sync.Mutex
)

type daoIml struct {
}

type ProductDao interface {
	FindById(id int64) (product *productCategoryProto.ProductCategory, err error)

	Insert(product *productCategoryProto.ProductCategory) (err error)

	SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *productCategoryProto.Out_GetProductCategorys, err error)

	Delete(ids []int64) (err error)

	Update(id int64, reqMap map[string]interface{}) (err error)

	//查询多等级信息
	GetProductCategoryDetailsList(limit int64, pages int64, order string) (rsp *productCategoryProto.Out_GetProductCategoryDetailsList, err error)
}

func GetDao() (ProductDao, error) {
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
