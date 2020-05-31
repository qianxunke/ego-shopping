module user-api

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-log/log v0.1.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/qianxunke/ego-shopping v0.0.0-20200527072247-4480a588825a
	google.golang.org/grpc v1.29.1
)

replace github.com/qianxunke/ego-shopping/ego-common-protos => ../ego-common-protos
