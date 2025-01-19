package service

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Rhaqim/wtbackend/domain"
	"github.com/Rhaqim/wtbackend/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketService struct {
	connections map[string]map[*websocket.Conn]struct{}
	mutext      sync.Mutex
	upgrader    websocket.Upgrader
	translator  domain.TranslationService
}

func NewWebsocketService(translator domain.TranslationService) *WebsocketService {
	return &WebsocketService{
		connections: make(map[string]map[*websocket.Conn]struct{}),
		mutext:      sync.Mutex{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		translator: translator,
	}
}

func (wsService *WebsocketService) Handle(c *gin.Context) {
	conn, err := wsService.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// If upgrade fails, respond with HTTP error (this is before hijacking)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close() // Close the connection when the function exits

	userID := c.Param("id")
	if userID == "" {
		// Send error message over WebSocket instead of HTTP
		conn.WriteJSON(map[string]string{"error": "User ID is required"})
		return
	}

	// Add connection to the list
	wsService.mutext.Lock()
	if _, ok := wsService.connections[userID]; !ok {
		wsService.connections[userID] = make(map[*websocket.Conn]struct{})
	}
	wsService.connections[userID][conn] = struct{}{}
	wsService.mutext.Unlock()

	var message model.Message

	// Handle messages from the client
	for {
		err := conn.ReadJSON(&message)
		if err != nil {
			// Break the loop on read error (e.g., client disconnect)
			fmt.Println("Error reading message:", err)
			break
		}

		message.ID = model.NewID()
		message.SenderID = userID
		message.Timestamp = message.Timestamp.UTC()

		// Set source and target languages (example values)
		message.SourceLanguage = "en"
		message.TargetLanguage = "fr"

		// Translate the message content
		new_message, err := wsService.translator.Translate(message.Content, message.SourceLanguage, message.TargetLanguage)
		if err != nil {
			// Send error message over WebSocket
			conn.WriteJSON(map[string]string{"error": "Translation failed: " + err.Error()})
			continue
		}

		message.TranslatedContent = new_message

		// Broadcast the translated message
		wsService.Broadcast(message)
	}

	// Remove the connection when it closes
	wsService.mutext.Lock()
	delete(wsService.connections[userID], conn)
	if len(wsService.connections[userID]) == 0 {
		delete(wsService.connections, userID)
	}
	wsService.mutext.Unlock()
}

func (wsService *WebsocketService) Broadcast(message model.Message) {
	wsService.mutext.Lock()
	defer wsService.mutext.Unlock()

	for conn := range wsService.connections[message.ReceiverID] {
		conn.WriteJSON(message)
	}
}

func (wsService *WebsocketService) Close() {
	wsService.mutext.Lock()
	defer wsService.mutext.Unlock()

	for userID, connections := range wsService.connections {
		for conn := range connections {
			conn.Close()
		}
		delete(wsService.connections, userID)
	}
}
