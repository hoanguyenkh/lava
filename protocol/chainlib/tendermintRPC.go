package chainlib

import (
	"context"
	"fmt"
	"time"

	"github.com/lavanet/lava/relayer/lavasession"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

type TendermintChainParser struct{}

func (apip *TendermintChainParser) ParseMsg(url string, data []byte, connectionType string) (ChainMessage, error) {
	return nil, nil
}

func (apip *TendermintChainParser) SetSpec(spec spectypes.Spec) {}

func (apip *TendermintChainParser) DataReliabilityEnabled() bool {
	// TODO
	return false
}

func (apip *TendermintChainParser) ChainBlockStats() (allowedBlockLagForQosSync int64, averageBlockTime time.Duration, blockDistanceForFinalizedData uint32) {
	//TODO
	return 0, 0, 0
}

func NewTendermintRpcChainParser() (chainParser *TendermintChainParser, err error) {
	return nil, fmt.Errorf("not implemented")
}

type TendermintRpcChainListener struct{}

func (apil *TendermintRpcChainListener) Serve() {}

func NewTendermintRpcChainListener(ctx context.Context, listenEndpoint *lavasession.RPCEndpoint, relaySender RelaySender) (apiListener *TendermintRpcChainListener) {
	// open up server for http implementing the api requested (currently implemented in serve_portal in chainproxy, endpoint at listenEndpoint
	// when receiving the data such as url, rpc data, headers (connectionType), use relaySender to wrap verify and send that data
	return nil
}