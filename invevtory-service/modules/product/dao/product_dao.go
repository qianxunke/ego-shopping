package dao

import (
	"fmt"
	productProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"sync"
)

var (
	dao *productDaoIml
	m   sync.Mutex
)

type productDaoIml struct {
}

type ProductDao interface {
	FindById(id string) (product *productProto.ProductDetails, err error)

	Insert(product *productProto.In_CreateProduct) (id string, err error)

	SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *productProto.Out_GetProducts, err error)

	Delete(deleteStatus int64, ids []string) (err error)

	Update(product *productProto.ProductDetails) (err error)
}

func GetProductDao() (ProductDao, error) {
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
	dao = &productDaoIml{}
}
