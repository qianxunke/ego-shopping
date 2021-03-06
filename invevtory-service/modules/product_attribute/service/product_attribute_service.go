package service

import (
	"fmt"
	"github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
	"log"
	"sync"
)

type service struct {
}

var (
	s *service
	m sync.RWMutex
)

//获取服务
func GetService() (*service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

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
	if !DB.HasTable(&product_attribute.ProductAttribute{}) {
		DB.CreateTable(&product_attribute.ProductAttribute{})
	}
	s = &service{}
}
