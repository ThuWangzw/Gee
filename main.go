package main

import (
	"net/http"

	"github.com/ThuWangzw/Gee/gee"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>HomePage</h1>")
	})
	r.Post("/info", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"name": c.PostForm("name"),
		})
	})
	r.Get("/info", func(c *gee.Context) {
		c.String(http.StatusOK, c.Query("name"))
	})
	r.Get("/book/:id/info", func(c *gee.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "this is book %s information", id)
	})
	r.Get("/static/*filepath", func(c *gee.Context) {
		filepath := c.Param("filepath")
		c.JSON(http.StatusOK, gee.H{
			"filepath": filepath,
		})
	})
	r.Run("0.0.0.0:9000")
}
