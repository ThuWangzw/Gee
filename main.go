package main

import (
	"fmt"
	"net/http"

	"github.com/ThuWangzw/Gee/gee"
)

func main() {
	gee := gee.New()
	gee.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "(get)hello!")
	})
	gee.Post("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "(post)hello!")
	})
	gee.Run("0.0.0.0:9000")
}
