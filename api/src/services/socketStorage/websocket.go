package socketstorage

import (
	"github.com/gorilla/websocket"
)

// on connect
var onConnect []func(Ws) = []func(Ws){}

func OnConnect(fn func(Ws)) {
	onConnect = append(onConnect, fn)
}

// web socket

type ws struct {
	onMessage map[string][]func(Message)
	onClose   []func()
	send      func(Message)
	close     func()
}

type Ws interface {
	OnMessage(topic string, handler func(Message))
	Send(Message)
	OnClose(func())
	Close()
}

func (ws *ws) OnMessage(topic string, fn func(Message)) {
	if _, found := ws.onMessage[topic]; !found {
		ws.onMessage[topic] = []func(Message){}
	}
	ws.onMessage[topic] = append(ws.onMessage[topic], fn)
}
func (ws *ws) OnClose(fn func())    { ws.onClose = append(ws.onClose, fn) }
func (ws *ws) Send(message Message) { ws.send(message) }
func (ws *ws) Close()               { ws.close() }

func NewWebSocket(conn *websocket.Conn) Ws {
	instance := &ws{
		send: func(msg Message) {
			if msg.Payload == nil {
				return
			}
			if err, ok := msg.Payload.(error); ok {
				msg.Payload = map[string]string{"error": err.Error()}
			}
			conn.WriteMessage(websocket.TextMessage, msg.Encode()) // here error is ignored because socket close in handled in receiving message
		},
		close:     func() { conn.Close() },
		onMessage: map[string][]func(Message){},
		onClose:   []func(){},
	}
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			message, err := DecodeMessage(msg)
			if err != nil {
				continue
			}
			if _, found := instance.onMessage[message.Topic]; found {
				for _, fn := range instance.onMessage[message.Topic] {
					go fn(message)
				}
			}
		}
		conn.Close()
		for _, fn := range instance.onClose {
			go fn()
		}
	}()
	return instance
}
