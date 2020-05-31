package dao

import (
	"fmt"
	branProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/brand"
	"sync"
)

var (
	dao *productDaoIml
	m   sync.Mutex
)

type productDaoIml struct {
}

type ProductDao interface {
	FindById(id int64) (product *branProto.Brand, err error)

	Insert(product *branProto.Brand) (err error)

	SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *branProto.Out_GetBrands, err error)

	Delete(ids []int64) (err error)

	Update(id int64, reqMap map[string]interface{}) (err error)
}

func GetBrandDao() (ProductDao, error) {
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
