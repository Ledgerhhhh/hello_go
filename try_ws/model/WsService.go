package model

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WsService struct {
	IP        string
	Port      int
	Upgrade   *websocket.Upgrader
	WsClients sync.Map
}
