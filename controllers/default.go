package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() { //这是一个方法，这个方法会被go的底层去绑到结构体上
	c.Data["Website"] = "beego.vip"       //意思：我要给网页传输数据，名字叫Website，值是beego.vip
	c.Data["Email"] = "astaxie@gmail.com" //给网页传一个数据，名字叫Email，值是那个邮箱
	c.TplName = "index.tpl"               //用views文件夹中的index.tpl这个网页文件展示给用户，tpl就是网页模板文件，类似HTML
}

func (u *UserController) GetUser() {
	//直接返回JSON给前端
	u.Data["json"] = (map[string]any{
		"code": 200,
		"msg":  "成功",
		"data": map[string]string{
			"username": "lizongqi",
			"nickname": "宗启",
		},
	})

	//关键，必须停止这一句，否则beego会去找页面
	u.ServeJSON()
	//停止执行，不渲染模板
	u.StopRun()
}
