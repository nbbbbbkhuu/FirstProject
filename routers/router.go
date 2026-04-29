package routers

import (
	"quickstart/controllers"

	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	/*MainController是在controllers这个文件夹里面，这句话的意思是调用MainController这个函数。用&的原因是Beego规定路由必须传指针，
	这样使用，beego框架会自动去找里面的get（）*/

	//格式：web.Router(接口地址, 控制器, "请求方式:方法名")
	web.Router("/user, &controllers.UserController{}, "get:GetUser")

}
