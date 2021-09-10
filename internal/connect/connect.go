package connect

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/google/uuid"

	"im/internal/pkg/logger"
	"im/pkg/config"
)

type Connect struct {
	ServerID string
}

var DefaultServer *Server

func NewConnect() *Connect {
	return &Connect{}
}

func (c *Connect) RunWS() {
	conf := config.GetConfig().Connect

	runtime.GOMAXPROCS(conf.Bucket.CPUs)

	if err := c.runWS(); err != nil {
		logger.Panic(err)
	}

	buckets := make([]*Bucket, conf.Bucket.CPUs)
	for i := range buckets {
		buckets[i] = NewBucket(uint32(i), BucketOption{
			DialogNum: conf.Bucket.DialogNum,
			GroupNum:  conf.Bucket.GroupNum,
			ArgAmount: conf.Bucket.ArgAmount,
			ArgSize:   conf.Bucket.ArgSize,
		})
	}

	operator := new(DefaultOperator)
	DefaultServer = NewServer(buckets, operator, ServerOption{
		WriteWait:      conf.Server.WriteWait * time.Second,
		PongWait:       conf.Server.PongWait * time.Second,
		PingPeriod:     conf.Server.PingPeriod * time.Second,
		MaxMessageSize: conf.Server.MaxMessageSize,
		RBufferSize:    conf.Server.RBufferSize,
		WBufferSize:    conf.Server.WBufferSize,
		BroadcastSize:  conf.Server.BroadcastSize,
	})

	c.ServerID = fmt.Sprintf("%s-%s", "ws", uuid.New().String())

	if err := c.runWSRPC(); err != nil {
		logger.Panicf("run websocket rpc server got error: %s", err.Error())
	}

	if err := c.runWS(); err != nil {
		logger.Panicf("run websocket server got error: %s", err.Error())
	}
}

func (c *Connect) runWSRPC() error {

	return nil
}

func (conn *Connect) runWS() error {
	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		conn.serverWS(DefaultServer, rw, r)
	})

	err := http.ListenAndServe(config.GetConfig().Connect.Websocket.Bind, nil)

	return err
}
