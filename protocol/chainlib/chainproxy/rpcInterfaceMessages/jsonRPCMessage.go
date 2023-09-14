package rpcInterfaceMessages

import (
	"encoding/json"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy/rpcclient"
	"github.com/lavanet/lava/protocol/parser"
	"github.com/lavanet/lava/utils"
)

var ErrFailedToConvertMessage = sdkerrors.New("RPC error", 1000, "failed to convert a message")

type JsonrpcMessage struct {
	Version                string               `json:"jsonrpc,omitempty"`
	ID                     json.RawMessage      `json:"id,omitempty"`
	Method                 string               `json:"method,omitempty"`
	Params                 interface{}          `json:"params,omitempty"`
	Error                  *rpcclient.JsonError `json:"error,omitempty"`
	Result                 json.RawMessage      `json:"result,omitempty"`
	chainproxy.BaseMessage `json:"-"`
}

func ConvertJsonRPCMsg(rpcMsg *rpcclient.JsonrpcMessage) (*JsonrpcMessage, error) {
	// Return an error if the message was not sent
	if rpcMsg == nil {
		return nil, ErrFailedToConvertMessage
	}

	msg := &JsonrpcMessage{
		Version: rpcMsg.Version,
		ID:      rpcMsg.ID,
		Method:  rpcMsg.Method,
		Error:   rpcMsg.Error,
		Result:  rpcMsg.Result,
	}

	if rpcMsg.Params != nil {
		msg.Params = rpcMsg.Params
	}

	return msg, nil
}

func ConvertBatchElement(batchElement rpcclient.BatchElemWithId) (JsonrpcMessage, error) {
	var JsonError *rpcclient.JsonError
	var ok bool
	if batchElement.Error != nil {
		JsonError, ok = batchElement.Error.(*rpcclient.JsonError)
		if !ok {
			return JsonrpcMessage{}, batchElement.Error
		}
	}
	var result json.RawMessage
	if batchElement.Result != nil {
		resultRef, ok := batchElement.Result.(*json.RawMessage)
		if !ok {
			return JsonrpcMessage{}, batchElement.Error
		}
		result = *resultRef
	}
	msg := JsonrpcMessage{
		Version: rpcclient.Vsn,
		ID:      batchElement.ID,
		Error:   JsonError,
		Result:  result,
	}

	return msg, nil
}

func (gm *JsonrpcMessage) UpdateLatestBlockInMessage(latestBlock uint64, modifyContent bool) (success bool) {
	return false
}

func (gm JsonrpcMessage) NewParsableRPCInput(input json.RawMessage) (parser.RPCInput, error) {
	msg := &JsonrpcMessage{}
	err := json.Unmarshal(input, msg)
	if err != nil {
		return nil, utils.LavaFormatError("failed unmarshaling JsonrpcMessage", err, utils.Attribute{Key: "input", Value: input})
	}

	// Make sure the response does not have an error
	if msg.Error != nil && msg.Result == nil {
		return nil, utils.LavaFormatError("response is an error message", msg.Error)
	}
	return ParsableRPCInput{Result: msg.Result}, nil
}

func (cp JsonrpcMessage) GetParams() interface{} {
	return cp.Params
}

func (cp JsonrpcMessage) GetResult() json.RawMessage {
	if cp.Error != nil {
		utils.LavaFormatWarning("GetResult() Request got an error from the node", nil, utils.Attribute{Key: "error", Value: cp.Error})
	}
	return cp.Result
}

func (cp JsonrpcMessage) ParseBlock(inp string) (int64, error) {
	return parser.ParseDefaultBlockParameter(inp)
}

func ParseJsonRPCMsg(data []byte) (msgRet []JsonrpcMessage, err error) {
	// connectionType is currently only used in rest API.
	// Unmarshal request
	var msg JsonrpcMessage
	err = json.Unmarshal(data, &msg)
	if err != nil {
		// we failed unmarshaling
		// try to parse a batch
		var batch []JsonrpcMessage
		errBatch := json.Unmarshal(data, &batch)
		if errBatch != nil {
			// failed parsing both as batch and jsonrpc return the first unmarshal error, unless the first charqacter is "["
			if data[0] == '[' {
				return nil, errBatch
			}
			return nil, err
		}
		return batch, nil
	}
	return []JsonrpcMessage{msg}, nil
}

type JsonrpcBatchMessage struct {
	batch []rpcclient.BatchElemWithId
	chainproxy.BaseMessage
}

func (jbm *JsonrpcBatchMessage) UpdateLatestBlockInMessage(latestBlock uint64, modifyContent bool) (success bool) {
	return false
}

func (jbm *JsonrpcBatchMessage) GetBatch() []rpcclient.BatchElemWithId {
	return jbm.batch
}

func NewBatchMessage(msgs []JsonrpcMessage) (JsonrpcBatchMessage, error) {
	batch := make([]rpcclient.BatchElemWithId, len(msgs))
	for idx, msg := range msgs {
		var elementParams []interface{}
		switch params := msg.Params.(type) {
		case []interface{}:
			elementParams = params
		default:
			return JsonrpcBatchMessage{}, fmt.Errorf("invalid params in batch, batching only supports ordered arguments %s %+v", msg.Method, params)
		}
		element := rpcclient.BatchElemWithId{
			Method: msg.Method,
			Args:   elementParams,
			Result: &json.RawMessage{}, // will unmarshal the result here, if it fails it sets Error field
			Error:  nil,
			ID:     msg.ID,
		}
		batch[idx] = element
	}
	return JsonrpcBatchMessage{batch: batch}, nil
}
