package models
// struct for messages

import (
)

type CreateRoomReq struct {
  ID      string `json:"idt"`
  Name    string `json:"name"`
}

type ChatMessage struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

func NewSystemMessage(roomID, content string) *ChatMessage {
  return &ChatMessage{
      Content:  content,
      RoomID:   roomID,
      Username: "system",
  }
}


