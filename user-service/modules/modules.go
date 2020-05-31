package modules

import userInfo "ego-user-service/modules/user_info/service"
import "ego-user-service/modules/access"


//初始化各个模块
func Init() {
	access.Init()
	userInfo.Init()
}
