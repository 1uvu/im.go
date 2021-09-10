package connect

import (
	"net/http"

	"github.com/gorilla/websocket"

	"im/internal/pkg/logger"
)

func (conn *Connect) serverWS(server *Server, rw http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  server.Option.RBufferSize,
		WriteBufferSize: server.Option.WBufferSize,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	connWS, err := upgrader.Upgrade(rw, r, nil)

	if err != nil {
		logger.Errorf("websocket server got error: %s", err.Error())
		return
	}

	d := NewDialog(0, server.Option.BroadcastSize)

	d.connWS = connWS

	go server.writePump(d, conn)
	go server.readPump(d, conn)
}
