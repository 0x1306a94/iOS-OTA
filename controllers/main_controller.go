package controllers

import (
	"github.com/astaxie/beego"
	"iOS-OTA/common"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "common/common_layout.tpl"
	c.TplName = "index.tpl"
	c.LayoutSections = common.LayoutSections
	css := []string{"/static/css/index.css"}
	c.Data["Css"] = css
	c.Data["Title"] = "iOS OTA 使用说明"
}
