package bean

import (
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"reflect"
	"strings"
)

type ProductSkuStock struct {
	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId      string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SkuCode        string  `protobuf:"bytes,3,opt,name=sku_code,json=skuCode,proto3" json:"sku_code,omitempty"`
	Price          float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock          int64   `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Sp1            string  `protobuf:"bytes,6,opt,name=sp1,proto3" json:"sp1,omitempty"`
	Sp2            string  `protobuf:"bytes,7,opt,name=sp2,proto3" json:"sp2,omitempty"`
	Sp3            string  `protobuf:"bytes,8,opt,name=sp3,proto3" json:"sp3,omitempty"`
	Pic            string  `protobuf:"bytes,9,opt,name=pic,proto3" json:"pic,omitempty"`
	Sale           int64   `protobuf:"varint,10,opt,name=sale,proto3" json:"sale,omitempty"`
	PromotionPrice float32 `protobuf:"fixed32,11,opt,name=promotion_price,json=promotionPrice,proto3" json:"promotion_price,omitempty"`
	LockStock      int64   `protobuf:"varint,12,opt,name=lock_stock,json=lockStock,proto3" json:"lock_stock,omitempty"`
	LowStock       int64   `protobuf:"varint,13,opt,name=low_stock,json=lowStock,proto3" json:"low_stock,omitempty"`
}

type Sku_out_log struct {
	Id          int64
	ProductId   string //商品id
	MemberId    int64  //
	SellNum     int64
	SkuId       int64
	OrderId     int64
	SkuPrice    float32
	TotalAmount float32
	CreateTime  string
}

func ProToStandard(in *product.ProductSkuStock, out *ProductSkuStock) (err error) {
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
