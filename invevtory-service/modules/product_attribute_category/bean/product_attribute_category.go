package bean

import (
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute_category"
	"reflect"
	"strings"
)

type ProductAttributeCategory struct {
	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AttributeCount int64  `protobuf:"varint,3,opt,name=attribute_count,json=attributeCount,proto3" json:"attribute_count,omitempty"`
	ParamCount     int64  `protobuf:"varint,4,opt,name=param_count,json=paramCount,proto3" json:"param_count,omitempty"`
	CreatedTime    string `protobuf:"bytes,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func ProToStandard(in *product_attribute_category.ProductAttributeCategory, out *ProductAttributeCategory) (err error) {
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
