package user_level

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qianxunke/ego-shopping/ego-common-protos/out/user_level"
	"google.golang.org/grpc"
	"user-api/common/api_common"
)

func Init(client *grpc.ClientConn) *ApiService {
	return &ApiService{
		serviceClient: user_level.NewUserLevelClient(client),
	}
}

type ApiService struct {
	serviceClient user_level.UserLevelClient
}

//获取某个用户的收件地址列表
func (apiService *ApiService) GetUserLevels(c *gin.Context) {
	var reqParameter user_level.In_GetUserLevels
	rsp, _ := apiService.serviceClient.GetUserLevels(context.TODO(), &reqParameter)
	//返回结果
	api_common.SrvResultListDone(c, rsp.UserLevelInfList, rsp.Limit, rsp.Pages, rsp.Total, &api_common.Error{Code: rsp.Error.Code, Message: rsp.Error.Message})
}
