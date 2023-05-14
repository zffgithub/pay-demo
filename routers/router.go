package routers

import (
	"shop/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.HomeAction{}, "*:ShowHome") //初始化首页
	web.Router("/pay", &controllers.PayController{}, "*:Pay")
	web.Router("/pay_requst.html", &controllers.ScanShopController{})
	web.Router("/scan.html", &controllers.ScanShopController{}, "*:ScanRender")
	web.Router("/error.html", &controllers.HomeAction{}, "*:ErrorPage")
}
