package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type UserChatServer struct {
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	for {
		// 在这里处理接收到的消息，你可以调用 Subscribe 函数或进行其他逻辑处理
		msg := "Hello, WebSocket Client!"
		tm := time.Now().Format("2006-01-02 15:04:05")
		message := fmt.Sprintf("[ws][%s]:%s", tm, msg)

		err := ws.WriteMessage(1, []byte(message))
		if err != nil {
			log.Println(err)
			break
		}
		time.Sleep(5 * time.Second) // 示例：每隔 5 秒发送一次消息
	}
}

func (UserChatServer) SendMsg(c *gin.Context) {
	fmt.Println("123")
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 启动消息处理器
	go MsgHandler(c, ws)
}

func main() {
	router := gin.Default()

	router.GET("/ws", UserChatServer{}.SendMsg)

	router.Run(":8080")
}
