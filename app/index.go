package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Description 测试
// @Tags 测试
// @Produce application/json
// @Param _ query validator.Pager false "_"
// @Success 200 {object} static.Response
// @Failure 500 {object} static.Response
// @Router / [GET]
func Index(c *gin.Context) {
	fmt.Fprintf(c.Writer, "index")
}
