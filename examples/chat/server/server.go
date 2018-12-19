package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/leesper/holmes"
	"github.com/json7/stw"
	"github.com/json7/stw/examples/chat"
)

// ChatServer is the chatting server.
type ChatServer struct {
	*stw.Server
}

// NewChatServer returns a ChatServer.
func NewChatServer() *ChatServer {
	onConnectOption := stw.OnConnectOption(func(conn stw.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})
	onErrorOption := stw.OnErrorOption(func(conn stw.WriteCloser) {
		holmes.Infoln("on error")
	})
	onCloseOption := stw.OnCloseOption(func(conn stw.WriteCloser) {
		holmes.Infoln("close chat client")
	})
	return &ChatServer{
		stw.NewServer(onConnectOption, onErrorOption, onCloseOption),
	}
}

func main() {
	defer holmes.Start().Stop()

	stw.Register(chat.ChatMessage, chat.DeserializeMessage, chat.ProcessMessage)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 12345))
	if err != nil {
		holmes.Fatalln("listen error", err)
	}
	chatServer := NewChatServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		chatServer.Stop()
	}()

	holmes.Infoln(chatServer.Start(l))
}
