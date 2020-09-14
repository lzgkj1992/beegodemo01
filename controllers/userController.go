package controllers

import (
	"github.com/astaxie/beego"
)

type UserInfo struct {
	UserName string `json:"username"`
	Age string `json:"age"`
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	u := UserInfo {
		UserName: "zhangsan",
		Age: "20",
	}

	c.Data["json"] = u
	c.ServeJSON()
}
