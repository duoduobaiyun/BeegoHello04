package controllers

import (
	"BeegoHello04/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//该方法用于post请求
func (c *MainController) Post() {
	//第一种方法
	//userName:=c.Ctx.Input.Query("user")
	//password:=c.Ctx.Input.Query("psd")
	////2、使用固定数据进行数据校验
	////admin 123456
	//if userName!="kun" || password!="wuqing"{
	//	//代表错误处理
	//	c.Ctx.ResponseWriter.Write([]byte("对不起,数据校验错误"))
	//	return
	//}
	////校验正确的情况

	//第二种分方法
	school := c.Ctx.Request.FormValue("school")
	class := c.Ctx.Request.FormValue("class")
	name := c.Ctx.Request.FormValue("name")
	fmt.Println(school, class, name)

	var person models.Person
	//解析前端提交出的json数据格式
	data, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据解析错误,请重新开始操作")
		return
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败,请认真检查代码")
		return
	}
	fmt.Println("大学", person.School)
	fmt.Println("班级", person.Class)
	fmt.Println("名字", person.Name)
}
