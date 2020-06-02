module ego-inventory-api

go 1.14

replace github.com/qianxunke/ego-shopping/ego-common-protos => ../ego-common-protos

replace github.com/qianxunke/ego-shopping/ego-common-utils => ../ego-common-utils

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gin-gonic/gin v1.6.3
	github.com/qianxunke/ego-shopping/ego-common-protos v0.0.0-00010101000000-000000000000
	github.com/qianxunke/ego-shopping/ego-common-utils v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.29.1
)
