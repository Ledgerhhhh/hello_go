package ledger

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func TestMw(ctx context.Context, c *app.RequestContext) {
	// your code...
	c.Next(ctx)
	ledger := c.Request.Header.Get("ledger")
	clientIP := c.RemoteAddr()
	fmt.Println(ledger)
	fmt.Println(clientIP)
	fmt.Println("我是中间件啊啊")
	c.JSON(consts.StatusOK, utils.H{
		"code":    consts.StatusOK,
		"message": "我是中间件啊啊",
	})
}
