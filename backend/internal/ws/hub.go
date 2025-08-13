package ws

import (
	"github.com/Roddyck/go-react-chess/internal/database"
	"github.com/Roddyck/go-react-chess/internal/game"
	"github.com/google/uuid"
)

type SessionStatus string

const (
	active = "active"
	full   = "full"
	ended  = "ended"
)

type Session struct {
	ID      uuid.UUID
	Players map[uuid.UUID]*Player
	Game    *game.Game
	Status  SessionStatus
	Name    string
}

type Hub struct {
	Sessions   map[uuid.UUID]*Session
	Register   chan *Player
	Unregister chan *Player
	Broadcast  chan *Message
	db         *database.Queries
}

func InitSession(sessionID uuid.UUID, name string) *Session {
	return &Session{
		ID:      sessionID,
		Players: make(map[uuid.UUID]*Player),
		Game:    game.NewGame(),
		Status:  active,
		Name:    name,
	}
}

func NewHub(dbQueries *database.Queries) *Hub {
	return &Hub{
		Sessions:   make(map[uuid.UUID]*Session, 0),
		Register:   make(chan *Player),
		Unregister: make(chan *Player),
		Broadcast:  make(chan *Message),
		db:         dbQueries,
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
					if player.ID != s.Game.Players[game.White] {
						s.Game.Players[game.Black] = player.ID
					}
				}

				if len(s.Players) == 2 {
					s.Status = full
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
								"msg":  "Player left the game",
								"game": h.Sessions[player.SessionID].Game,
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
					message.Data["game"] = h.Sessions[message.SessionID].Game
					player.Message <- message
				}
			}
		}
	}
}
