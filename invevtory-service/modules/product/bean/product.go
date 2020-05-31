package bean

import (
	"errors"
	"github.com/goinggo/mapstructure"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"reflect"
	"strings"
)

type Product struct {
	Id                         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BrandId                    int64   `protobuf:"varint,2,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	ProductCategoryId          int64   `protobuf:"varint,3,opt,name=product_category_id,json=productCategoryId,proto3" json:"product_category_id,omitempty"`
	FeightTemplateId           int64   `protobuf:"varint,4,opt,name=feight_template_id,json=feightTemplateId,proto3" json:"feight_template_id,omitempty"`
	ProductAttributeCategoryId int64   `protobuf:"varint,5,opt,name=product_attribute_category_id,json=productAttributeCategoryId,proto3" json:"product_attribute_category_id,omitempty"`
	Name                       string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Pic                        string  `protobuf:"bytes,7,opt,name=pic,proto3" json:"pic,omitempty"`
	ProductSn                  string  `protobuf:"bytes,8,opt,name=product_sn,json=productSn,proto3" json:"product_sn,omitempty"`
	DeleteStatus               int64   `protobuf:"varint,9,opt,name=delete_status,json=deleteStatus,proto3" json:"delete_status,omitempty"`
	PublishStatus              int64   `protobuf:"varint,10,opt,name=publish_status,json=publishStatus,proto3" json:"publish_status,omitempty"`
	NewStatus                  int64   `protobuf:"varint,11,opt,name=new_status,json=newStatus,proto3" json:"new_status,omitempty"`
	RecommandStatus            int64   `protobuf:"varint,12,opt,name=recommand_status,json=recommandStatus,proto3" json:"recommand_status,omitempty"`
	VerifyStatus               int64   `protobuf:"varint,13,opt,name=verify_status,json=verifyStatus,proto3" json:"verify_status,omitempty"`
	Sort                       int64   `protobuf:"varint,14,opt,name=sort,proto3" json:"sort,omitempty"`
	Sale                       int64   `protobuf:"varint,15,opt,name=sale,proto3" json:"sale,omitempty"`
	Price                      float32 `protobuf:"fixed32,16,opt,name=price,proto3" json:"price,omitempty"`
	PromotionPrice             float32 `protobuf:"fixed32,17,opt,name=promotion_price,json=promotionPrice,proto3" json:"promotion_price,omitempty"`
	GiftGrowth                 int64   `protobuf:"varint,18,opt,name=gift_growth,json=giftGrowth,proto3" json:"gift_growth,omitempty"`
	GiftPoint                  int64   `protobuf:"varint,19,opt,name=gift_point,json=giftPoint,proto3" json:"gift_point,omitempty"`
	UsePointLimit              int64   `protobuf:"varint,20,opt,name=use_point_limit,json=usePointLimit,proto3" json:"use_point_limit,omitempty"`
	SubTitle                   string  `protobuf:"bytes,21,opt,name=sub_title,json=subTitle,proto3" json:"sub_title,omitempty"`
	Description                string  `protobuf:"bytes,22,opt,name=description,proto3" json:"description,omitempty"`
	OriginalPrice              float32 `protobuf:"fixed32,23,opt,name=original_price,json=originalPrice,proto3" json:"original_price,omitempty"`
	Stock                      int64   `protobuf:"varint,24,opt,name=stock,proto3" json:"stock,omitempty"`
	LowStock                   int64   `protobuf:"varint,25,opt,name=low_stock,json=lowStock,proto3" json:"low_stock,omitempty"`
	Unit                       string  `protobuf:"bytes,26,opt,name=unit,proto3" json:"unit,omitempty"`
	Weight                     float32 `protobuf:"fixed32,27,opt,name=weight,proto3" json:"weight,omitempty"`
	PreviewStatus              int64   `protobuf:"varint,28,opt,name=preview_status,json=previewStatus,proto3" json:"preview_status,omitempty"`
	ServiceIds                 string  `protobuf:"bytes,29,opt,name=service_ids,json=serviceIds,proto3" json:"service_ids,omitempty"`
	Keywords                   string  `protobuf:"bytes,30,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Note                       string  `protobuf:"bytes,31,opt,name=note,proto3" json:"note,omitempty"`
	AlbumPics                  string  `protobuf:"bytes,32,opt,name=album_pics,json=albumPics,proto3" json:"album_pics,omitempty"`
	DetailTitle                string  `protobuf:"bytes,33,opt,name=detail_title,json=detailTitle,proto3" json:"detail_title,omitempty"`
	DetailDesc                 string  `protobuf:"bytes,34,opt,name=detail_desc,json=detailDesc,proto3" json:"detail_desc,omitempty"`
	DetailHtml                 string  `protobuf:"bytes,35,opt,name=detail_html,json=detailHtml,proto3" json:"detail_html,omitempty"`
	DetailMobileHtml           string  `protobuf:"bytes,36,opt,name=detail_mobile_html,json=detailMobileHtml,proto3" json:"detail_mobile_html,omitempty"`
	PromotionStartTime         string  `protobuf:"bytes,37,opt,name=promotion_start_time,json=promotionStartTime,proto3" json:"promotion_start_time,omitempty"`
	PromotionEndTime           string  `protobuf:"bytes,38,opt,name=promotion_end_time,json=promotionEndTime,proto3" json:"promotion_end_time,omitempty"`
	PromotionPerLimit          int64   `protobuf:"varint,39,opt,name=promotion_per_limit,json=promotionPerLimit,proto3" json:"promotion_per_limit,omitempty"`
	PromotionType              int64   `protobuf:"varint,40,opt,name=promotion_type,json=promotionType,proto3" json:"promotion_type,omitempty"`
	BrandName                  string  `protobuf:"bytes,41,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	ProductCategoryName        string  `protobuf:"bytes,42,opt,name=product_category_name,json=productCategoryName,proto3" json:"product_category_name,omitempty"`
	CreatedTime                string  `protobuf:"bytes,43,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdateTime                 string  `protobuf:"bytes,44,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func ProToStandard(in *product.Product, out *Product) (err error) {
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
