package main

import (
	"context"
	"net"

	"github.com/leesper/holmes"
	"github.com/json7/stw"
	"github.com/json7/stw/examples/pingpong"
)

var (
	rspChan = make(chan string)
)

func main() {
	defer holmes.Start().Stop()

	stw.Register(pingpong.PingPontMessage, pingpong.DeserializeMessage, ProcessPingPongMessage)

	c, err := net.Dial("tcp", "127.0.0.1:12346")
	if err != nil {
		holmes.Fatalln(err)
	}

	conn := stw.NewClientConn(0, c)
	defer conn.Close()

	conn.Start()
	req := pingpong.Message{
		Info: "ping",
	}
	for {
		conn.Write(req)
		holmes.Infoln(<-rspChan)
	}
}

// ProcessPingPongMessage handles business logic.
func ProcessPingPongMessage(ctx context.Context, conn stw.WriteCloser) {
	rsp := stw.MessageFromContext(ctx).(pingpong.Message)
	rspChan <- rsp.Info
}
