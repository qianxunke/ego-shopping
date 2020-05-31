module inventory-service

go 1.13

replace (
	github.com/qianxunke/ego-shopping/ego-common-protos => ../ego-common-protos
	github.com/qianxunke/ego-shopping/ego-plugins => ../ego-plugins
)

require (
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/kataras/golog v0.0.15 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/qianxunke/ego-shopping/ego-common-protos v0.0.0-00010101000000-000000000000 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/shopspring/decimal v1.2.0
	golang.org/x/tools v0.0.0-20200530233709-52effbd89c51
)
