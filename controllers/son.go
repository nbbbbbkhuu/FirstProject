package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type AddController struct {
	web.Controller // 👈 继承 Beego 父控制器，AddController是我自己的子struct，web.Controller = Beego 官方父 struct
	//之所可以进行重写，是因为web.Controller 自带了以下被重写的方法
}

func (this *AddController) Prepare() { //方法的重写，重写为Prepare

}

func (this *AddController) Get() { //重写GET方法：返回一个HTML页面
	this.Data["content"] = "value"    //给页面传值
	this.Latout = "admin/layout.html" //布局模板
	this.TplName = "admin/add.tpl"    //页面模板
}

func (this *AddController) Post() {
	pkgname := this.GetString("pkgname") //获取前端表单传来的数据
	content := this.GetString("content")
	pk := models.getCruPkg(pkgname) //查分类是否存在，不存在就不创建
	if pk.Id == 0 {
		var pp models.PkgEntity
		pp.Pid = 0
		pp.Pathname = pkgname
		pp.Intro = pkgname
		models.InsertPkg(pp)
		pk = models.GetCruPkg(pkgname)
	}
	//新建文章，存入数据库
	var at models.Article
	at.Pkgid = pk.Id
	at.Content = content
	models.InsertArticle(at)
	//提交完成，跳转到后台首页
	this.Ctx.Redirect(302, "/admin/index")
}
