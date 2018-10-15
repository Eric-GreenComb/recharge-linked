package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/recharge-linked/config"
)

// Cors Cors
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		_clientIP := c.ClientIP()
		if !config.HasInIPs(_clientIP) {
			fmt.Println("!!!!!!!!!! illegal access : ", _clientIP)
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
