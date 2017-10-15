package main

import (
	"github.com/gin-gonic/gin"
	"github.com/easonlin404/swag-test/api"
)


func main() {
	r := gin.New()
	r.GET("/testapi/test/:some_id", api.Test)
	r.Run()

}
