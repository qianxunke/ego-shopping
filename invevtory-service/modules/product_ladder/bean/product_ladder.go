package bean

import (
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"reflect"
	"strings"
)

type ProductLadder struct {
	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Count     int64   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Discount  float32 `protobuf:"fixed32,4,opt,name=discount,proto3" json:"discount,omitempty"`
	Price     float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
}

func ProToStandard(in *product.ProductLadder, out *ProductLadder) (err error) {
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
	err = mapstructure.Decode(m, out)
	return
}
