package product

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"github.com/qianxunke/ego-shopping/ego-common-utils/api_common"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

func Init(client *grpc.ClientConn) *ApiService {
	return &ApiService{
		serviceClient: product.NewProductHandlerClient(client),
	}
}

type ApiService struct {
	serviceClient product.ProductHandlerClient
}

func (userApiService *ApiService) GetProduct(c *gin.Context) {
	req := &product.In_GetProductById{}
	req.Id = c.Param("id")
	if len(req.Id) <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api_common.Error{Code: http.StatusBadRequest, Message: "参数非法"})
		return
	}
	//调用后台服务
	rsp, _ := userApiService.serviceClient.GetProductById(context.TODO(), req)
	//返回结果
	api_common.SrvResultDone(c, rsp.Product, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

func (userApiService *ApiService) UpdateProduct(c *gin.Context) {
	req := &product.In_UpdateProductInfo{}
	Id := c.Param("id")
	if len(Id) <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api_common.Error{Code: http.StatusBadRequest, Message: "参数非法"})
		return
	}
	if err := c.ShouldBindJSON(&req.Product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api_common.Error{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	//调用后台服务
	rsp, _ := userApiService.serviceClient.UpdateProductInfo(context.TODO(), req)
	//返回结果
	api_common.SrvResultDone(c, rsp.Product, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

//获取商品列表
func (userApiService *ApiService) GetProducts(c *gin.Context) {
	reqParameter := &product.In_GetProducts{}
	reqParameter.Limit, _ = strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	reqParameter.Pages, _ = strconv.ParseInt(c.DefaultQuery("pages", "1"), 10, 64)
	reqParameter.SearchKey = c.DefaultQuery("search_key", "")
	reqParameter.StartTime = c.DefaultQuery("start_time", "")
	reqParameter.EndTime = c.DefaultQuery("end_time", "")
	rsp, _ := userApiService.serviceClient.GetProducts(context.TODO(), reqParameter)
	//返回结果
	api_common.SrvResultListDone(c, rsp.ProductList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})

}

//获取商品列表
func (userApiService *ApiService) GetProductsByEs(c *gin.Context) {
	/*
		reqParameter := &productEs.In_GetProducts{}
		reqParameter.Limit, _ = strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
		reqParameter.Pages, _ = strconv.ParseInt(c.DefaultQuery("pages", "1"), 10, 64)
		reqParameter.SearchKey = c.DefaultQuery("search_key", "")
		reqParameter.StartTime = c.DefaultQuery("start_time", "")
		reqParameter.EndTime = c.DefaultQuery("end_time", "")
		rsp, _ := myClients.ProductEsClient.GetProducts(context.TODO(), reqParameter)
		//返回结果
		api_common.SrvResultListDone(c, rsp.ProductList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})

	*/

}

func (userApiService *ApiService) AddProduct(c *gin.Context) {
	req := &product.In_CreateProduct{}
	req.ProductDetail = &product.ProductDetails{}
	if err := c.ShouldBindJSON(&req.ProductDetail); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api_common.Error{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	rsp, _ := userApiService.serviceClient.CreateProduct(context.TODO(), req)
	api_common.SrvResultDone(c, nil, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

func (userApiService *ApiService) DeleteProducts(c *gin.Context) {
	req := &product.In_DeleteProducts{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api_common.Error{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	rsp, _ := userApiService.serviceClient.DeleteProducts(context.TODO(), req)
	api_common.SrvResultDone(c, nil, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}
