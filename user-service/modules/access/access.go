package access

import (
	"fmt"
	"github.com/go-log/log"
	"github.com/qianxunke/ego-shopping/ego-plugins/jwt"
	"sync"
)

//到时候会有人实现你的哈
type service struct {
}

var (
	s  *service
	m  sync.RWMutex
	cfg = &jwt.Jwt{}
)

//接口
type Service interface {
	//生成toke，
	MakeAccessToken(subject *Subject) (ret string, err error);

	//得到缓存的token
	GetCacheAccessToken(subject *Subject) (ret string, err error);

	//清除用户token
	DelUserAccessToken(token string) (err error);

	//解析token获取用户信息
	AuthenticationFromToken(tk string)(subject *Subject,err error)

}

//获取服务
func GetService() (Service, error) {
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
	cfg.SecretKey="qianxunke"

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	s = &service{}
}
