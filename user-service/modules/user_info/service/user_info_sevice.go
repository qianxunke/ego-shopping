package service

import (
	"fmt"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/user/user_info"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
	"log"
	"sync"
)

var (
	s *userInfoService
	m sync.Mutex
)

//service 服务
type userInfoService struct {
}

func GetService() (*userInfoService, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

//初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	DB := db.MasterEngine()
	if DB == nil {
		log.Fatal("数据库初始化出错！")
		return
	}
	if !DB.HasTable(&user_info.UserInf{}) {
		DB.CreateTable(&user_info.UserInf{})
	}
	s = &userInfoService{}
}
