package service

import (
	"fmt"
	branProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/brand"
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
	s = &service{}
}
