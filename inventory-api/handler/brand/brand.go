package brand

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/brand"
	"github.com/qianxunke/ego-shopping/ego-common-utils/api_common"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func Init(client *grpc.ClientConn) *ApiService {
	return &ApiService{
		serviceClient: brand.NewBrandHandlerClient(client),
	}
}

type ApiService struct {
	serviceClient brand.BrandHandlerClient
}

func (userApiService *ApiService) GetBrand(c *gin.Context) {
	req := &brand.In_GetBrandById{}
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("[Api] 请求参数不合法！"))
		return
	}
	//调用后台服务
	rsp, _ := userApiService.serviceClient.GetBrandById(context.TODO(), req)
	//返回结果
	api_common.SrvResultDone(c, rsp.Brand, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

func (userApiService *ApiService) GetBrands(c *gin.Context) {
	reqParameter := &brand.In_GetBrands{}
	reqParameter.Limit, _ = strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	reqParameter.Pages, _ = strconv.ParseInt(c.DefaultQuery("pages", "1"), 10, 64)
	reqParameter.SearchKey = c.DefaultQuery("search_key", "")
	reqParameter.StartTime = c.DefaultQuery("start_time", "")
	reqParameter.EndTime = c.DefaultQuery("end_time", "")
	rsp, _ := userApiService.serviceClient.GetBrands(context.TODO(), reqParameter)
	//返回结果
	api_common.SrvResultListDone(c, rsp.BrandList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})

}
