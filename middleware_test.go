package ginception

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go setupServer()
	// Replace with something better
	time.Sleep(1 * time.Second)
	m.Run()
}

func setupServer() {
	r := gin.Default()
	r.Use(Middleware())
	r.GET("/", func(context *gin.Context) {
		panic("test panic")
	})
	r.GET("/ok", func(context *gin.Context) {
		context.String(200, "OK")
	})
	err := r.Run(":7885")
	if err != nil {
		panic(err)
	}
}

func TestMiddleware(t *testing.T) {
	req, err := http.Get("http://localhost:7885")
	if err != nil {
		t.Fatal(err)
	}
	if req.StatusCode != 500 {
		t.Fatal("server send responded with code other then 500:", req.StatusCode)
	}
}

func TestMiddlewareOk(t *testing.T) {
	req, err := http.Get("http://localhost:7885/ok")
	if err != nil {
		t.Fatal(err)
	}
	if req.StatusCode != 200 {
		t.Fatal("server send responded with code other then 200:", req.StatusCode)
	}

	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "OK" {
		t.Log(content)
		t.Fatal("Server returned wrong response")
	}
}
