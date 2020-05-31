package bean

import (
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute"
	"reflect"
	"strings"
)

type ProductAttribute struct {
	Id                         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ProductAttributeCategoryId int64  `protobuf:"varint,2,opt,name=product_attribute_category_id,json=productAttributeCategoryId,proto3" json:"product_attribute_category_id"`
	Name                       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	SelectType                 int64  `protobuf:"varint,4,opt,name=select_type,json=selectType,proto3" json:"select_type"`
	InputType                  int64  `protobuf:"varint,5,opt,name=input_type,json=inputType,proto3" json:"input_type"`
	InputList                  string `protobuf:"bytes,6,opt,name=input_list,json=inputList,proto3" json:"input_list"`
	Sort                       int64  `protobuf:"varint,7,opt,name=sort,proto3" json:"sort"`
	FilterType                 int64  `protobuf:"varint,8,opt,name=filter_type,json=filterType,proto3" json:"filter_type"`
	SearchType                 int64  `protobuf:"varint,9,opt,name=search_type,json=searchType,proto3" json:"search_type"`
	RelatedStatus              int64  `protobuf:"varint,10,opt,name=related_status,json=relatedStatus,proto3" json:"related_status"`
	HandAddStatus              int64  `protobuf:"varint,11,opt,name=hand_add_status,json=handAddStatus,proto3" json:"hand_add_status"`
	Type                       int64  `protobuf:"varint,12,opt,name=type,proto3" json:"type"`
}

func ProToStandard(in *product_attribute.ProductAttribute, out *ProductAttribute) (err error) {
	if in == nil || out == nil {
		err = errors.New("[ProToStandard] in or out is nil")
	}
	m := make(map[string]interface{})
	elem := reflect.ValueOf(&in).Elem()
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
