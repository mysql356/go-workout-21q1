package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ShowConfigController struct {
	beego.Controller
}

func (c *ShowConfigController) Get() {

	//app, _ := beego.GetConfig("app", "appname", "")
	//fmt.Println(beego.GetConfig("app", "appname", ""))

	app, _ := beego.AppConfig.String("appname")
	devport, _ := beego.AppConfig.String("dev::httpport")
	prodport, _ := beego.AppConfig.String("prod::httpport")
	testport, _ := beego.AppConfig.String("test::httpport")

	c.Data["app"] = app
	c.Data["devport"] = devport
	c.Data["prodport"] = prodport
	c.Data["testport"] = testport

	//fmt.Println(beego.AppConfig.String("appname"))
	c.TplName = "showConfig.tpl"
}
