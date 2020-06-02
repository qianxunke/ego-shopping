package elasticsearch

import (
	"context"
	"fmt"
	"github.com/go-log/log"
	"github.com/kataras/golog"
	"github.com/olivere/elastic/v7"
	"sync"
)

var (
	hostUrl      string
	masterEngine *elastic.Client
	lock         sync.Mutex
)

func MasterEngine() (elasticClient *elastic.Client, err error) {
	if masterEngine != nil {
		goto EXiST
	}
	//锁住
	lock.Lock()
	defer lock.Unlock()
	if masterEngine != nil {
		goto EXiST
	}
	err = createEngine()
	return masterEngine, err

EXiST:
	_, code, err := masterEngine.Ping(hostUrl).Do(context.Background())
	if err != nil {
		golog.Errorf("@@@ Elasticsearch 连接异常挂掉!! ERR: %s code: %d", err, code)
		err = createEngine()
	}
	return masterEngine, err
}

func createEngine() (err error) {
	/*	c := config.C()

	 */
	elasticInfo := &ElasticConfigInfo{}
	elasticInfo.Host = "localhost"
	elasticInfo.Index = "ego_shopping"
	elasticInfo.Port = 9200
	elasticInfo.Type = "goods"
	//err = c.App("elasticsearch", elasticInfo)
	if err != nil {
		log.Logf("[Elasticsearch] %s", err)
		return
	}
	client, err := elastic.NewClient(elastic.SetURL(getConnURL(elasticInfo)))
	if err != nil {
		log.Logf("@@@ Elasticsearch 连接异常挂掉!! %s", err)
	}
	masterEngine = client
	return
}

func getConnURL(info *ElasticConfigInfo) (url string) {
	url = fmt.Sprintf("%s:%d",
		info.Host,
		info.Port)
	hostUrl = url
	return
}
