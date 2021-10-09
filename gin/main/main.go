/*
 * @Author: Brightness
 * @Date: 2021-10-08 16:53:25
 * @LastEditors: Brightness
 * @LastEditTime: 2021-10-08 17:21:37
 * @Description:
 */
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
