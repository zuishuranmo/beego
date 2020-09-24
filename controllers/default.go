package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"sgwe/models"
)

type MainController struct {
	beego.Controller
}

/*
 *方法的重写:
 *如果一个结构体包含某个方法a,
 */

func (c *MainController) Get() {//匿名字段：一个结构体可以包含另一个结构体，另一个结构体可以直接以类型的方式被包含。外部结构体会自动用有匿名字段的所有属性和方法
	//1.获取请求数据
	userName :=c.Ctx.Input.Query("user")
	password :=c.Ctx.Input.Query("psd")
	//2.使用固定数据进行数据校检
	//admin  123456
	if userName !="admin"|| password != "123456" {
		c.Ctx.ResponseWriter.Write([]byte("数据校检失败"))
	}else {
		c.Ctx.WriteString("数据校检成功")
	}
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "zuishuranmo.com"
	c.TplName = "index.tpl"
}

//编写一个post请求
//func (c *MainController) Post()  {
//	//1.接收post请求的参数
//	name := c.Ctx.Request.FormValue("name")
//	age :=c.Ctx.Request.FormValue("age")
//	//sex :=c.Ctx.Request.FormValue("sex")
//	//2.进行数据校检
//	if name !="admin" && age != "18" {
//		c.Ctx.WriteString("数据校检失败")
//		return
//	}else {
//		c.Ctx.WriteString("数据校检成功")
//	}
//
//}
func (c *MainController) Post() {
	//1.解析前段提交的json格式的数据
	var mine models.Mine
	dataBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil {
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	err = json.Unmarshal(dataBytes,&mine)
	if err != nil {
		c.Ctx.WriteString("数据接收失败2，请重试")
		return
	}
	fmt.Println("姓名：",mine.Name)
	fmt.Println("生日：",mine.Birthday)
	fmt.Println("地址：",mine.Address)
	fmt.Println("别名：",mine.Nick)
	c.Ctx.WriteString("数据解析成功")
}