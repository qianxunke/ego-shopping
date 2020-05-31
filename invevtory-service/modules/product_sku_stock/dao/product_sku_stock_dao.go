package dao

import (
	"fmt"
	productSkuStockProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_sku_stock"
	"sync"
)

var (
	dao *daoIml
	m   sync.Mutex
)

type daoIml struct {
}

type ProductDao interface {
	FindById(id int64) (product *productSkuStockProto.ProductSkuStock, err error)

	Insert(product *productSkuStockProto.ProductSkuStock) (err error)

	SimpleQuery(productId string) (rsp *productSkuStockProto.Out_GetProductSkuStocks, err error)

	Delete(ids []int64) (err error)

	Update(id int64, reqMap map[string]interface{}) (err error)

	UpdateStock(id int64, productId string, memberId int64, stockNum int64) (rsp *productSkuStockProto.SkuOutLog, err error)
}

func GetDao() (ProductDao, error) {
	if dao == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
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
