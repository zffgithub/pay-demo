package controllers

import (
	"shop/models/system"
	"strconv"
	"strings"

	"shop/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type PayController struct {
	web.Controller
}

func (c *PayController) Pay() {
	cName := strings.TrimSpace(c.GetString("c_name"))
	cPhone := strings.TrimSpace(c.GetString("c_phone"))
	amount := strings.TrimSpace(c.GetString("amount"))
	lesson := strings.TrimSpace(c.GetString("lesson"))

	// flash := web.NewFlash()
	// if orderNo == "" {
	// 	flash.Error("订单号为空")
	// 	flash.Store(&c.Controller)
	// 	c.Redirect("/error.html", 302)
	// 	return
	// }
	// amount := strings.TrimSpace(c.GetString("amount"))
	if !c.judgeAmount(amount) {
		// flash.Error("金额有误")
		// flash.Store(&c.Controller)
		c.Redirect("/error.html", 302)
		return
	}

	customeInfo := system.CustomerInfo{
		CName:      cName,
		CPhone:     cPhone,
		SumMoney:   amount,
		Lesson:     lesson,
		Ip:         c.Ctx.Request.RemoteAddr,
		CreateTime: utils.GetBasicDateTime(),
		UpdateTime: utils.GetBasicDateTime(),
	}

	if !system.InsertItem(customeInfo) {
		c.Redirect("/pay_ok.html", 302)
		return
	} else {
		c.Redirect("/pay_fail.html", 302)
		return
	}

}

func (c *PayController) judgeAmount(amount string) bool {
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		logs.Error("输入金额有误", err)
		return false
	}

	return true
}
