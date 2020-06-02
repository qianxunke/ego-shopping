module github.com/qianxunke/ego-shopping/ego-plugins

go 1.14

replace go.etcd.io/etcd => github.com/qianxunke/etcd v0.0.0-20200520232829-54ba9589114f

require (
	github.com/go-log/log v0.2.0
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/jinzhu/gorm v1.9.12
	github.com/kataras/golog v0.0.15
	github.com/micro/go-micro v1.18.0
	github.com/olivere/elastic/v7 v7.0.15
	github.com/onsi/ginkgo v1.12.2 // indirect
	go.etcd.io/etcd v0.0.0-00010101000000-000000000000
)
