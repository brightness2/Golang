/*
 * @Author: Brightness
 * @Date: 2021-10-09 08:57:27
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-11 15:18:57
 * @Description:
 */
package controllers

import (
	"beeWeb/controllers/common"
	"fmt"
	"strconv"
)

type ParamStruct struct {
	Name string `form:"user_name"`
	Age  int    `form:"age"`
}

type MainController struct {
	common.BaseController
}

func (c *MainController) Get() {

	//读取配置
	// user, _ := beego.AppConfig.String("db_user")
	id := c.GetString("id") //获取url参数

	c.Data["id"] = id

	c.TplName = "index.html"

}

func (c *MainController) Post() {
	//获取表单数据
	if c.IsAjax() {
		param := ParamStruct{}
		param.Name = c.GetString("user_name")
		param.Age, _ = strconv.Atoi(c.GetString("age"))
		c.Json(param, 0, "success")
	} else {
		param := ParamStruct{}
		c.ParseForm(&param)
		fmt.Println(param)
		c.Ctx.WriteString("提交成功")
	}

}
