package ginception

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var exceptionPageTemplate *template.Template

func init() {
	exceptionPageTemplate = template.New(exceptionPage)
	if exceptionPageTemplate == nil {
		panic("failed to create template")
	}
}
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
