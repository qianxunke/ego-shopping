package service

import (
	"context"
	"ego-user-service/modules/access"
	"ego-user-service/modules/user_info/dao"
	"github.com/qianxunke/ego-shopping/ego-common-protos/out/user_info"
	checkutil "ego-user-service/utils"
	"errors"
	"net/http"
	"time"
)



//用户登陆
func (s *userInfoService) DoneUserLogin(ctx context.Context, req *user_info.InDoneUserLogin) (rsp *user_info.OutDoneUserLogin, err error) {

	rsp = &user_info.OutDoneUserLogin{}
	if req.LoginType == 1 { //使用用户名。密码登陆
		loginByUserName(req, rsp)
	} else if req.LoginType == 2 { //验证码登陆
		loginByTelephone(req, rsp)
	} else {
		rsp.Error = &user_info.Error{
			Code:    http.StatusBadRequest,
			Message: "请求非法,未设置登陆方式！",
		}
	}
	if rsp.Error.Code != http.StatusOK {
		rsp.UserInf = nil
		return
	}
	acS, err := access.GetService()
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Token, err = acS.MakeAccessToken(&access.Subject{
		ID:   rsp.UserInf.UserId,
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	_ = hideUserPrivate(rsp.UserInf)

	return rsp, nil
}

//获取用户详细信息
func (s *userInfoService) GetUserInfo(ctx context.Context, req *user_info.InGetUserInfo) (rsp *user_info.OutGetUserInfo, err error) {
	rsp = &user_info.OutGetUserInfo{}
	//获取某个用户的信息
	rsp.UserInf, err = dao.GetUserInfoById(req.UserId)
	if err != nil && len(rsp.UserInf.UserId) == 0 {
		rsp.Error = &user_info.Error{
			Code:    http.StatusBadRequest,
			Message: "没有找到任何数据！",
		}
		return
	}
	rsp.Roles = []string{"admin"}
	_ = hideUserPrivate(rsp.UserInf)
	rsp.Error = &user_info.Error{
		Code:    http.StatusOK,
		Message: "OK",
	}
	return
}

//获取用户列表
func (s *userInfoService) GetUserInfoList(ctx context.Context, req *user_info.InGetUserInfoList) (rsp *user_info.OutGetUserInfoList, err error) {
	//对参数鉴权
	if req.Limit == 0 {
		req.Limit = 10 //默认10个分页
	}
	if req.Limit > 1000 { //每一页数量
		req.Limit = 1000
	}
	if req.Pages <= 0 { //页数
		req.Pages = 1
	}
	rsp = dao.GetUserInfoList(req.SearchKey, req.StartTime, req.EndTime, req.Pages, req.Limit)
	if rsp.Error.Code != http.StatusOK {
		err = errors.New(rsp.Error.Message)
		return
	}

	var message string
	if rsp.Total > 0 && len(rsp.UserInfList) > 0 {
		for i := 0; i < len(rsp.UserInfList); i++ {
			_ = hideUserPrivate(rsp.UserInfList[i])
		}
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &user_info.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	return
}

//修改用户信息
func (s *userInfoService) UpdateUserInfo(ctx context.Context, req *user_info.InUpdateUserInfo) (rsp *user_info.OutUpdateUserInfo, err error) {

	rsp = &user_info.OutUpdateUserInfo{}
	//这里只是修改普通信息
	updataData := map[string]interface{}{}
	if len(req.UserInf.Birthday) > 0 {
		updataData["birthday"] = req.UserInf.Birthday
	}
	if len(req.UserInf.Gender) > 0 {
		updataData["gender"] = req.UserInf.Gender
	}

	if len(req.UserInf.NikeName) > 0 {
		updataData["nike_name"] = req.UserInf.NikeName
	}
	if req.UserInf.IdentityCardType > 0 && len(req.UserInf.IdentityCardNo) > 0 {
		updataData["identity_card_type"] = req.UserInf.IdentityCardType
		updataData["identity_card_no"] = req.UserInf.IdentityCardNo
	}
	if len(updataData) <= 0 {
		rsp.Error = &user_info.Error{
			Code:    http.StatusBadRequest,
			Message: "没有任何数据需要修改！",
		}
		return
	}
	updataData["modified_time"] = time.Now().Format("2006-01-02 15:04:05")
	err = dao.UpdateUserInfo(req.UserInf.UserId, updataData)
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &user_info.Error{
		Code:    http.StatusOK,
		Message: "修改成功！",
	}
	return
}

//用户注册
func (s *userInfoService) DoneUserRegister(ctx context.Context, req *user_info.InDoneUserRegister) (rsp *user_info.OutDoneUserRegister, err error) {
	rsp = &user_info.OutDoneUserRegister{}
	user, err := dao.GetUserInfoByPhoneOrUserName(req.Userinf.UserName, req.Userinf.MobilePhone)
	if err == nil && user != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusBadRequest,
			Message: "用户已存在",
		}
		return
	}
	if len(req.VerificationCode) == 0 {
		rsp.Error = &user_info.Error{
			Code:    http.StatusBadRequest,
			Message: "验证码为空",
		}
		return
	}
	err = verificationTelphone(req.Userinf.MobilePhone)
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return
	}
	md5str, err := getMd5Password(req.Userinf.Password)
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	req.Userinf.Password = md5str
	req.Userinf.UserStats = 1
	req.Userinf.RegisterTime = time.Now().Format("2006-01-02 15:04:05")
	req.Userinf.UserPoint = 0
	req.Userinf.ModifiedTime = req.Userinf.RegisterTime
	req.Userinf.Birthday = req.Userinf.ModifiedTime
	err = dao.Insert(req.Userinf)
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusInternalServerError,
			Message: "创建用户失败",
		}
		return
	}
	rsp.UserInf, err = dao.GetUserInfoByPhoneOrUserName(req.Userinf.MobilePhone, req.Userinf.UserName)
	if err != nil || len(rsp.UserInf.UserId) <= 0 {
		rsp.Error = &user_info.Error{
			Code:    http.StatusInternalServerError,
			Message: "查询用户失败",
		}
		return
	}
	acS, err := access.GetService()
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Token, err = acS.MakeAccessToken(&access.Subject{
		ID:   rsp.UserInf.UserId,
		Name: req.Userinf.UserName,
	})
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusInternalServerError,
			Message: "token 生成出错",
		}
		return
	}
	rsp.Error = &user_info.Error{
		Code:    http.StatusOK,
		Message: "注册成功",
	}
	_ = hideUserPrivate(rsp.UserInf)
	return
}

//获取验证码
func (s *userInfoService) GetVerificationCode(ctx context.Context, req *user_info.InGetVerificationCode) (rsp *user_info.OutGetVerificationCode, err error) {
	rsp = &user_info.OutGetVerificationCode{}
	if req.Telephone == "" {
		rsp.Error = &user_info.Error{
			Message: "电话号码为空",
			Code:    http.StatusBadRequest,
		}
		return
	}

	if checkutil.ValiTephone(req.Telephone) {
		rsp.Error = &user_info.Error{
			Message: "电话号码不合法",
			Code:    http.StatusBadRequest,
		}
		return
	}

	//判断该手机验证码是否还有效
	err = verificationTelphone(req.Telephone)
	if err == nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusBadRequest,
			Message: "验证码已发，请耐心等等，或请一分钟后再次请求！",
		}
		return
	}
	err = sendVerificationCode(req.Telephone, 3000)
	if err != nil {
		rsp.Error = &user_info.Error{
			Code:    http.StatusBadRequest,
			Message: "验证码发送失败，请重试！",
		}
		return
	}
	rsp.Error = &user_info.Error{
		Code:    http.StatusOK,
		Message: "验证码发送成功",
	}
	return
}



