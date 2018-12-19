package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"github.com/leesper/holmes"
	"github.com/json7/stw"
	"github.com/json7/stw/examples/chat"
)

func main() {
	defer holmes.Start().Stop()

	stw.Register(chat.ChatMessage, chat.DeserializeMessage, nil)

	c, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		holmes.Fatalln(err)
	}

	onConnect := stw.OnConnectOption(func(c stw.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onError := stw.OnErrorOption(func(c stw.WriteCloser) {
		holmes.Infoln("on error")
	})

	onClose := stw.OnCloseOption(func(c stw.WriteCloser) {
		holmes.Infoln("on close")
	})

	onMessage := stw.OnMessageOption(func(msg stw.Message, c stw.WriteCloser) {
		fmt.Print(msg.(chat.Message).Content)
	})

	options := []stw.ServerOption{
		onConnect,
		onError,
		onClose,
		onMessage,
		stw.ReconnectOption(),
	}

	conn := stw.NewClientConn(0, c, options...)
	defer conn.Close()

	conn.Start()
	for {
		reader := bufio.NewReader(os.Stdin)
		talk, _ := reader.ReadString('\n')
		if talk == "bye\n" {
			break
		} else {
			msg := chat.Message{
				Content: talk,
			}
			if err := conn.Write(msg); err != nil {
				holmes.Infoln("error", err)
			}
		}
	}
	fmt.Println("goodbye")
}
