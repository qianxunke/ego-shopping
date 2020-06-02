package product_attribute

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute"
	"github.com/qianxunke/ego-shopping/ego-common-utils/api_common"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func Init(client *grpc.ClientConn) *ApiService {
	return &ApiService{
		serviceClient: product_attribute.NewProductAttributeHandlerClient(client),
	}
}

type ApiService struct {
	serviceClient product_attribute.ProductAttributeHandlerClient
}

func (userApiService *ApiService) Get(c *gin.Context) {
	req := &product_attribute.In_GetProductAttributeById{}
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("[Api] 请求参数不合法！"))
		return
	}
	//调用后台服务
	rsp, _ := userApiService.serviceClient.GetProductAttributeById(context.TODO(), req)
	//返回结果
	api_common.SrvResultDone(c, rsp.ProductAttribute, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

//获取列表
func (userApiService *ApiService) List(c *gin.Context) {
	reqParameter := &product_attribute.In_GetProductAttributes{}
	reqParameter.Limit, _ = strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	reqParameter.Pages, _ = strconv.ParseInt(c.DefaultQuery("pages", "1"), 10, 64)
	reqParameter.SearchKey = c.DefaultQuery("search_key", "")
	reqParameter.StartTime = c.DefaultQuery("start_time", "")
	reqParameter.EndTime = c.DefaultQuery("end_time", "")
	rsp, _ := userApiService.serviceClient.GetProductAttributes(context.TODO(), reqParameter)
	//返回结果
	api_common.SrvResultListDone(c, rsp.ProductAttributeList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})

}

//获取列表
func (userApiService *ApiService) ListInfo(c *gin.Context) {
	reqParameter := &product_attribute.In_GetProductAttributeList{}
	reqParameter.Cid, _ = strconv.ParseInt(c.Param("cid"), 10, 64)
	reqParameter.TypeValue, _ = strconv.ParseInt(c.DefaultQuery("type", "-1"), 10, 64)
	rsp, _ := userApiService.serviceClient.GetProductAttributeList(context.TODO(), reqParameter)
	//返回结果
	api_common.SrvResultListDone(c, rsp.ProductAttributeList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})

}
