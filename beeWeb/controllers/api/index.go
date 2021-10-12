/*
 * @Author: Brightness
 * @Date: 2021-10-09 16:49:20
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-11 17:11:13
 * @Description:api
 */

package api

import (
	"beeWeb/controllers/common"

	"github.com/beego/beego/v2/core/validation"
)

type IndexController struct {
	common.BaseController
}

type Param struct {
	Name string
	Age  int
}

func (c *IndexController) Get() {
	//数据校验
	p := Param{"", 60}
	valid := validation.Validation{}
	valid.Required(p.Name, "name")
	// valid.MaxSize(u.Name, 15, "nameMax")
	valid.Range(p.Age, 0, 18, "age")
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			c.Json(err, 0, "success")
		}
	} else {
		c.Json("api 接口", 0, "success")

	}
}
