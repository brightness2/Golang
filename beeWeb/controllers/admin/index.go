/*
 * @Author: Brightness
 * @Date: 2021-10-09 16:30:53
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-09 16:43:45
 * @Description:后台 index 控制器
 */

package admin

import (
	"beeWeb/controllers/common"
)

type IndexController struct {
	common.BaseController
}

func (c *IndexController) Get() {

	c.TplName = "admin/index.html"
}
