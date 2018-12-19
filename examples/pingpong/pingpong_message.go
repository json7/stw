package pingpong

import "github.com/json7/stw"

const (
	// PingPontMessage defines message number.
	PingPontMessage int32 = 1
)

// Message defines message format.
type Message struct {
	Info string
}

// MessageNumber returns the message number.
func (pp Message) MessageNumber() int32 {
	return PingPontMessage
}

// Serialize serializes Message into bytes.
func (pp Message) Serialize() ([]byte, error) {
	return []byte(pp.Info), nil
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message stw.Message, err error) {
	if data == nil {
		return nil, stw.ErrNilData
	}
	info := string(data)
	msg := Message{
		Info: info,
	}
	return msg, nil
}

// func ProcessPingPongMessage(ctx stw.Context, conn stw.Connection) {
//   if serverConn, ok := conn.(*stw.ServerConnection); ok {
//     if serverConn.GetOwner() != nil {
//       connections := serverConn.GetOwner().GetAllConnections()
//       for v := range connections.IterValues() {
//         c := v.(stw.Connection)
//         c.Write(ctx.Message())
//       }
//     }
//   }
// }
