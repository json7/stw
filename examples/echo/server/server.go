package main

import (
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/leesper/holmes"
	"github.com/json7/stw"
	"github.com/json7/stw/examples/echo"
)

// EchoServer represents the echo server.
type EchoServer struct {
	*stw.Server
}

// NewEchoServer returns an EchoServer.
func NewEchoServer() *EchoServer {
	onConnect := stw.OnConnectOption(func(conn stw.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onClose := stw.OnCloseOption(func(conn stw.WriteCloser) {
		holmes.Infoln("closing client")
	})

	onError := stw.OnErrorOption(func(conn stw.WriteCloser) {
		holmes.Infoln("on error")
	})

	onMessage := stw.OnMessageOption(func(msg stw.Message, conn stw.WriteCloser) {
		holmes.Infoln("receving message")
	})

	return &EchoServer{
		stw.NewServer(onConnect, onClose, onError, onMessage),
	}
}

func main() {
	defer holmes.Start().Stop()

	runtime.GOMAXPROCS(runtime.NumCPU())

	stw.Register(echo.Message{}.MessageNumber(), echo.DeserializeMessage, echo.ProcessMessage)

	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		holmes.Fatalf("listen error %v", err)
	}
	echoServer := NewEchoServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		echoServer.Stop()
	}()

	echoServer.Start(l)
}
