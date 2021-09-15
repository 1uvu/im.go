package config

import "time"

type ConnectConfig struct {
	SessionExpireTime time.Duration       `yaml:"sessionExpireTime"`
	Auth              ConnectAuth         `yaml:"auth"`
	Websocket         ConnectWebsocket    `yaml:"websocket"`
	WebsocketRPC      ConnectWebsocketRPC `yaml:"websocketRPC"`
	TCP               ConnectTCP          `yaml:"tcp"`
	TCPRPC            ConnectTCPRPC       `yaml:"tcpRPC"`
	Bucket            ConnectBucket       `yaml:"bucket"`
	Server            ConnectServer       `yaml:"server"`
}

type ConnectAuth struct {
	CertPath string `yaml:"certPath"`
	KeyPath  string `yaml:"keyPath"`
}

type ConnectWebsocket struct {
	ServerID string `yaml:"serverID"`
	Bind     string `yaml:"bind"`
}

type ConnectWebsocketRPC struct {
	Address string `yaml:"address"`
}

type ConnectTCP struct {
	ServerID         string `yaml:"serverID"`
	Bind             string `yaml:"bind"`
	SendBuffer       int    `yaml:"sendBuffer"`
	ReceiveBuffer    int    `yaml:"receiveBuffer"`
	KeepAlive        bool   `yaml:"keepalive"`
	Reader           int    `yaml:"reader"`
	ReadBuffer       int    `yaml:"readBuffer"`
	ReadBufferSize   int    `yaml:"readBufferSize"`
	Writer           int    `yaml:"writer"`
	WriterBuffer     int    `yaml:"writeBuffer"`
	WriterBufferSize int    `yaml:"writeBufferSize"`
}

type ConnectTCPRPC struct {
	Address string `yaml:"address"`
}

type ConnectBucket struct {
	CPUs      int    `yaml:"cpus"`
	DialogNum uint64 `yaml:"dialogNum"`
	GroupNum  uint64 `yaml:"groupNum"`
	ArgAmount uint64 `yaml:"argAmount"`
	ArgSize   uint64 `yaml:"argSize"`
	SrvProto  int    `yaml:"srvProto"`
}

type ConnectServer struct {
	WriteWait      time.Duration `yaml:"writeWait"`
	PongWait       time.Duration `yaml:"pongWait"`
	PingPeriod     time.Duration `yaml:"pingPeriod"`
	MaxMessageSize uint64        `yaml:"maxMessageSize"`
	RBufferSize    int           `yaml:"rBufferSize"`
	WBufferSize    int           `yaml:"wBufferSize"`
	BroadcastSize  int           `yaml:"BroadcastSize"`
}
