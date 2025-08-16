package ws

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Roddyck/go-react-chess/internal/database"
	"github.com/Roddyck/go-react-chess/internal/game"
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
		_, message, err := p.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := Message{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("error unmarshalling message: ", err)
			return
		}

		handleMessage(msg, hub)
	}
}

func handleMessage(msg Message, hub *Hub) {
	switch msg.Action {
	case "move":
		session := hub.Sessions[msg.SessionID]
		move := &game.Move{
			From: &game.Position{},
			To:   &game.Position{},
		}
		parseMove(msg, move)
		err := session.Game.HandleMove(move)
		if err != nil {
			message := &Message{
				Action:    "error",
				SessionID: msg.SessionID,
				Data: map[string]any{
					"error": err.Error(),
				},
			}
			hub.Broadcast <- message
			return
		}

		message := &Message{
			Action:    "session_update",
			SessionID: msg.SessionID,
			Data: map[string]any{
				"game": session.Game,
			},
		}
		hub.Broadcast <- message

		if session.Game.Status != game.Active {
			handleGameEnd(msg.SessionID, hub)
		}
	case "draw_offer":
		hub.Broadcast <- &Message{
			Action:    "draw_offer",
			SessionID: msg.SessionID,
			Data: map[string]any{
				"user_id": msg.Data["user_id"],
			},
		}
	case "draw_accept":
		session := hub.Sessions[msg.SessionID]
		session.Game.Status = game.Draw
		hub.Broadcast <- &Message{
			Action:    "draw_accept",
			SessionID: msg.SessionID,
			Data: map[string]any{
				"user_id": msg.Data["user_id"],
				"game":    session.Game,
			},
		}
		handleGameEnd(msg.SessionID, hub)
	default:
	}
}

// this is a hacky (just idiotic) way to parse the move data
// i hate it, probably having map[string]any in message data is a bad idea
func parseMove(msg Message, move *game.Move) {
	from := msg.Data["move"].(map[string]any)["from"].(map[string]any)
	to := msg.Data["move"].(map[string]any)["to"].(map[string]any)

	posFrom := &game.Position{
		X: int(from["x"].(float64)),
		Y: int(from["y"].(float64)),
	}
	posTo := &game.Position{
		X: int(to["x"].(float64)),
		Y: int(to["y"].(float64)),
	}
	move.From = posFrom
	move.To = posTo
}

func handleGameEnd(sessionID uuid.UUID, hub *Hub) {
	if session, ok := hub.Sessions[sessionID]; ok {
		session.Status = ended

		board, err := json.Marshal(session.Game.Board)
		if err != nil {
			log.Println("error marshalling board", err)
		}
		history, err := json.Marshal(session.Game.History)
		if err != nil {
			log.Println("error marshalling history", err)
		}
		players, err := json.Marshal(session.Game.Players)
		if err != nil {
			log.Println("error marshalling players", err)
		}

		ctx := context.Background()
		err = hub.db.UpdateGame(ctx, database.UpdateGameParams{
			ID:      session.Game.ID,
			Board:   board,
			Turn:    string(session.Game.Turn),
			History: history,
			Players: players,
		})
		if err != nil {
			log.Println("error updating game", err)
		}

		for _, player := range session.Players {
			hub.Unregister <- player
		}
		delete(hub.Sessions, sessionID)
	}
}
