package middlewares

import (

	"github.com/gin-gonic/gin"

	"net/http"

)

func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {

		method := c.Request.Method

		// 必须，指定允许的域名
		c.Header("Access-Control-Allow-Origin", "*")
		// 可选，指定允许的请求方式
		c.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,OPTIONS,PUT,PATCH")
		// 可选，指定自定义 header 参数，多个用 , 隔开
		c.Header("Access-Control-Allow-Headers", "Token")
		// 可选，指定是否允许携带 cookie
		c.Header("Access-Control-Allow-Credentials", "true")
		// 可选，指定时间内减少发送「预检」请求
		c.Header("Access-Control-Max-Age", "60")

		if method == "OPTIONS" {

			c.AbortWithStatus(http.StatusNoContent)

		}

		c.Next()

	}
}
