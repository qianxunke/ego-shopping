module inventory-service

go 1.13

replace (
	github.com/qianxunke/ego-shopping/ego-common-protos => ../ego-common-protos
	github.com/qianxunke/ego-shopping/ego-plugins => ../ego-plugins
	go.etcd.io/etcd => github.com/qianxunke/etcd v0.0.0-20200520232829-54ba9589114f
)

require (
	github.com/go-log/log v0.2.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/micro/go-micro v1.18.0
	github.com/qianxunke/ego-shopping/ego-common-protos v0.0.0-00010101000000-000000000000
	github.com/qianxunke/ego-shopping/ego-plugins v0.0.0-00010101000000-000000000000
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v1.2.0
	golang.org/x/tools v0.0.0-20200530233709-52effbd89c51
	google.golang.org/grpc v1.29.1
)
