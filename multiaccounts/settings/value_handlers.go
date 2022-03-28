package settings

import (
	"fmt"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/status-im/status-go/eth-node/types"
	"github.com/status-im/status-go/multiaccounts/errors"
	"github.com/status-im/status-go/params"
	"github.com/status-im/status-go/protocol/protobuf"
	"github.com/status-im/status-go/sqlite"
	"go.uber.org/zap"
)

func StringFromSyncProtobuf(ss *protobuf.SyncSetting) interface{} {
	return ss.GetValueString()
}

func BoolFromSyncProtobuf(ss *protobuf.SyncSetting) interface{} {
	return ss.GetValueBool()
}

func BytesFromSyncProtobuf(ss *protobuf.SyncSetting) interface{} {
	return ss.GetValueBytes()
}

func Int64FromSyncProtobuf(ss *protobuf.SyncSetting) interface{} {
	return ss.GetValueInt64()
}

func BoolHandler(value interface{}) (interface{}, error) {
	_, ok := value.(bool)
	if !ok {
		return value, errors.ErrInvalidConfig
	}

	return value, nil
}

func JSONBlobHandler(value interface{}) (interface{}, error) {
	spew.Dump("---------------------------")
    spew.Dump(value)
    d := &sqlite.JSONBlob{Data: value}
    spew.Dump(d)
	spew.Dump("---------------------------")
	l := zap.L()
	l.Info(spew.Sdump(value))
	l.Info("---------------------------------hi")
	///just trying to get some debug output to log files
	fmt.Println("------------------------hello world")

    return d, nil
}

func AddressHandler(value interface{}) (interface{}, error) {
	str, ok := value.(string)
	if ok {
		value = types.HexToAddress(str)
	} else {
		return value, errors.ErrInvalidConfig
	}
	return value, nil
}

func NodeConfigHandler(value interface{}) (interface{}, error) {
	jsonString, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	nodeConfig := new(params.NodeConfig)
	err = json.Unmarshal(jsonString, nodeConfig)
	if err != nil {
		return nil, err
	}

	return nodeConfig, nil
}
