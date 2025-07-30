package ws

import (
	"github.com/Roddyck/go-react-chess/backend/internal/game"
	"github.com/google/uuid"
)

type Session struct {
	ID      uuid.UUID
	Players map[uuid.UUID]*Player
	Game    *game.Game
}

type Hub struct {
	Sessions   map[uuid.UUID]*Session
	Register   chan *Player
	Unregister chan *Player
	Broadcast  chan *Message
}

func InitSession(sessionID uuid.UUID) *Session {
	return &Session{
		ID:      sessionID,
		Players: make(map[uuid.UUID]*Player),
		Game:    game.NewGame(),
	}
}

func NewHub() *Hub {
	return &Hub{
		Sessions:   make(map[uuid.UUID]*Session),
		Register:   make(chan *Player),
		Unregister: make(chan *Player),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case player := <-h.Register:
			if _, ok := h.Sessions[player.SessionID]; ok {
				s := h.Sessions[player.SessionID]

				if _, ok := s.Players[player.ID]; !ok {
					s.Players[player.ID] = player
				}
			}
		case player := <-h.Unregister:
			if _, ok := h.Sessions[player.SessionID]; ok {
				if _, ok := h.Sessions[player.SessionID].Players[player.ID]; ok {
					if len(h.Sessions[player.SessionID].Players) == 0 {
						h.Broadcast <- &Message{
							Action:    "player_left",
							SessionID: player.SessionID,
							Data: map[string]any{
								"msg": "Player left the game",
							},
						}
					}

					delete(h.Sessions[player.SessionID].Players, player.ID)
					close(player.Message)
				}
			}
		case message := <-h.Broadcast:
			if _, ok := h.Sessions[message.SessionID]; ok {
				for _, player := range h.Sessions[message.SessionID].Players {
					player.Message <- message
				}
			}
		}
	}
}
