package routers

import (
	"quickstart/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/get", &controllers.MainController{})
	beego.Router("/show-config", &controllers.ShowConfigController{})

	//Defalult : Static path - http://localhost:8080/static/s1.jpg
	beego.SetStaticPath("/down1", "download1") //http://localhost:8080/down1/r1.jpg
	beego.SetStaticPath("/down2", "download2") //http://localhost:8080/down2/s5.jpg

}
