package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		hub: hub,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) CreateSession(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error decoding request", err)
		return
	}

	session := InitSession(params.ID)
	h.hub.Sessions[params.ID] = session

	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error marshalling session", err)
		return
	}

	w.Write(data)
}

func (h *Handler) JoinSession(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading connection", err)
		return
	}


	roomID := r.PathValue("roomID")
	username := r.URL.Query().Get("username")
	userID := r.URL.Query().Get("userID")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		log.Println("error parsing user id", err)
		return
	}

	sessionID, err := uuid.Parse(roomID)
	if err != nil {
		log.Println("error parsing session id", err)
		return
	}

	player := &Player{
		Conn:      conn,
		Message:   make(chan *Message),
		ID:        userUUID,
		SessionID: sessionID,
		Name:      username,
	}

	message := &Message{
		Action:    "player_joined",
		SessionID: player.SessionID,
		Data: map[string]any{
			"msg": "Player joined the game",
			"game_id": h.hub.Sessions[sessionID].Game.ID,
		},
	}

	h.hub.Register <- player
	h.hub.Broadcast <- message

	go player.writeMessage()
	go player.readMessage(h.hub)
}
