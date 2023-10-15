package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandlerWs(c *gin.Context) {
	// 升级成websocket协议
	ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalln(err, "err")
	}
	fmt.Println("ws connect success")
	// 完成时关闭连接释放资源
	defer ws.Close()

	go func() {
		// 监听连接“完成”事件，其实也可以说丢失事件
		<-c.Done()
		fmt.Println("ws lost connection")
	}()

	for {
		// 读取客户端发送过来的消息，如果没发就会一直阻塞住
		mt, message, err := ws.ReadMessage()

		if err != nil {
			fmt.Println("read error")
			fmt.Println(err)
			break
		}
		if string(message) == "ping" {
			message = []byte("恭喜您连接成功，您可以发送消息了")
		}
		// 将读取到的消息写回客户端，这里是原封不动的写回去
		fmt.Print("receive message: ", mt, "\n")
		err = ws.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
			break
		}

	}
}
