package user_addr

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/user/user_addr"
	"google.golang.org/grpc"
	"net/http"
	"user-api/common/api_common"
)

func Init(client *grpc.ClientConn) *UserApiService {
	return &UserApiService{
		serviceClient: user_addr.NewUserAddrServiceClient(client),
	}
}

type UserApiService struct {
	serviceClient user_addr.UserAddrServiceClient
}

func (userApiService *UserApiService) Add(c *gin.Context) {
	var req user_addr.In_CreateUserAddr
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("[Api] 请求参数不合法！"))
		return
	}

	if req.UserAddr == nil {
		api_common.SrvResultDone(c, nil, &api_common.Error{Code: http.StatusBadRequest, Message: "paramters cannot is nil"})
		return
	}
	if len(req.UserAddr.Province) == 0 {
		api_common.SrvResultDone(c, nil, &api_common.Error{Code: http.StatusBadRequest, Message: "Province cannot is nil"})
		return
	}
	if len(req.UserAddr.City) == 0 {
		api_common.SrvResultDone(c, nil, &api_common.Error{Code: http.StatusBadRequest, Message: "City cannot is nil"})
		return
	}
	if len(req.UserAddr.Address) == 0 {
		api_common.SrvResultDone(c, nil, &api_common.Error{Code: http.StatusBadRequest, Message: "Address cannot is nil"})
		return
	}
	if len(req.UserAddr.District) == 0 {
		api_common.SrvResultDone(c, nil, &api_common.Error{Code: http.StatusBadRequest, Message: "District cannot is nil"})
		return
	}
	if req.UserAddr.UserId = api_common.GetHeadUserId(c); len(req.UserAddr.UserId) == -1 {
		api_common.SrvResultDone(c, nil, &api_common.Error{Code: http.StatusBadRequest, Message: "Illegal request"})
		return
	}
	//调用后台服务
	rsp, _ := userApiService.serviceClient.CreateUserAddr(context.TODO(), &req)
	//返回结果
	api_common.SrvResultDone(c, rsp.UserAddr, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}

//获取某个用户的收件地址列表
func (userApiService *UserApiService) GetUserAddrs(c *gin.Context) {
	var reqParameter user_addr.In_GetOneUserAddrs
	if err := c.ShouldBindJSON(&reqParameter); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("[Api] 请求参数不合法！"))
		return
	}
	if reqParameter.UserId = api_common.GetHeadUserId(c); len(reqParameter.UserId) == 0 {
		api_common.SrvResultListDone(c, nil, reqParameter.Limit, 0, reqParameter.Pages, &api_common.Error{Code: http.StatusBadRequest, Message: "Illegal request"})
		return
	}
	rsp, _ := userApiService.serviceClient.GetOneUserAddrs(context.TODO(), &reqParameter)
	//返回结果
	api_common.SrvResultListDone(c, rsp.UserAddrList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}
