package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"sync"

	//"singo/serializer"
	"github.com/olahol/melody"
)

type GopherInfo struct {
	ID, X, Y string
}

// FangTest 测试接口
func WsSocket(c *gin.Context) {
	fmt.Println("进入WsSocket方法")

	m := melody.New()

	m.HandleRequest(c.Writer, c.Request)

	/*c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "我是WsSocket测试方法 返回！",
	})*/

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println("进入m.HandleMessage方法")
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

}

func WsSocket2(c *gin.Context) {
	fmt.Println("进入WsSocket2方法")

	mrouter := melody.New()
	gophers := make(map[*melody.Session]*GopherInfo)
	lock := new(sync.Mutex)
	counter := 0

	mrouter.HandleRequest(c.Writer, c.Request)

	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println("进入m.HandleMessage方法1")
		mrouter.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println("进入m.HandleMessage方法2")
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

	mrouter.HandleConnect(func(s *melody.Session) {
		fmt.Println("进入m.HandleConnect")
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
		fmt.Println("进入m.HandleDisconnect")
		lock.Lock()
		mrouter.BroadcastOthers([]byte("dis "+gophers[s].ID), s)
		delete(gophers, s)
		lock.Unlock()
	})




}


