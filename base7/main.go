package main

import (
	"din"
	"log"
	"net/http"
	"time"
)

func onlyForV2() din.HandlerFunc {
	return func(c *din.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := din.New()
	r.Use(din.Logger()) // global midlleware
	r.GET("/", func(c *din.Context) {
		c.HTML(http.StatusOK, "<h1>Hello din</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *din.Context) {
			// expect /hello/dinktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}