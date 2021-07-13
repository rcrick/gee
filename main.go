package main

import (
	"net/http"

	"gee"
)

func main() {
	g := gee.New()
	g.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "URL.Path = %q\n", c.Path)
	})

	g.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})

	})
	g.Run(":9999")
}
