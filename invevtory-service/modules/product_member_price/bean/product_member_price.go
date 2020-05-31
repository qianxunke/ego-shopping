package bean

import (
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"reflect"
	"strings"
)

type ProductMemberPrice struct {
	Id              int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId       string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	MemberLevelId   int64   `protobuf:"varint,3,opt,name=member_level_id,json=memberLevelId,proto3" json:"member_level_id,omitempty"`
	MemberPrice     float32 `protobuf:"fixed32,4,opt,name=member_price,json=memberPrice,proto3" json:"member_price,omitempty"`
	MemberLevelName string  `protobuf:"bytes,5,opt,name=member_level_name,json=memberLevelName,proto3" json:"member_level_name,omitempty"`
}

func ProToStandard(in *product.ProductMemberPrice, out *ProductMemberPrice) (err error) {
	if out == nil {
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
