package mw

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"log"
	"time"
)

type login struct {
	Username string `form:"username,required" json:"username,required"`
	Password string `form:"password,required" json:"password,required"`
}

var identityKey = "id"

func PingHandler(c context.Context, ctx *app.RequestContext) {
	user, _ := ctx.Get(identityKey)
	ctx.JSON(200, utils.H{
		"message": fmt.Sprintf("username:%v", user.(*User).UserName),
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func MyMiddleware(ctx context.Context, c *app.RequestContext) {
	c.Next(ctx)

	ledger := c.Request.Header.Get("ledger")

	clientIP := c.RemoteAddr()

	fmt.Println(ledger)

	fmt.Println(clientIP)

	c.JSON(consts.StatusOK, utils.H{
		"code":    consts.StatusOK,
		"message": "我是中间件啊啊",
	})
}

func Cors(h *server.Hertz) {
	//h.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"https://foo.com"},
	//	AllowMethods:     []string{"PUT", "PATCH"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "https://github.com"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
}

// LoggingMiddleware  是一个简单的中间件，用于打印请求信息,全局异常捕获
func LoggingMiddleware(ctx context.Context, c *app.RequestContext) {
	// 过滤器链的作用,先放行过去
	c.Next(ctx)
	// 放行回来执行的逻辑
	if len(c.Errors) == 0 {
		// 没有收集到异常直接返回
		fmt.Println("return")
		return
	}
	hertzErr := c.Errors[0]
	// 获取errors包装的err
	err := hertzErr.Unwrap()
	// 打印异常堆栈
	logger.CtxErrorf(ctx, "%+v", err)
	// 获取原始err
	err = errors.Unwrap(err)
	// todo 进行错误代码进行判断
	c.JSON(500, utils.H{
		"code":    consts.StatusOK,
		"message": err.Error(),
	})
}

// Jwt jwt校验
func Jwt(h *server.Hertz) {
	// the jwt mw
	authMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// 登录成功之后的生成jwt的token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		// TODO token校验出现一些问题,我存放在请求头里面但是还是有问题
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		// 用于校验登录的,用户名和密码
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVals login
			if err := c.BindAndValidate(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Hertz",
					FirstName: "CloudWeGo",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		// 执行权限的校验
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		// 未认证的时候返回的数据
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": "请先进行登录",
			})
		},
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	h.POST("/login", authMiddleware.LoginHandler)
	// 配置404页面
	h.NoRoute(authMiddleware.MiddlewareFunc(), func(ctx context.Context, c *app.RequestContext) {
		claims := jwt.ExtractClaims(ctx, c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, map[string]string{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := h.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/ping", PingHandler)
	}

}
