package controllers

import (
	"github.com/astaxie/beego"
)

type Article struct {
	Title string `xml:"title"`
	Content string `xml:"content"`
}

type GoodsController struct {
	beego.Controller
}

/*
RestfulAPI 设计指南主要是对 API 接口进行了规定，它要求获取数据使用 Get，增加数据使用Post，修改数据用Put，删除数据用delete
*/
func (c *GoodsController) Get() {
	c.Ctx.WriteString("执行查询操作")
}

func (c *GoodsController) DoAdd() {  // post
	c.Ctx.WriteString("执行增加操作")
}

func (c *GoodsController) DoEdit() {  // put
	c.Ctx.WriteString("执行修改操作")
}

func (c *GoodsController) DoDelete() {  // delete
	c.Ctx.WriteString("执行删除操作")
}
