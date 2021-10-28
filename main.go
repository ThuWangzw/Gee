package main

import (
	"net/http"
	"time"

	"github.com/ThuWangzw/Gee/gee"
	log "github.com/sirupsen/logrus"
)

func LogMiddleWare() func(*gee.Context) {
	return func(c *gee.Context) {
		startTime := time.Now()
		path := c.Path
		c.Next()
		log.WithFields(log.Fields{
			"URL":    path,
			"Time":   time.Since(startTime),
			"Method": c.Method,
		}).Print("")
	}
}

func main() {
	r := gee.New()
	apiGroup := r.Group("/api")
	apiGroup.Get("/book/:id/info", func(c *gee.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "this is book %s information", id)
	})

	staticGroup := r.Group("/static")
	staticGroup.Get("/*filepath", func(c *gee.Context) {
		filepath := c.Param("filepath")
		c.JSON(http.StatusOK, gee.H{
			"filepath": filepath,
		})
	})

	r.Get("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>HomePage</h1>")
	})
	r.Use(LogMiddleWare())
	r.Run("0.0.0.0:9000")
}
