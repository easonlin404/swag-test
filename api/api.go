package api

import (
	"github.com/gin-gonic/gin"
	"github.com/easonlin404/swag-test/conf"
	"fmt"
)

// @Summary Add a new pet to the store
// @Description get string by ID
// @Success 200 {array} conf.GetMeActifResult
// @Router /testapi/get-string-by-int/{some_id} [get]
func Test(c *gin.Context) {
	//write your code

	r:=conf.GetMeActifResult{}
	fmt.Print(r)
}

