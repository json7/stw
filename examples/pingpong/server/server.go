package main

import (
	"context"
	"net"
	"runtime"

	"github.com/leesper/holmes"
	"github.com/json7/stw"
	"github.com/json7/stw/examples/pingpong"
)

// PingPongServer defines pingpong server.
type PingPongServer struct {
	*stw.Server
}

// NewPingPongServer returns PingPongServer.
func NewPingPongServer() *PingPongServer {
	onConnect := stw.OnConnectOption(func(conn stw.WriteCloser) bool {
		holmes.Infoln("on connect")
		return true
	})

	onError := stw.OnErrorOption(func(conn stw.WriteCloser) {
		holmes.Infoln("on error")
	})

	onClose := stw.OnCloseOption(func(conn stw.WriteCloser) {
		holmes.Infoln("closing pingpong client")
	})

	return &PingPongServer{
		stw.NewServer(onConnect, onError, onClose),
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer holmes.Start().Stop()
	stw.MonitorOn(12345)
	stw.Register(pingpong.PingPontMessage, pingpong.DeserializeMessage, ProcessPingPongMessage)

	l, err := net.Listen("tcp", ":12346")
	if err != nil {
		holmes.Fatalln("listen error", err)
	}

	server := NewPingPongServer()

	server.Start(l)
}

// ProcessPingPongMessage handles business logic.
func ProcessPingPongMessage(ctx context.Context, conn stw.WriteCloser) {
	ping := stw.MessageFromContext(ctx).(pingpong.Message)
	holmes.Infoln(ping.Info)
	rsp := pingpong.Message{
		Info: "pong",
	}
	conn.Write(rsp)
}
