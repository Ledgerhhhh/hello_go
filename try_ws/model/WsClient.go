package model

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

// WsClient is client information
type WsClient struct {

	// Session number
	SessionID string

	// Client number
	ClientID string

	// App key
	AppKey string

	// Account number
	AccountID string

	// Access token
	AccessToken string

	// House number
	HouseID string

	// Whether or not client authorized
	IsClientAuthorized bool

	// Whether or not account authorized
	IsAccountAuthorized bool

	// Connect time
	ConnectTime time.Time

	// Bind time
	BindTime time.Time

	// Websocket connection
	Conn *websocket.Conn

	// Lock
	mutex sync.Mutex

	// AuthType
	// 0:External client authorization
	// 1:Internal client authorization
	AuthType string
}

// Close is close client connection
func (wc *WsClient) Close() error {
	return wc.Conn.Close()
}

// SendMessage is send message to client
func (wc *WsClient) SendMessage(message []byte) error {
	wc.mutex.Lock()
	err := wc.Conn.WriteMessage(websocket.TextMessage, message)
	wc.mutex.Unlock()
	return err
}
