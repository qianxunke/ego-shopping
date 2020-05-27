package db

import (
	"github.com/go-log/log"
	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
	"sync"
)

var (
	masterEngine *gorm.DB //主数据库
	slaveEngine  *gorm.DB //从数据库
	lock         sync.Mutex
)

func init() {
	//basic.Register
}

//配置数据库主库
func MasterEngine() *gorm.DB {

	if masterEngine != nil {
		goto EXiST
	}
	//锁住
	lock.Lock()
	defer lock.Unlock()
	if masterEngine != nil {
		goto EXiST
	}
	createEngine(true)
	return masterEngine

EXiST:
	var err = masterEngine.DB().Ping()
	if err != nil {
		golog.Errorf("@@@ 数据库 master 节点连接异常挂掉!! %s", err)
		createEngine(true)
	}
	return masterEngine
}

// 从库，单例
func SlaveEngine() *gorm.DB {
	if slaveEngine != nil {
		goto EXIST
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		goto EXIST
	}

	createEngine(false)
	return slaveEngine

EXIST:
	var err = slaveEngine.DB().Ping()
	if err != nil {
		golog.Errorf("@@@ 数据库 slave 节点连接异常挂掉!! %s", err)
		createEngine(false)
	}
	return slaveEngine
}

func createEngine(isMaster bool) {
	cfg := &db{}
	cfg.Mysql.URL="root:root@(127.0.0.1:3306)/qianxunke_user?charset=utf8&parseTime=true&loc=Asia%2FShanghai"
	cfg.Mysql.Debug = true
	cfg.Mysql.Enable = true
	cfg.Mysql.MaxIdleConnection = 10
	cfg.Mysql.MaxOpenConnection = 10
	if !cfg.Mysql.Enable {
		log.Logf("[initMysql] 未启用Mysql")
		return
	}
	engine, err := gorm.Open("mysql", cfg.Mysql.URL)
	if err != nil {
		golog.Fatalf("@@@ 初始化数据库连接失败!! %s", err)
		return
	}
	//是否启用日志记录器，将会在控制台打印sql
	engine.LogMode(cfg.Mysql.Debug)
	if cfg.Mysql.MaxIdleConnection > 0 {
		engine.DB().SetMaxIdleConns(cfg.Mysql.MaxIdleConnection)
	}
	if cfg.Mysql.MaxOpenConnection > 0 {
		engine.DB().SetMaxOpenConns(cfg.Mysql.MaxOpenConnection)
	}
	if isMaster {
		masterEngine = engine
	} else {
		slaveEngine = engine
	}
}
