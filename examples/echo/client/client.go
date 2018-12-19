package main

import (
	"fmt"
	"net"
	"time"

	"github.com/leesper/holmes"
	"github.com/json7/stw"
	"github.com/json7/stw/examples/echo"
)

func main() {
	stw.Register(echo.Message{}.MessageNumber(), echo.DeserializeMessage, nil)

	c, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		holmes.Fatalln(err)
	}

	onConnect := stw.OnConnectOption(func(conn stw.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onError := stw.OnErrorOption(func(conn stw.WriteCloser) {
		holmes.Infoln("on error")
	})

	onClose := stw.OnCloseOption(func(conn stw.WriteCloser) {
		holmes.Infoln("on close")
	})

	onMessage := stw.OnMessageOption(func(msg stw.Message, conn stw.WriteCloser) {
		echo := msg.(echo.Message)
		fmt.Printf("%s\n", echo.Content)
	})

	conn := stw.NewClientConn(0, c, onConnect, onError, onClose, onMessage)

	echo := echo.Message{
		Content: "hello, world",
	}

	conn.Start()

	for i := 0; i < 10; i++ {
		time.Sleep(60 * time.Millisecond)
		err := conn.Write(echo)
		if err != nil {
			holmes.Errorln(err)
		}
	}
	holmes.Debugln("hello")
	conn.Close()
}
