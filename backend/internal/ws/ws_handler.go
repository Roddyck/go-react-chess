package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Roddyck/go-react-chess/internal/database"
	"github.com/Roddyck/go-react-chess/internal/game"
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

	userID := r.Context().Value("userID").(uuid.UUID)

	session := InitSession(params.ID)
	h.hub.Sessions[params.ID] = session
	session.Game.Players[game.White] = userID

	board, err := json.Marshal(session.Game.Board)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error marshalling board", err)
		return
	}
	history, err := json.Marshal(session.Game.History)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error marshalling history", err)
		return
	}
	players, err := json.Marshal(session.Game.Players)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "error marshalling players", err)
		return
	}

	g, err := h.hub.db.CreateGame(r.Context(), database.CreateGameParams{
		Board:   board,
		Turn:    string(session.Game.Turn),
		History: history,
		Players: players,
	})

	h.hub.Sessions[session.ID].Game.ID = g.ID

	util.RespondWithJSON(w, http.StatusCreated, params)
}

func (h *Handler) GetSessions(w http.ResponseWriter, r *http.Request) {
	type responseParams struct {
		SessionID string `json:"session_id"`
		GameID    string `json:"game_id"`
	}
	sessions := h.hub.Sessions

	var response []responseParams
	for sessionID, session := range sessions {
		response = append(response, responseParams{
			SessionID: sessionID.String(),
			GameID:    session.Game.ID.String(),
		})
	}

	util.RespondWithJSON(w, http.StatusOK, response)
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
		Data: map[string]any{
			"game": h.hub.Sessions[sessionID].Game,
		},
	}

	h.hub.Register <- player
	h.hub.Broadcast <- message

	go player.writeMessage()
	player.readMessage(h.hub)
}
