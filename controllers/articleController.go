package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type article struct {
	Title string `form:"title"`
	Content string `form:"content"`
}

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	//page := c.GetString("page")
	//fmt.Println(page)
	//c.Ctx.WriteString("默认get请求" + page）
	c.TplName = "article/add.tpl"
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加文章")
}

func (c *ArticleController) EditArticle() {
	c.Ctx.WriteString("编辑文章")
}

//接收Post请求，执行增加
func (c *ArticleController) DoAddArticle() {
	title := c.GetString("title")
	content := c.GetString("content")
	c.Ctx.WriteString("执行增加" + title + "----" + content)
}

func (c *ArticleController) DoEditArticle() {
	a := article{}
	err := c.ParseForm(&a)
	if err == nil {
		fmt.Println(a)
		c.Ctx.WriteString(a.Title + "----" + a.Content)
	} else {
		c.Ctx.WriteString("执行增加")
	}
}