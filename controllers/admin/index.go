package admin

import "gin"

type Index struct {
	Base Base
}

func (_ *Index) Index(c *gin.Context) {
	c.HTML(200, "admin/index.html", map[string]interface{}{
		"title"	:	"first title",
	})
}