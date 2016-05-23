package main

import (
  "runtime"
  "log"
  "github.com/leesper/tao"
  "github.com/leesper/tao/examples/echo"
)

func init() {
  log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type EchoServer struct {
  tao.Server
}

func NewEchoServer(addr string) *EchoServer {
  return &EchoServer {
    tao.NewTCPServer(addr),
  }
}

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  tao.MessageMap.Register(echo.EchoMessage{}.MessageNumber(), echo.DeserializeEchoMessage)
  tao.HandlerMap.Register(echo.EchoMessage{}.MessageNumber(), echo.NewEchoMessageHandler)

  echoServer := NewEchoServer(":18342")
  defer echoServer.Close()

  echoServer.SetOnConnectCallback(func(client tao.Connection) bool {
    log.Printf("On connect\n")
    return true
  })

  echoServer.SetOnErrorCallback(func() {
    log.Printf("On error\n")
  })

  echoServer.SetOnCloseCallback(func(client tao.Connection) {
    log.Printf("Closing client\n")
  })

  echoServer.SetOnMessageCallback(func(msg tao.Message, client tao.Connection) {
    log.Printf("Receving message\n")
  })

  echoServer.Start()
}
