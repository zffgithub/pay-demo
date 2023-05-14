package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/rs/xid"
)

type HomeAction struct {
	web.Controller
}

/*加载首页及数据*/
func (c *HomeAction) ShowHome() {
	//取值
	siteName, _ := web.AppConfig.String("siteName")
	orderNo := xid.New().String()
	productName := "测试应用-支付功能体验(非商品消费)"

	//数据回显
	c.Data["siteName"] = siteName
	c.Data["pname"] = productName
	c.Data["orderNo"] = orderNo
	c.TplName = "index.html"
}

func (c *HomeAction) ErrorPage() {
	flash := web.ReadFromRequest(&c.Controller)
	e := flash.Data["error"]
	c.Data["error"] = e
	c.TplName = "error.html"
}
