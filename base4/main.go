package main

import (
	"net/http"
	"din"
)

func main() {
	r := din.New()
	r.GET("/", func(c *din.Context) {
		c.HTML(http.StatusOK, "<h1>Hello din</h1>")
	})
	r.GET("/hello", func(c *din.Context) {
		// expect /hello?name=dinktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *din.Context) {
		c.JSON(http.StatusOK, din.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}