package main

import (
	"day7-recovery/gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "welcome to gee")
	})
	r.GET("/err", func(c *gee.Context) {
		ret := []string{"123123"}
		c.String(http.StatusOK, "%d", ret[10000])
	})
	r.Run(":9999")
}
