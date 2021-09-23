package connect

import (
	"net/http"
	"runtime"
	"time"

	"im/internal/pkg/logger"
	"im/internal/pkg/rpc"
	"im/pkg/config"
)

type Connect struct {
}

var DefaultServer *Server

func NewConnect() *Connect {
	return &Connect{}
}

func (connect *Connect) RunWS() {
	conf := config.GetConfig().Connect

	runtime.GOMAXPROCS(conf.Bucket.CPUs)

	if err := connect.runWS(); err != nil {
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

	if err := connect.runWSRPC(); err != nil {
		logger.Panicf("run websocket rpc server got error: %s", err.Error())
	}

	if err := connect.runWS(); err != nil {
		logger.Panicf("run websocket server got error: %s", err.Error())
	}
}

func (connect *Connect) runWSRPC() error {
	err := rpc.RunRPCServer(config.GetConfig().Common.ETCD.ServerPathConnect, config.GetConfig().Connect.WebsocketRPC.RPCAddress, new(ConnectRPCServer))
	return err
}

func (connect *Connect) runWS() error {
	http.HandleFunc("/ws", func(rw http.ResponseWriter, r *http.Request) {
		connect.serverWS(DefaultServer, rw, r)
	})

	err := http.ListenAndServe(config.GetConfig().Connect.Websocket.Bind, nil)

	return err
}
