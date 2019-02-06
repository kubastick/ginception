package ginception

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"
)

type rtException struct {
	Exception      string
	StackTrace     string
	Headers        http.Header
	Cookies        []*http.Cookie
	WithoutCookies bool
	HaveCookies    bool
	Query          string
}

var exceptionPageTemplate *template.Template

func init() {
	var err error
	exceptionPageTemplate, err = template.New("exception page").Parse(exceptionPage)
	if err != nil {
		panic("failed to create template" + err.Error())
	}
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log incident
				log.Println("Ginception caught panic")
				debug.PrintStack()

				// Prepare model
				exception := rtException{}
				exception.StackTrace = string(debug.Stack())
				exception.Exception = template.HTMLEscapeString(r.(string))
				exception.Headers = c.Request.Header
				exception.Query = c.Request.RequestURI
				exception.Cookies = c.Request.Cookies()
				if len(exception.Cookies) < 1 {
					exception.WithoutCookies = true
				}
				exception.HaveCookies = !exception.WithoutCookies

				// Render template
				result := bytes.Buffer{}
				err := exceptionPageTemplate.Execute(&result, exception)

				// Return it
				c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", result.Bytes())
				if err != nil {
					log.Println(err.Error())
				}
			}
		}()
		// Skip to next template
		c.Next()
	}
}
