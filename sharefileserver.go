package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/sharefile", func(c *gin.Context) {
		bucket := c.Query("bucket")
		file := c.Query("file")
		c.JSON(200, gin.H{
			"bucket": bucket,
			"file":   file,
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a service under development, let's call : http://localhost:8080/sharefile?file=tata&bucket=toto")
	})
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
