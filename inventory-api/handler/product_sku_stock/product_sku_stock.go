package product_sku_stock

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_sku_stock"
	"github.com/qianxunke/ego-shopping/ego-common-utils/api_common"
	"google.golang.org/grpc"
)

func Init(client *grpc.ClientConn) *ApiService {
	return &ApiService{
		serviceClient: product_sku_stock.NewProductSkuStockHandlerClient(client),
	}
}

type ApiService struct {
	serviceClient product_sku_stock.ProductSkuStockHandlerClient
}

//获取列表
func (this *ApiService) GetProductSkuStocks(c *gin.Context) {
	reqParameter := &product_sku_stock.In_GetProductSkuStocks{}
	reqParameter.ProductId = c.Param("product_id")
	rsp, _ := this.serviceClient.GetProductSkuStocks(context.TODO(), reqParameter)
	api_common.SrvResultListDone(c, rsp.ProductSkuStockList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}
