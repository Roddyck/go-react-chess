package ws

import (
	//"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	Conn      *websocket.Conn
	Message   chan *Message
	ID        uuid.UUID `json:"id"`
	SessionID uuid.UUID `json:"session_id"`
	Name      string    `json:"name"`
}

type Message struct {
	Action    string         `json:"action"`
	SessionID uuid.UUID      `json:"session_id"`
	Data      map[string]any `json:"data"`
}

func (p *Player) writeMessage() {
	defer func() {
		p.Conn.Close()
	}()

	for {
		message, ok := <-p.Message
		if !ok {
			return
		}

		p.Conn.WriteJSON(message)
	}
}

func (p *Player) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- p
		p.Conn.Close()
	}()

	for {
		_, _, err := p.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("error: %v", err)
			}
		}

		msg := &Message{
			Action:    "player_joined",
			SessionID: p.SessionID,
			Data: map[string]any{},
		}

		hub.Broadcast <- msg
	}
}
