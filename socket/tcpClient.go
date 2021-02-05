package socket

import (
	"fmt"
	"net"
	"singo/socket/proto"
	"strconv"
)

// TCP Client端
func TcpClient() {
	fmt.Println("进入TcpClient的main方法")

	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`+strconv.Itoa(i)
		//msg := `Hello`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		fmt.Println("发送的数据是: ")
		fmt.Println(data)
		conn.Write(data)
	}

}

