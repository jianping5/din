package main

import (
	"din"
	"net/http"
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

	r.GET("/hello/:name", func(c *din.Context) {
		// expect /hello/dinktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *din.Context) {
		c.JSON(http.StatusOK, din.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}