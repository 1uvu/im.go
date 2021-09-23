package task

import (
	"encoding/json"

	"im/internal/pkg/logger"
	"im/pkg/common"
	"im/pkg/config"
	"im/pkg/proto"
)

type PushArg struct {
	ServerIDx string
	GroupID   int
	UserID    uint64
	Msg       []byte
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

	switch iparam.GetOp() {
	case proto.OpPeerPush:
		param := iparam.(proto.TaskPeerPushParam)
		pushArg := &PushArg{
			UserID:    param.UserID,
			ServerIDx: param.ServerIDx,
			Msg:       param.Msg,
		}

		writeChan(pushArg, pushChan[common.RandInt(config.GetConfig().Task.PushChanCap)])
	case proto.OpGroupPush:
		param := iparam.(proto.TaskGroupPushParam)
		task.groupPush(&param)
	case proto.OpGroupCount:
		param := iparam.(proto.TaskGroupCountParam)
		task.groupCount(&param)
	case proto.OpGroupInfo:
		param := iparam.(proto.TaskGroupInfoParam)
		task.groupInfo(&param)
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

// test 测试上面的能不能通过, 能就删掉下面的

// func (task *Task) Push(paramStr string) {
// 	var iparam proto.ITaskParam
// 	readParam(paramStr, iparam)

// 	switch iparam.GetOp() {
// 	case proto.OpPeerPush:
// 		param := new(proto.TaskPeerPushParam)
// 		readParam(paramStr, param)
// 		pushArg := &PushArg{
// 			UserID:    param.UserID,
// 			ServerIDx: param.ServerIDx,
// 			Msg:       param.Msg,
// 		}

// 		writeChan(pushArg, pushChan[common.RandInt(config.GetConfig().Task.PushChanCap)])
// 	case proto.OpGroupPush:
// 		param := new(proto.TaskGroupPushParam)
// 		readParam(paramStr, param)
// 		task.groupPush(param)
// 	case proto.OpGroupCount:
// 		param := new(proto.TaskGroupCountParam)
// 		readParam(paramStr, param)
// 		task.groupCount(param)
// 	case proto.OpGroupInfo:
// 		param := new(proto.TaskGroupInfoParam)
// 		readParam(paramStr, param)
// 		task.groupInfo(param)
// 	default:
// 		logger.Errorf("task got a unknown op: %s", proto.OPText(iparam.GetOp()))
// 	}
// }

// func readParam(paramStr string, param proto.ITaskParam) {

// 	if err := json.Unmarshal([]byte(paramStr), param); err != nil {
// 		logger.Errorf("task peer push json unmarshal got error: %v", err)
// 	}
// }
