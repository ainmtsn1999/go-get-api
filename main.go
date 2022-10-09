package main

import (
	"github.com/ainmtsn1999/go-get-api/controllers/timecontroller"
	"github.com/ainmtsn1999/go-get-api/controllers/weathercontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/weather", weathercontroller.Index)
	r.GET("/api/time", timecontroller.Index)
	r.Run()

}
