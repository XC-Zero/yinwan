package logger

import (
	"github.com/gorilla/websocket"
	"golang.org/x/sync/syncmap"
	"net/http"

	"time"
)

var blackList syncmap.Map
var blackListCopy []string
var OperateLogSocketClient *websocket.Conn
var SystemLogSocketClient *websocket.Conn

var SocketClient = websocket.Upgrader{
	HandshakeTimeout:  30 * time.Second,
	ReadBufferSize:    2048,
	WriteBufferSize:   2048,
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       nil,
	EnableCompression: false,
}

func checkOrigin(req *http.Request) bool {
	if _, ok := blackList.Load(req.RemoteAddr); ok {
		return false
	}
	return true
}

func ReadBlackList() []string {
	return blackListCopy
}

func AddBlackList(ip string) {
	var temp []string
	blackList.Store(ip, true)
	blackList.Range(func(key, value interface{}) bool {
		v, _ := value.(string)
		temp = append(temp, v)
		return true
	})
	blackListCopy = temp
}

func DeleteBlackList(ip string) {
	blackList.Delete(ip)
}
