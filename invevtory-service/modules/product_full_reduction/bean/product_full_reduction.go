package bean

import (
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"reflect"
	"strings"
)

type ProductFullReduction struct {
	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId   string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	FullPrice   float32 `protobuf:"fixed32,3,opt,name=full_price,json=fullPrice,proto3" json:"full_price,omitempty"`
	ReducePrice int64   `protobuf:"varint,4,opt,name=reduce_price,json=reducePrice,proto3" json:"reduce_price,omitempty"`
}

func ProToStandard(in *product.ProductFullReduction, out *ProductFullReduction) (err error) {
	if in == nil || out == nil {
		err = errors.New("[ProToStandard] in or out is nil")
	}
	m := make(map[string]interface{})
	elem := reflect.ValueOf(in).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		if !strings.Contains(relType.Field(i).Name, "XXX") {
			m[relType.Field(i).Name] = elem.Field(i).Interface()
		}
	}
	//将 map 转换为指定的结构体
	err = mapstructure.Decode(m, out)
	return
}
