package service

import (
	"context"
	"ego-user-service/modules/user_info"
)

type UserInfoService struct {
}

//用户登陆
func (s *UserInfoService) DoneUserLogin(context.Context, *user_info.InDoneUserLogin) (*user_info.OutDoneUserLogin, error) {

	return &user_info.OutDoneUserLogin{Token: "123", UserInf: nil, Error: nil}, nil
}

//获取用户详细信息
func (s *UserInfoService) GetUserInfo(context.Context, *user_info.InGetUserInfo) (*user_info.OutGetUserInfo, error) {

	return &user_info.OutGetUserInfo{Error: nil, UserInf: &user_info.UserInf{NikeName: "千寻客"}}, nil

}

//获取用户列表
func (s *UserInfoService) GetUserInfoList(context.Context, *user_info.InGetUserInfoList) (*user_info.OutGetUserInfoList, error) {

	return &user_info.OutGetUserInfoList{Error: nil, Pages: 0, Limit: 10, Total: 100}, nil

}

//修改用户信息
func (s *UserInfoService) UpdateUserInfo(context.Context, *user_info.InUpdateUserInfo) (*user_info.OutUpdateUserInfo, error) {

	return &user_info.OutUpdateUserInfo{Error: nil,UserInf: &user_info.UserInf{Password: "123456"}}, nil

}

//用户注册
func (s *UserInfoService) DoneUserRegister(context.Context, *user_info.InDoneUserRegister) (*user_info.OutDoneUserRegister, error) {

	return &user_info.OutDoneUserRegister{Error: nil,UserInf: &user_info.UserInf{NikeName: "DoneUserRegister"}},nil

}

//获取验证码
func (s *UserInfoService) GetVerificationCode(context.Context, *user_info.InGetVerificationCode) (*user_info.OutGetVerificationCode, error) {

	return &user_info.OutGetVerificationCode{Error: nil},nil

}
