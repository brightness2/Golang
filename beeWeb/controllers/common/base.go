/*
 * @Author: Brightness
 * @Date: 2021-10-09 15:30:11
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-09 16:34:36
 * @Description:基础 controller
 */

package common

import beego "github.com/beego/beego/v2/server/web"

type JsonStruct struct {
	Data interface{}
	Code int
	Msg  string
}

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Json(data interface{}, code int, msg string) {
	res := JsonStruct{data, code, msg}
	c.Data["json"] = res
	c.ServeJSON()
	c.StopRun()
}
