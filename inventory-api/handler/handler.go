package handler

import (
	"ego-inventory-api/handler/brand"
	"ego-inventory-api/handler/product"
	"ego-inventory-api/handler/product_attribute"
	"ego-inventory-api/handler/product_attribute_category"
	"ego-inventory-api/handler/product_category"
	"ego-inventory-api/handler/product_sku_stock"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Init() {

}

//注册路由
func RegiserRouter(sClient *grpc.ClientConn, router *gin.Engine) {
	hystrix.DefaultTimeout = 5000
	inventoryRout := router.Group("/inventory")
	{
		productService := product.Init(sClient)
		productRouter := inventoryRout.Group("/product")
		{
			productRouter.GET("/list", productService.GetProductsByEs)
			productRouter.GET("/details/:id", productService.GetProduct)
			productRouter.POST("/create", productService.AddProduct)
			productRouter.POST("/update/:id", productService.UpdateProduct)
			productRouter.POST("/delete", productService.DeleteProducts)
		}

		brandService := brand.Init(sClient)
		brandRouter := inventoryRout.Group("/brand")
		{
			brandRouter.GET("/list", brandService.GetBrands)
			brandRouter.GET("/", brandService.GetBrand)
		}

		productCategoryService := product_category.Init(sClient)
		productCategoryRouter := inventoryRout.Group("/productCategory")
		{
			productCategoryRouter.GET("/list", productCategoryService.GetProductCategorys)
			productCategoryRouter.GET("/", productCategoryService.GetProductCategory)
			productCategoryRouter.GET("/detailsList", productCategoryService.GetProductCategoryDetails)
		}

		productAttributeCategoryService := product_attribute_category.Init(sClient)
		productAttributeCategoryRouter := inventoryRout.Group("/productAttributeCategory")
		{
			productAttributeCategoryRouter.GET("/list", productAttributeCategoryService.List)
			productAttributeCategoryRouter.GET("/", productAttributeCategoryService.Get)
		}

		productAttributService := product_attribute.Init(sClient)
		productAttributeRouter := inventoryRout.Group("/productAttribute")
		{
			productAttributeRouter.GET("/list/:cid", productAttributService.ListInfo)
			productAttributeRouter.GET("/list", productAttributService.List)
			//	productAttributeRouter.GET("/:id", productAttributService.Get)
		}

		skuStockService := product_sku_stock.Init(sClient)
		skuStockRouter := inventoryRout.Group("/sku")
		{
			skuStockRouter.GET("/:product_id", skuStockService.GetProductSkuStocks)
			//	brandRouter.GET("/", skuStockService.GetBrand)
		}
	}
	//user_level
}
