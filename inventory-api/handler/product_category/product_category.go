package product_category

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_category"
	"github.com/qianxunke/ego-shopping/ego-common-utils/api_common"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func Init(client *grpc.ClientConn) *ApiService {
	return &ApiService{
		serviceClient: product_category.NewProductCategoryHandlerClient(client),
	}
}

type ApiService struct {
	serviceClient product_category.ProductCategoryHandlerClient
}

func (userApiService *ApiService) GetProductCategory(c *gin.Context) {
	req := &product_category.In_GetProductCategoryById{}
	if err := c.ShouldBindQuery(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("[Api] 请求参数不合法！"))
		return
	}
	//调用后台服务
	rsp, _ := userApiService.serviceClient.GetProductCategoryById(context.TODO(), req)
	//返回结果
	api_common.SrvResultDone(c, rsp.ProductCategory, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

//获取列表
func (userApiService *ApiService) GetProductCategorys(c *gin.Context) {
	reqParameter := &product_category.In_GetProductCategorys{}
	reqParameter.Limit, _ = strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	reqParameter.Pages, _ = strconv.ParseInt(c.DefaultQuery("pages", "1"), 10, 64)
	reqParameter.SearchKey = c.DefaultQuery("search_key", "")
	reqParameter.StartTime = c.DefaultQuery("start_time", "")
	reqParameter.EndTime = c.DefaultQuery("end_time", "")
	rsp, _ := userApiService.serviceClient.GetProductCategorys(context.TODO(), reqParameter)
	//返回结果
	api_common.SrvResultListDone(c, rsp.ProductCategoryList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

//获取列表
func (userApiService *ApiService) GetProductCategoryDetails(c *gin.Context) {
	reqParameter := &product_category.In_GetProductCategoryDetailsList{}
	rsp, _ := userApiService.serviceClient.GetProductCategoryDetailsList(context.TODO(), reqParameter)
	var rsp_map []ReqMap
	var isExit bool
	for i := 0; i < len(rsp.ProductCategoryList); i++ {
		isExit = false
		for j := 0; j < len(rsp_map); j++ {
			if rsp.ProductCategoryList[i].Id == rsp_map[j].Id {
				isExit = true
				rsp_map[j].Childrens = append(rsp_map[j].Childrens, Children{Id: rsp.ProductCategoryList[i].ChildId, Name: rsp.ProductCategoryList[i].ChildName})
				break
			}
		}
		if !isExit {
			rsp_map = append(rsp_map, ReqMap{Id: rsp.ProductCategoryList[i].Id, Name: rsp.ProductCategoryList[i].Name})
			rsp_map[len(rsp_map)-1].Childrens = append(rsp_map[len(rsp_map)-1].Childrens, Children{Id: rsp.ProductCategoryList[i].ChildId, Name: rsp.ProductCategoryList[i].Name})
		}
	}
	//返回结果
	api_common.SrvResultListDone(c, rsp_map, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

type ReqMap struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Childrens []Children `json:"children,omitempty"`
}
type Children struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
