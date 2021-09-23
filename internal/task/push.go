package task

import (
	"encoding/json"
	"math/rand"

	"im/internal/pkg/logger"
	"im/pkg/config"
	"im/pkg/proto"
)

type PushArg struct {
	ServerID string
	GroupID  int
	UserID   uint64
	Count    uint64
	Msg      []byte
}

var pushChan []chan *PushArg

func init() {
	pushChan = make([]chan *PushArg, config.GetConfig().Task.PushChanCap)
}

func (task *Task) DoPush() {
	for i := range pushChan {
		pushChan[i] = make(chan *PushArg, config.GetConfig().Task.PushParamsCap)

		go func(ch <-chan *PushArg) {
			var arg *PushArg
			for {
				arg = readChan(ch)
				task.peerPush(arg)
			}
		}(pushChan[i])
	}
}

func (task *Task) Push(paramStr string) {
	var iparam proto.ITaskParam

	if err := json.Unmarshal([]byte(paramStr), iparam); err != nil {
		logger.Errorf("task peer push json unmarshal got error: %v", err)
	}

	// todo
	switch iparam.GetOp() {
	case proto.OpPeerPush:
		param := iparam.(proto.TaskPeerPushParam)
		pushArg := &PushArg{
			ServerID: param.ServerID,
		}

		writeChan(pushArg, pushChan[rand.Int()%config.GetConfig().Task.PushChanCap])
	case proto.OpGroupPush:
		param := iparam.(proto.TaskPeerPushParam)
		pushArg := &PushArg{
			ServerID: param.ServerID,
		}
		task.groupPush(pushArg)
	case proto.OpGroupCount:
		param := iparam.(proto.TaskPeerPushParam)
		pushArg := &PushArg{
			ServerID: param.ServerID,
		}
		task.groupPush(pushArg)
	case proto.OpGroupInfo:
		param := iparam.(proto.TaskPeerPushParam)
		pushArg := &PushArg{
			ServerID: param.ServerID,
		}
		task.groupPush(pushArg)
	default:
		logger.Errorf("task got a unknown op: %s", proto.OPText(iparam.GetOp()))
	}
}

func readChan(ch <-chan *PushArg) *PushArg {
	return <-ch
}

func writeChan(arg *PushArg, ch chan<- *PushArg) {
	ch <- arg
}
