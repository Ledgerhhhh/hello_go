package service

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_ws/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

var WsService = new(model.WsService)

func HandleFunc(w http.ResponseWriter, r *http.Request) {

	// 升级 HTTP 连接为 WebSocket 连接
	WsService.Upgrade = &websocket.Upgrader{
		// WebSocket 连接的读取缓冲区大小
		ReadBufferSize: 1024,
		// WebSocket 连接的写入缓冲区大小
		WriteBufferSize: 1024,
		// 检查 WebSocket 连接请求的来源
		CheckOrigin: func(r *http.Request) bool {
			// 表示接受所有连接请求的来源
			return true
		},
	}
	conn, err := WsService.Upgrade.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	_ = connHandler(conn)
}

func connHandler(conn *websocket.Conn) error {
	ws := &model.WsClient{
		SessionID:           uuid.New().String(),
		Conn:                conn,
		IsClientAuthorized:  false,
		IsAccountAuthorized: false,
		ConnectTime:         time.Now(),
	}
	// 这个方法对于防止长时间无响应的连接非常有用，可以避免程序在等待无响应的连接上浪费时间。
	err := conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(120)))
	if err != nil {
		return fmt.Errorf("set read deadline error: %s\n", err)
	}
	defer func(wc *model.WsClient) {
		if err := recover(); err != nil {
			_ = fmt.Errorf("WebSocket connHandler - panic ：%s", err)
		}
		err := wsConnectionCloseHandler(wc)
		if err != nil {
			return
		}
	}(ws)
	// 保存用户的session
	WsService.WsClients.Store(ws.SessionID, ws)
	fmt.Println("Client connected")

	// 发送用户他的id
	err = ws.SendMessage([]byte(ws.SessionID))
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			var err *websocket.CloseError
			if errors.As(err, &err) {
				// Client actively disconnected
				fmt.Println("Client actively disconnected")
				break
			} else {
				fmt.Printf("read message error: %s\n", err)
				continue
			}
		}
		if websocket.TextMessage != messageType {
			_ = fmt.Errorf("read message error: %s\n", err)
			continue
		}
		if websocket.CloseMessage == messageType {
			_ = fmt.Errorf("read message error: %s\n", err)
			break
		}
		var ms model.Message
		// 将json转换成结构体
		err = json.Unmarshal(message, &ms)
		if err != nil {
			_ = fmt.Errorf("read message error: %s\n", err)
		}
		fmt.Println(string(message) + "消息")
		if ms.To != "" {
			// 发送给别人的
			value, ok := WsService.WsClients.Load(ms.To)
			if ok {
				_ = value.(*model.WsClient).SendMessage(message)
			} else {
				_ = fmt.Errorf("receive message error: %s\n", ms)
			}
		} else {
			fmt.Printf("Received message: %s\n", ms)
		}
		// 增加客户端连接时间
		_ = conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(120)))
	}
	return err
}

func wsConnectionCloseHandler(wc *model.WsClient) error {
	err := wc.Close()
	if err != nil {
		return fmt.Errorf("close client error: %s\n", err)
	}
	return err
}

func ListenService() {
	http.HandleFunc("/", HandleFunc)
	log.Fatal(http.ListenAndServe(myconfig.GConfig.WsConfig.IP+":"+strconv.Itoa(myconfig.GConfig.WsConfig.Port), nil))
}
