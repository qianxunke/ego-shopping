module ego-user-service

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-log/log v0.2.0
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.6 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/kataras/golog v0.0.15 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/onsi/ginkgo v1.12.2 // indirect
	github.com/qianxunke/ego-shopping/ego-common-protos v0.0.0-00010101000000-000000000000
	github.com/qianxunke/ego-shopping/ego-plugins v0.0.0-00010101000000-000000000000
	github.com/satori/go.uuid v1.2.0
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/grpc v1.29.1
)

replace (
	github.com/qianxunke/ego-shopping/ego-common-protos => ../ego-common-protos
	github.com/qianxunke/ego-shopping/ego-plugins => ../ego-plugins
	go.etcd.io/etcd => github.com/qianxunke/etcd v0.0.0-20200520232829-54ba9589114f
)
