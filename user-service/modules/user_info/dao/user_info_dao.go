package dao

import (
	"ego-user-service/utils/uuid"
	userInfoProto "github.com/qianxunke/ego-shopping/ego-common-protos/out/user_info"
	"errors"
	"fmt"
	"github.com/go-log/log"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
	"net/http"
)

func UserIsExit(userName string) (u *userInfoProto.UserInf, err error) {
	DB := db.MasterEngine()
	u = &userInfoProto.UserInf{}
	err = DB.Table("user_infs").Where("user_name = ?", userName).Scan(&u).Error
	if err != nil {
		return nil, err
	}
	if len(u.UserId) <= 0 {
		return nil, errors.New("user no exit !")
	}
	return u, err
}

func UpdateUserInfo(userId string, userMap map[string]interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("%v", r))
			}
		}
	}()
	DB := db.MasterEngine()
	err = DB.Model(&userInfoProto.UserInf{}).Where("user_id =?", userId).Updates(userMap, true).Error
	return err
}

func GetUserInfoList(searchKey string, startTime string, endTime string, pages int64, limit int64) (rsp *userInfoProto.OutGetUserInfoList) {
	DB := db.MasterEngine()
	rsp = &userInfoProto.OutGetUserInfoList{}
	var err error
	if len(searchKey) == 0 {
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&userInfoProto.UserInf{}).Where("register_time > ?", endTime).Order("user_id desc").Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("register_time > ? ", startTime).Order("user_id desc").Offset((- 1) * limit).Limit(limit).Find(&rsp.UserInfList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&userInfoProto.UserInf{}).Where("register_time < ? ", endTime).Order("user_id desc").Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("register_time < ? ", endTime).Order("user_id desc").Offset((pages - 1) * limit).Limit(limit).Find(&rsp.UserInfList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&userInfoProto.UserInf{}).Where("register_time  between ? and ?", startTime, endTime).Order("user_id desc").Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("register_time  between ? and ?", startTime, endTime).Order("user_id desc").Offset((pages - 1) * limit).Limit(limit).Find(&rsp.UserInfList).Error
			}
		} else {
			//先统计
			err = DB.Model(&userInfoProto.UserInf{}).Order("user_id desc").Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Order("user_id desc").Offset((pages - 1) * limit).Limit(limit).Find(&rsp.UserInfList).Error
			}
		}
	} else {
		key := "%" + searchKey + "%"
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&userInfoProto.UserInf{}).Where("(nike_name like ? or user_name like ? or mobile_phone like ?) and register_time > ? ", key, key, key, startTime).Order("user_id desc").Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Model(&userInfoProto.UserInf{}).Where("(nike_name like ? or user_name like ? or mobile_phone like ?) and register_time > ? ", key, key, key, startTime).Order("user_id desc").Offset((pages - 1) * limit).Limit(limit).Find(&rsp.UserInfList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&userInfoProto.UserInf{}).Where("(nike_name like ? or user_name like ? or mobile_phone like ?) and register_time < ? ", key, key, key, endTime).Order("user_id desc").Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(nike_name like ? or user_name like ? or mobile_phone like ?) and register_time < ? ", key, key, key, endTime).Order("user_id desc").Offset((pages - 1) * limit).Limit(limit).Find(&rsp.UserInfList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&userInfoProto.UserInf{}).Where("(nike_name like ? or user_name like ? or mobile_phone like ?) and register_time between ? and ?", key, key, key, startTime, endTime).Order("user_id desc").Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(nike_name like ? or user_name like ? or mobile_phone like ?) and register_time between ? and ?", key, key, key, startTime, endTime).Order("user_id desc").Offset((pages - 1) * limit).Limit(limit).Find(&rsp.UserInfList).Error
			}
		} else {
			err = DB.Model(&userInfoProto.UserInf{}).Where("1=1 and (nike_name like ? or user_name like ? or mobile_phone like ?)", key, key, key, startTime).Order("user_id desc").Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("1=1 and (nike_name like ? or user_name like ? or mobile_phone like ?)", key, key, key, startTime).Order("user_id desc").Offset((pages - 1) * limit).Limit(limit).Find(&rsp.UserInfList).Error
			}
		}
	}
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &userInfoProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Limit = limit
	rsp.Pages = pages
	return

}

/**
get userinfo by id
*/
func GetUserInfoById(userId string) (user *userInfoProto.UserInf, err error) {
	DB := db.MasterEngine()
	user = &userInfoProto.UserInf{}
	err = DB.Where("user_id = ?", userId).First(&user).Error
	return
}

/**
  user exis?
*/
func GetUserInfoByPhoneOrUserName(mobilePhone string, userName string) (user *userInfoProto.UserInf, err error) {
	DB := db.MasterEngine()
	user = &userInfoProto.UserInf{}
	err = DB.Where(" mobile_phone = ? or user_name= ?", mobilePhone, userName).First(&user).Error
	return
}

func Insert(user *userInfoProto.UserInf) (err error) {
	user.UserId=uuid.GetUuid()
	DB := db.MasterEngine()
	err = DB.Create(user).Error
	return
}



