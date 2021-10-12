/*
 * @Author: Brightness
 * @Date: 2021-10-09 17:49:29
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-09 18:14:49
 * @Description:
 */
package common

import (
	"strings"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	space := strings.Split(c.Ctx.Request.RequestURI, "/")[1]
	if space == "api" {
		c.Json(nil, 404, "接口不存在")
	} else {
		c.TplName = "error/404.html"
	}
}

func (c *ErrorController) Error501() {
	space := strings.Split(c.Ctx.Request.RequestURI, "/")[1]
	if space == "api" {
		c.Json(nil, 501, "接口不存在")
	} else {
		c.TplName = "error/501.html"
	}

}

func (c *ErrorController) ErrorDb() {
	space := strings.Split(c.Ctx.Request.RequestURI, "/")[1]
	if space == "api" {
		c.Json(nil, 1004, "接口不存在")
	} else {
		c.TplName = "error/Db.html"
	}
}
