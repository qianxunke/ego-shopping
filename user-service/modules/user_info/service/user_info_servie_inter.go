package service

import (
	"crypto/md5"
	"ego-user-service/clients"
	"errors"
	"fmt"
	userInfoProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/user/user_info"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func loginByUserName(req *userInfoProto.InDoneUserLogin, rsp *userInfoProto.OutDoneUserLogin) {
	if rsp == nil {
		rsp = &userInfoProto.OutDoneUserLogin{Token: ""}
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusBadRequest,
			Message: "[loginByUserName] Enter parameter is nil",
		}
		return
	}
	if len(req.UserName) == 0 || len(req.Password) == 0 {
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusBadRequest,
			Message: "用户名或密码为空！",
		}
		return
	}
	rsp.UserInf = &userInfoProto.UserInf{}
	DB := db.MasterEngine()
	err := DB.Table("user_infs").Where("user_name = ?", req.UserName).Scan(&rsp.UserInf).Error
	if err != nil || len(rsp.UserInf.UserId) <= 0 {
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusBadRequest,
			Message: "用户不存在！",
		}
		return
	}
	md5str, err := getMd5Password(req.Password)
	if err != nil {
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	if strings.Compare(md5str, rsp.UserInf.Password) != 0 {
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusBadRequest,
			Message: "密码不正确",
		}
		return
	}
	//登陆成功
	rsp.Error = &userInfoProto.Error{
		Code: http.StatusOK,
	}
	rsp.UserInf.Password = ""
	return
}

func loginByTelephone(req *userInfoProto.InDoneUserLogin, rsp *userInfoProto.OutDoneUserLogin) (isOk bool) {
	if rsp == nil {
		rsp = &userInfoProto.OutDoneUserLogin{Token: ""}
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusInternalServerError,
			Message: "[loginByUserName] Enter parameter is nil",
		}
		isOk = false
		return
	}
	if len(req.MobilePhone) == 0 || len(req.VerificationCode) == 0 {
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusBadRequest,
			Message: "用户名或密码为空!",
		}
		isOk = false
		return
	}
	//判断该手机验证码是否还有效
	err := verificationTelphone(req.MobilePhone)
	if err != nil {
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusBadRequest,
			Message: "验证码失效，请重新获取",
		}
		isOk = false
		return
	}
	rsp.UserInf = &userInfoProto.UserInf{}
	DB := db.MasterEngine()
	err = DB.Table("user_infs").Where("mobile_phone = ?", req.MobilePhone).Scan(&rsp.UserInf).Error
	if err != nil || len(rsp.UserInf.UserId) <= 0 {
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusBadRequest,
			Message: "用户不存在！",
		}
		isOk = false
		return
	}
	//登陆成功
	rsp.Error = &userInfoProto.Error{
		Code:    http.StatusOK,
		Message: "登陆成功",
	}
	isOk = false
	rsp.UserInf.Password = ""
	return
}

func verificationTelphone(telephone string) (err error) {
	if len(telephone) == 0 {
		err = errors.New("电话号码为空")
		return
	}
	//从redis 获取验证码
	code, err := clients.GetRedis().Do("GET", telephone).Int64()
	if err != nil || code <= 0 {
		//将redis令牌清除
		err = errors.New("验证码已过期")
		return
	}

	return
}

func getMd5Password(password string) (md5str string, err error) {
	//将获取密码MD5值
	w := md5.New()
	_, err = io.WriteString(w, password) //将str写入到w中
	if err != nil {
		return
	}
	md5str = fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return
}

func hideUserPrivate(user *userInfoProto.UserInf) (err error) {
	if user == nil {
		err = errors.New("入参为空！")
		return
	}
	user.Password = ""
	if user.IdentityCardType == 1 && len(user.IdentityCardNo) >= 18 {
		user.IdentityCardNo = strings.ReplaceAll(user.IdentityCardNo, user.IdentityCardNo[6:14], "*********")
	}
	if len(user.MobilePhone) >= 11 {
		user.MobilePhone = strings.ReplaceAll(user.MobilePhone, user.MobilePhone[4:7], "****")
	}
	return nil
}

func sendVerificationCode(phone string, ex int64) (err error) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	int6, err := strconv.ParseInt(vcode, 10, 64)
	if err != nil {
		return
	}
	log.Println("Send code : " + vcode)
	//往redis写入验证码
	err = clients.GetRedis().Do("SET", phone, int6, "Ex", ex).Err()
	return
}
