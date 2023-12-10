package main

import (
	"bufio"
	"fmt"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	//testJSON()
	//testXML()
	//testYaml()
	//testParam()
	//testQuery()
	//testBody()
	//test404()
	//testMW()
	//testPath()
	testFile()
}

func testFile() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		open, err := file.Open()
		if err != nil {
			fmt.Println(err)
			return
		}
		all, err := io.ReadAll(open)
		if err != nil {
			fmt.Println(err)
			return
		}
		openFile, err := os.OpenFile("myFile.jpg", os.O_CREATE|os.O_WRONLY, 0655)
		defer openFile.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		writer := bufio.NewWriter(openFile)
		writer.Write(all)
		writer.Flush()
		fmt.Println("File uploaded and saved successfully")
	})
	r.Run(":8000")
}

func testPath() {
	r := gin.Default()
	r.GET("/hhh", func(context *gin.Context) {
		fmt.Println(context.FullPath())       // /hhh
		fmt.Println(context.Request.URL.Path) // /hhh
	})
	r.Run(":8000")
}

func testEngin() {
	//engineNew := gin.New()
	//engineDefault := gin.Default()
	//engineNew.GET("/", func(context *gin.Context) {
	//	context.BindJSON()
	//})
	//engineDefault.GET("/", func(context *gin.Context) {
	//
	//})

}
func mw(c *gin.Context) {
	now := time.Now()
	c.Next()
	since := time.Since(now)
	fmt.Println(since)
}
func testMW() {
	r := gin.Default()
	r.Use(mw)
	r.GET("/", mw, func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(consts.StatusOK, gin.H{
			"message": "hhh",
		})
	})
	r.Run(":8000")
}
func test404() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(consts.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.Run(":8000")
}
func testBody() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		username := c.Query("username")
		address := c.Query("address")
		c.JSON(consts.StatusOK, gin.H{
			"code":     200,
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8000")
}
func testQuery() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		username := c.Query("username")
		address := c.Query("address")
		c.JSON(consts.StatusOK, gin.H{
			"code":     200,
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8000")
}

func testParam() {
	r := gin.Default()
	r.GET("/user/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(consts.StatusOK, gin.H{
			"code":     200,
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8000")
}

func testYaml() {
	r := gin.Default()
	r.GET("/getJson", func(c *gin.Context) {
		c.YAML(consts.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	r.GET("/someJson", func(c *gin.Context) {
		m := msg{
			Name:    "ledger",
			Message: "message",
			Age:     100,
		}
		c.YAML(consts.StatusOK, m)
	})
	r.Run(":8000")
}
func testXML() {
	r := gin.Default()
	r.GET("/getJson", func(c *gin.Context) {
		c.XML(consts.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	r.GET("/someJson", func(c *gin.Context) {
		m := msg{
			Name:    "ledger",
			Message: "message",
			Age:     100,
		}
		c.XML(consts.StatusOK, m)
	})
	r.Run(":8000")
}

type msg struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Age     int    `json:"age"`
}

func testJSON() {
	r := gin.Default()
	r.GET("/getJson", func(c *gin.Context) {
		c.JSON(consts.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	r.GET("/someJson", func(c *gin.Context) {
		m := msg{
			Name:    "ledger",
			Message: "message",
			Age:     100,
		}
		c.JSON(consts.StatusOK, m)
	})
	r.Run(":8000")
}
func testRestful() {
	// 1.创建路由
	r := gin.Default()
	r.GET("/book", func(context *gin.Context) {
		context.JSON(consts.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(consts.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(consts.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(consts.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	r.Run(":8000")
}
