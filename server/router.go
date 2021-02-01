package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"os"
	"singo/api"
	"singo/middleware"
	"strconv"
	"strings"
	"sync"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register",api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired(),middleware.JWTAuthMiddleware())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}


	// fang  自定义api
	v2 := r.Group("/api/v2")
	{
		v2.POST("/auth",api.AuthHandler)  //登录获取token
		v2.GET("/home", middleware.JWTAuthMiddleware(), api.AuthhomeHandler) //测试token

		v2.GET("fang/test", api.FangTest)
		v2.GET("test/byte", api.TestByte)

		//gorm crud test
		v2.GET("gorm/add", api.GormAdd) //添加
		v2.GET("gorm/update", api.GormUpdate) //添加
		v2.GET("gorm/select", api.GormSelect) //查询demo
		v2.GET("gorm/page", api.GormPage) //查询demo 分页

		//view crud test  视图层操作crud
		v2.GET("view/index", api.ViewIndex) //首页展示
		v2.GET("view/page", api.ViewPage) //视图层 分页demo

		//sqlx crud test
		v2.GET("sqlx/add", api.SqlxAdd) //添加
		v2.GET("sqlx/update", api.SqlxUpdate) //添加
		v2.GET("sqlx/select", api.SqlxSelect) //查询demo
		v2.GET("sqlx/page", api.SqlxPage) //查询demo

		v2.GET("user/checkpwd", api.UserCheckPwd) //查询demo

		//websocket测试
		v2.GET("ws/test",api.WsSocket) //websocket
		v2.GET("ws/test2",api.WsSocket2) //websocket
		type GopherInfo struct {
			ID, X, Y string
		}
		mrouter := melody.New()
		gophers := make(map[*melody.Session]*GopherInfo)
		lock := new(sync.Mutex)
		counter := 0

		v2.GET("/ws/test3", func(c *gin.Context) {
			fmt.Println("ws/test3")
			mrouter.HandleRequest(c.Writer, c.Request)
		})
		mrouter.HandleConnect(func(s *melody.Session) {
			fmt.Println("进入m.HandleConnect方法")
			lock.Lock()
			for _, info := range gophers {
				s.Write([]byte("set " + info.ID + " " + info.X + " " + info.Y))
			}
			gophers[s] = &GopherInfo{strconv.Itoa(counter), "0", "0"}
			s.Write([]byte("iam " + gophers[s].ID))
			counter += 1
			lock.Unlock()
		})

		mrouter.HandleDisconnect(func(s *melody.Session) {
			fmt.Println("进入m.HandleDisconnect方法")
			lock.Lock()
			mrouter.BroadcastOthers([]byte("dis "+gophers[s].ID), s)
			delete(gophers, s)
			lock.Unlock()
		})

		mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
			fmt.Println("进入m.HandleMessage方法3")
			fmt.Println(msg)
			p := strings.Split(string(msg), " ")
			lock.Lock()
			info := gophers[s]
			if len(p) == 2 {
				info.X = p[0]
				info.Y = p[1]
				mrouter.BroadcastOthers([]byte("set "+info.ID+" "+info.X+" "+info.Y), s)
			}
			lock.Unlock()
		})

	}
	return r
}
