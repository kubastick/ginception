# Ginception - Catch panics like a pro
This is exception page middleware that look's like `ASP.NET core` one for `Golang Gin-Gionic`

![](https://raw.githubusercontent.com/kubastick/ginception/master/assets/example-page.png?token=AYaz47c76RD61RM2G7qdaAInGYEbWwuOks5cZIFbwA%3D%3D)

Features:  
- Stack trace of panic
- Query URL
- All cookies and their values
- Request headers

### Instalation
`go get github.com/kubastick/ginception`  
This package is following Go Mod (vgo) modules system
### Usage
```
import (
  "github.com/kubastick/ginception"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(ginception.Middleware()) // Attach ginception middleware
  r.GET("/", func(context *gin.Context) {
    panic("test panic")
  })
  r.Run(":7885")
}
```
