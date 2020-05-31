package bean

import (
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"reflect"
	"strings"
)

type ProductAttributeValue struct {
	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId          string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductAttributeId int64  `protobuf:"varint,3,opt,name=product_attribute_id,json=productAttributeId,proto3" json:"product_attribute_id,omitempty"`
	Value              string `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func ProToStandard(in *product.ProductAttributeValue, out *ProductAttributeValue) (err error) {
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
