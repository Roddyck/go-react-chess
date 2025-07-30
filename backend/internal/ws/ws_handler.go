package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Roddyck/go-react-chess/util"
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
		util.RespondWithError(w, http.StatusBadRequest, "error decoding request body", err)
		return
	}

	session := InitSession(params.ID)
	h.hub.Sessions[params.ID] = session

	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error marshalling response", err)
		return
	}

	w.Write(data)
}

func (h *Handler) GetSessions(w http.ResponseWriter, r *http.Request) {
	type responseParams struct {
		IDs []uuid.UUID `json:"ids"`
	}
	sessions := h.hub.Sessions

	var response responseParams
	for _, session := range sessions {
		response.IDs = append(response.IDs, session.ID)
	}

	data, err := json.Marshal(response)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error marshalling response", err)
		return
	}

	w.Write(data)
}

func (h *Handler) JoinSession(w http.ResponseWriter, r *http.Request) {
	roomID := r.PathValue("roomID")
	username := r.URL.Query().Get("username")
	userID := r.URL.Query().Get("userID")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "user id not found in context", nil)
		return
	}

	sessionID, err := uuid.Parse(roomID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "session id not found in context", nil)
		return
	}

	if len(h.hub.Sessions[sessionID].Players) == 2 {
		util.RespondWithError(w, http.StatusBadRequest, "session is full", nil)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading connection", err)
		return
	}

	player := &Player{
		Conn:      conn,
		Message:   make(chan *Message, 10),
		ID:        userUUID,
		SessionID: sessionID,
		Name:      username,
	}

	message := &Message{
		Action:    "player_joined",
		SessionID: player.SessionID,
		Data:      make(map[string]any),
	}

	h.hub.Register <- player
	h.hub.Broadcast <- message

	go player.writeMessage()
	player.readMessage(h.hub)
}
