package service

import (
	"fmt"
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
