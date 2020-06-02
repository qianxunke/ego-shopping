package dao

import (
	"fmt"
	productEsProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"sync"
)

var (
	dao *productDaoIml
	m   sync.Mutex
)

type productDaoIml struct {
}

type ProductDao interface {
	FindById(id string) (product *productEsProto.Product, err error)

	Insert(product *productEsProto.Product) (err error)

	SimpleQuery(limit int, pages int, key string, startTime string, endTime string, order string) (rsp *productEsProto.Out_GetProducts, err error)

	Delete(ids []string) (err error)

	Update(id string, reqMap map[string]interface{}) (err error)
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
	dao = &productDaoIml{}
}
