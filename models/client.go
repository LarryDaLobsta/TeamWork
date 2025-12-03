package models


import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"github.com/gofiber/contrib/websocket"
)


type Client struct {
	Conn     *websocket.Conn
	Message  chan *ChatMessage
	ID       string `json:"id"`
	RoomID   string `json:"roomid"`
	Username string `json:"username"`
}

// write the message to the message attibute for client
func (c *Client) writeMessage() {
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		// Build HTML snippet for HTMX to insert
		htmlMsg := fmt.Sprintf(
			`<div id="messages" hx-swap-oob="beforeend">
				<div class="message"><strong>%s:</strong> %s</div>
				</div>`,
			html.EscapeString(message.Username),
			html.EscapeString(message.Content),
		)

		if err := c.Conn.WriteMessage(websocket.TextMessage, []byte(htmlMsg)); err != nil {
			log.Println("write message error:", err)
			return
		}
	}
}

// read the messages from the hub/chat room
func (c *Client) ReadMessage(ChS *ChatRoomServer) {
	defer func() {
		ChS.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, raw, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("read message error: %v", err)
			}
			break
		}

		// handle the message if the web socket connection is still good
		log.Printf("Raw from client: %s", raw)

		// HTMX WebSocket payload format
		var payload struct {
			Text    string                 `json:"text"`
			Headers map[string]interface{} `json:"HEADERS"`
		}

		if err := json.Unmarshal(raw, &payload); err != nil {
			log.Printf("json error: %v", err)
			continue
		}

		log.Printf("Parsed text from %s: %s", c.Username, payload.Text)

		msg := &ChatMessage{
			Content:  payload.Text,
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		// now broadcast the message to the correct Room
		ChS.Broadcast <- msg
	}
}