package modules

import (
	brandPro "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/brand"
	productPro "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	productAttributePro "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute"
	"inventory-service/modules/brand"
	m_brand_s "inventory-service/modules/brand/service"
	m_product_s "inventory-service/modules/product/service"
	m_productAttribute_s "inventory-service/modules/product_attribute/service"

	product_attribute_categoryPro "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute_category"
	product_categoryPro "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_category"
	product_sku_stockPro "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_sku_stock"
	m_product_attribute_category_s "inventory-service/modules/product_attribute_category/service"
	m_product_category_s "inventory-service/modules/product_category/service"
	m_product_sku_stock_s "inventory-service/modules/product_sku_stock/service"

	"google.golang.org/grpc"
	"inventory-service/modules/product"
	"inventory-service/modules/product_attribute"
	"inventory-service/modules/product_attribute_category"
	"inventory-service/modules/product_category"
	"inventory-service/modules/product_sku_stock"
	"log"
)

func Init() {
	product.Init()
	brand.Init()
	product_category.Init()
	product_attribute_category.Init()
	product_attribute.Init()
	product_sku_stock.Init()
}

func RegisterService(s *grpc.Server) {
	//brand
	brandS, err := m_brand_s.GetService()
	if err != nil {
		log.Fatal("brand service init error")
		return
	}
	brandPro.RegisterBrandHandlerServer(s, brandS)
	//product
	productS, err := m_product_s.GetService()
	if err != nil {
		log.Fatal("product service init error")
		return
	}
	productPro.RegisterProductHandlerServer(s, productS)
	//productAttribute
	productAttributeS, err := m_productAttribute_s.GetService()
	if err != nil {
		log.Fatal("productAttribute service init error")
		return
	}
	productAttributePro.RegisterProductAttributeHandlerServer(s, productAttributeS)

	product_attribute_categoryS, err := m_product_attribute_category_s.GetService()
	if err != nil {
		log.Fatal("product_attribute_category service init error")
		return
	}
	product_attribute_categoryPro.RegisterProductAttributeCategoryHandlerServer(s, product_attribute_categoryS)
	//product_category
	product_categoryS, err := m_product_category_s.GetService()
	if err != nil {
		log.Fatal("product_category service init error")
		return
	}
	product_categoryPro.RegisterProductCategoryHandlerServer(s, product_categoryS)

	product_sku_stockS, err := m_product_sku_stock_s.GetService()
	if err != nil {
		log.Fatal("product_sku_stock init error")
		return
	}
	product_sku_stockPro.RegisterProductSkuStockHandlerServer(s, product_sku_stockS)

}
