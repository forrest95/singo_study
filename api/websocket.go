package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"sync"
	"github.com/olahol/melody"
)

type GopherInfo struct {
	ID, X, Y string
}

var mrouter = melody.New()
var gophers = make(map[*melody.Session]*GopherInfo)
var lock = new(sync.Mutex)
var counter = 0

/*func init(){
	fmt.Println("我是init方法11111111111")
}*/

// 此处用不上！
func WebSocket(c *gin.Context) {
	fmt.Println("ws/test websocket 测试")
	mrouter.HandleRequest(c.Writer, c.Request)

	/*c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "我是WsSocket测试方法 返回！",
	})*/
}

func HandleDisconnect(s *melody.Session) {
	fmt.Println("进入 api.HandleDisconnect")
	lock.Lock()
	mrouter.BroadcastOthers([]byte("dis "+gophers[s].ID), s)
	delete(gophers, s)
	lock.Unlock()
}

func HandleConnect(s *melody.Session) {
	fmt.Println("进入 api.HandleConnect")
	lock.Lock()
	for _, info := range gophers {
		s.Write([]byte("set " + info.ID + " " + info.X + " " + info.Y))
	}
	gophers[s] = &GopherInfo{strconv.Itoa(counter), "0", "0"}
	s.Write([]byte("iam " + gophers[s].ID))
	counter += 1
	lock.Unlock()
}

func HandleMessage1(s *melody.Session, msg []byte) {
	fmt.Println("api.HandleMessage1")
	mrouter.BroadcastFilter(msg, func(q *melody.Session) bool {
		return q.Request.URL.Path == s.Request.URL.Path
	})
}

func HandleMessage(s *melody.Session, msg []byte) {
	fmt.Println("进入api.HandleMessage方法2")
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
}









