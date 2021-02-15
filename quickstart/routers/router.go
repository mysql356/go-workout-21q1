package routers

import (
	"quickstart/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/get", &controllers.MainController{})
}
