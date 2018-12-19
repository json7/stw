package chat

import (
	"context"

	"github.com/leesper/holmes"
	"github.com/json7/stw"
)

const (
	// ChatMessage is the message number of chat message.
	ChatMessage int32 = 1
)

// Message defines the chat message.
type Message struct {
	Content string
}

// MessageNumber returns the message number.
func (cm Message) MessageNumber() int32 {
	return ChatMessage
}

// Serialize Serializes Message into bytes.
func (cm Message) Serialize() ([]byte, error) {
	return []byte(cm.Content), nil
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message stw.Message, err error) {
	if data == nil {
		return nil, stw.ErrNilData
	}
	content := string(data)
	msg := Message{
		Content: content,
	}
	return msg, nil
}

// ProcessMessage handles the Message logic.
func ProcessMessage(ctx context.Context, conn stw.WriteCloser) {
	holmes.Infof("ProcessMessage")
	s, ok := stw.ServerFromContext(ctx)
	if ok {
		msg := stw.MessageFromContext(ctx)
		s.Broadcast(msg)
	}
}
