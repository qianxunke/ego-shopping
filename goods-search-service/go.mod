module ego-goods-search-service

go 1.14

replace (
	github.com/qianxunke/ego-shopping/ego-common-protos => ../ego-common-protos
	github.com/qianxunke/ego-shopping/ego-plugins => ../ego-plugins
	go.etcd.io/etcd => github.com/qianxunke/etcd v0.0.0-20200520232829-54ba9589114f
)

require (
	github.com/go-log/log v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/olivere/elastic/v7 v7.0.15
	github.com/qianxunke/ego-shopping/ego-common-protos v0.0.0-00010101000000-000000000000
	github.com/qianxunke/ego-shopping/ego-plugins v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.29.1
)
