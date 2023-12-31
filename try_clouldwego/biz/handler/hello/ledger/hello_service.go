// Code generated by hertz generator.

package ledger

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	ledger "ledger/biz/model/hello/ledger"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var req ledger.HelloReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(ledger.HelloResp)

	resp.RespBody = "id: " + req.Id +
		" age: " + strconv.FormatInt(int64(*req.Age), 10) +
		" hobbies: " + strings.Join(req.Hobbies, ",")
	host := c.Host()
	fmt.Println(string(host))
	c.NotFound()
	c.JSON(consts.StatusOK, resp)
}
