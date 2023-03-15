package chainlib

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/lavanet/lava/protocol/chainlib/chainproxy/rpcInterfaceMessages"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy/rpcclient"
	"github.com/lavanet/lava/protocol/lavasession"
	"github.com/lavanet/lava/utils"

	pairingtypes "github.com/lavanet/lava/x/pairing/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/lavanet/lava/protocol/common"
	"github.com/lavanet/lava/relayer/metrics"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

type RestChainParser struct {
	spec       spectypes.Spec
	rwLock     sync.RWMutex
	serverApis map[string]spectypes.ServiceApi
	BaseChainParser
}

// NewRestChainParser creates a new instance of RestChainParser
func NewRestChainParser() (chainParser *RestChainParser, err error) {
	return &RestChainParser{}, nil
}

func (apip *RestChainParser) CraftMessage(serviceApi spectypes.ServiceApi, craftData *CraftData) (ChainMessageForSend, error) {
	if craftData != nil {
		// chain fetcher sends the replaced request inside data
		return apip.ParseMsg(string(craftData.Data), nil, craftData.ConnectionType)
	}

	restMessage := rpcInterfaceMessages.RestMessage{
		Msg:  nil,
		Path: serviceApi.GetName(),
	}
	return apip.newChainMessage(&serviceApi, &serviceApi.ApiInterfaces[0], spectypes.NOT_APPLICABLE, restMessage), nil
}

// ParseMsg parses message data into chain message object
func (apip *RestChainParser) ParseMsg(url string, data []byte, connectionType string) (ChainMessage, error) {
	// Guard that the RestChainParser instance exists
	if apip == nil {
		return nil, errors.New("RestChainParser not defined")
	}

	// Check api is supported and save it in nodeMsg
	serviceApi, err := apip.getSupportedApi(url)
	if err != nil {
		return nil, err
	}

	apiInterface := GetApiInterfaceFromServiceApi(serviceApi, connectionType)
	if apiInterface == nil {
		return nil, fmt.Errorf("could not find the interface %s in the service %s", connectionType, serviceApi.Name)
	}

	// Construct restMessage
	restMessage := rpcInterfaceMessages.RestMessage{
		Msg:  data,
		Path: url,
	}
	if connectionType == http.MethodGet {
		// support for optional params, our listener puts them inside Msg data
		restMessage = rpcInterfaceMessages.RestMessage{
			Msg:  nil,
			Path: url + string(data),
		}
	}

	// TODO fix requested block
	nodeMsg := apip.newChainMessage(serviceApi, apiInterface, spectypes.NOT_APPLICABLE, restMessage)
	return nodeMsg, nil
}

func (*RestChainParser) newChainMessage(serviceApi *spectypes.ServiceApi, apiInterface *spectypes.ApiInterface, requestBlock int64, restMessage rpcInterfaceMessages.RestMessage) *parsedMessage {
	nodeMsg := &parsedMessage{
		serviceApi:     serviceApi,
		apiInterface:   apiInterface,
		msg:            restMessage,
		requestedBlock: requestBlock,
	}
	return nodeMsg
}

// getSupportedApi fetches service api from spec by name
func (apip *RestChainParser) getSupportedApi(name string) (*spectypes.ServiceApi, error) {
	// Guard that the RestChainParser instance exists
	if apip == nil {
		return nil, errors.New("RestChainParser not defined")
	}

	// Acquire read lock
	apip.rwLock.RLock()
	defer apip.rwLock.RUnlock()

	// Fetch server api by name
	api, ok := matchSpecApiByName(name, apip.serverApis)

	// Return an error if spec does not exist
	if !ok {
		return nil, errors.New("rest api not supported")
	}

	// Return an error if api is disabled
	if !api.Enabled {
		return nil, errors.New("api is disabled")
	}

	return &api, nil
}

// SetSpec sets the spec for the RestChainParser
func (apip *RestChainParser) SetSpec(spec spectypes.Spec) {
	// Guard that the RestChainParser instance exists
	if apip == nil {
		return
	}

	// Add a read-write lock to ensure thread safety
	apip.rwLock.Lock()
	defer apip.rwLock.Unlock()

	// extract server and tagged apis from spec
	serverApis, taggedApis := getServiceApis(spec, spectypes.APIInterfaceRest)

	// Set the spec field of the RestChainParser object
	apip.spec = spec
	apip.serverApis = serverApis
	apip.BaseChainParser.SetTaggedApis(taggedApis)
}

// DataReliabilityParams returns data reliability params from spec (spec.enabled and spec.dataReliabilityThreshold)
func (apip *RestChainParser) DataReliabilityParams() (enabled bool, dataReliabilityThreshold uint32) {
	// Guard that the RestChainParser instance exists
	if apip == nil {
		return false, 0
	}

	// Acquire read lock
	apip.rwLock.RLock()
	defer apip.rwLock.RUnlock()

	// Return enabled and data reliability threshold from spec
	return apip.spec.Enabled, apip.spec.GetReliabilityThreshold()
}

// ChainBlockStats returns block stats from spec
// (spec.AllowedBlockLagForQosSync, spec.AverageBlockTime, spec.BlockDistanceForFinalizedData)
func (apip *RestChainParser) ChainBlockStats() (allowedBlockLagForQosSync int64, averageBlockTime time.Duration, blockDistanceForFinalizedData uint32, blocksInFinalizationProof uint32) {
	// Guard that the JsonRPCChainParser instance exists
	if apip == nil {
		return 0, 0, 0, 0
	}

	// Acquire read lock
	apip.rwLock.RLock()
	defer apip.rwLock.RUnlock()

	// Convert average block time from int64 -> time.Duration
	averageBlockTime = time.Duration(apip.spec.AverageBlockTime) * time.Millisecond

	// Return values
	return apip.spec.AllowedBlockLagForQosSync, averageBlockTime, apip.spec.BlockDistanceForFinalizedData, apip.spec.BlocksInFinalizationProof
}

type RestChainListener struct {
	endpoint    *lavasession.RPCEndpoint
	relaySender RelaySender
	logger      *common.RPCConsumerLogs
}

// NewRestChainListener creates a new instance of RestChainListener
func NewRestChainListener(ctx context.Context, listenEndpoint *lavasession.RPCEndpoint, relaySender RelaySender, rpcConsumerLogs *common.RPCConsumerLogs) (chainListener *RestChainListener) {
	// Create a new instance of JsonRPCChainListener
	chainListener = &RestChainListener{
		listenEndpoint,
		relaySender,
		rpcConsumerLogs,
	}

	return chainListener
}

// Serve http server for RestChainListener
func (apil *RestChainListener) Serve(ctx context.Context) {
	// Guard that the RestChainListener instance exists
	if apil == nil {
		return
	}

	// Setup HTTP Server
	app := fiber.New(fiber.Config{})

	app.Use(favicon.New())

	chainID := apil.endpoint.ChainID
	apiInterface := apil.endpoint.ApiInterface
	// Catch Post
	app.Post("/:dappId/*", func(c *fiber.Ctx) error {
		apil.logger.LogStartTransaction("rest-http")

		msgSeed := apil.logger.GetMessageSeed()

		path := "/" + c.Params("*")

		// TODO: handle contentType, in case its not application/json currently we set it to application/json in the Send() method
		// contentType := string(c.Context().Request.Header.ContentType())
		dappID := extractDappIDFromFiberContext(c)
		analytics := metrics.NewRelayAnalytics(dappID, chainID, apiInterface)
		utils.LavaFormatInfo("in <<<", &map[string]string{"path": path, "dappID": dappID, "msgSeed": msgSeed})
		requestBody := string(c.Body())
		reply, _, err := apil.relaySender.SendRelay(ctx, path, requestBody, http.MethodPost, dappID, analytics)
		go apil.logger.AddMetricForHttp(analytics, err, c.GetReqHeaders())

		if err != nil {
			// Get unique GUID response
			errMasking := apil.logger.GetUniqueGuidResponseForError(err, msgSeed)

			// Log request and response
			apil.logger.LogRequestAndResponse("http in/out", true, http.MethodPost, path, requestBody, errMasking, msgSeed, err)

			// Set status to internal error
			c.Status(fiber.StatusInternalServerError)

			// Construct json response
			response := convertToJsonError(errMasking)

			// Return error json response
			return c.SendString(response)
		}
		// Log request and response
		apil.logger.LogRequestAndResponse("http in/out", false, http.MethodPost, path, requestBody, string(reply.Data), msgSeed, nil)

		// Return json response
		return c.SendString(string(reply.Data))
	})

	// Catch the others
	app.Use("/:dappId/*", func(c *fiber.Ctx) error {
		apil.logger.LogStartTransaction("rest-http")
		msgSeed := apil.logger.GetMessageSeed()

		query := "?" + string(c.Request().URI().QueryString())
		path := "/" + c.Params("*")
		dappID := extractDappIDFromFiberContext(c)
		utils.LavaFormatInfo("in <<<", &map[string]string{"path": path, "dappID": dappID, "msgSeed": msgSeed})
		analytics := metrics.NewRelayAnalytics(dappID, chainID, apiInterface)

		reply, _, err := apil.relaySender.SendRelay(ctx, path, query, http.MethodGet, dappID, analytics)
		go apil.logger.AddMetricForHttp(analytics, err, c.GetReqHeaders())
		if err != nil {
			// Get unique GUID response
			errMasking := apil.logger.GetUniqueGuidResponseForError(err, msgSeed)

			// Log request and response
			apil.logger.LogRequestAndResponse("http in/out", true, http.MethodGet, path, "", errMasking, msgSeed, err)

			// Set status to internal error
			c.Status(fiber.StatusInternalServerError)

			// Construct json response
			response := convertToJsonError(errMasking)

			// Return error json response
			return c.SendString(response)
		}
		// Log request and response
		apil.logger.LogRequestAndResponse("http in/out", false, http.MethodGet, path, "", string(reply.Data), msgSeed, nil)

		// Return json response
		return c.SendString(string(reply.Data))
	})

	// Go
	err := app.Listen(apil.endpoint.NetworkAddress)
	if err != nil {
		utils.LavaFormatError("app.Listen(listenAddr)", err, nil)
	}
}

type RestChainProxy struct {
	BaseChainProxy
	nodeUrl string
}

func NewRestChainProxy(ctx context.Context, nConns uint, rpcProviderEndpoint *lavasession.RPCProviderEndpoint, averageBlockTime time.Duration) (ChainProxy, error) {
	if len(rpcProviderEndpoint.NodeUrl) == 0 {
		return nil, utils.LavaFormatError("rpcProviderEndpoint.NodeUrl list is empty missing node url", nil, &map[string]string{"chainID": rpcProviderEndpoint.ChainID, "ApiInterface": rpcProviderEndpoint.ApiInterface})
	}
	rcp := &RestChainProxy{
		BaseChainProxy: BaseChainProxy{averageBlockTime: averageBlockTime},
		nodeUrl:        strings.TrimSuffix(rpcProviderEndpoint.NodeUrl[0], "/"),
	}
	return rcp, nil
}

func (rcp *RestChainProxy) SendNodeMsg(ctx context.Context, ch chan interface{}, chainMessage ChainMessageForSend) (relayReply *pairingtypes.RelayReply, subscriptionID string, relayReplyServer *rpcclient.ClientSubscription, err error) {
	if ch != nil {
		return nil, "", nil, utils.LavaFormatError("Subscribe is not allowed on rest", nil, nil)
	}
	httpClient := http.Client{
		Timeout: LocalNodeTimePerCu(chainMessage.GetServiceApi().ComputeUnits),
	}

	rpcInputMessage := chainMessage.GetRPCMessage()
	nodeMessage, ok := rpcInputMessage.(rpcInterfaceMessages.RestMessage)
	if !ok {
		return nil, "", nil, utils.LavaFormatError("invalid message type in rest, failed to cast RPCInput from chainMessage", nil, &map[string]string{"rpcMessage": fmt.Sprintf("%+v", rpcInputMessage)})
	}

	var connectionTypeSlected string = http.MethodGet
	// if ConnectionType is default value or empty we will choose http.MethodGet otherwise choosing the header type provided
	if chainMessage.GetInterface().Type != "" {
		connectionTypeSlected = chainMessage.GetInterface().Type
	}

	msgBuffer := bytes.NewBuffer(nodeMessage.Msg)
	url := rcp.nodeUrl + nodeMessage.Path

	relayTimeout := LocalNodeTimePerCu(chainMessage.GetServiceApi().ComputeUnits)
	// check if this API is hanging (waiting for block confirmation)
	if chainMessage.GetInterface().Category.HangingApi {
		relayTimeout += rcp.averageBlockTime
	}
	connectCtx, cancel := context.WithTimeout(ctx, relayTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(connectCtx, connectionTypeSlected, url, msgBuffer)
	if err != nil {
		return nil, "", nil, err
	}

	// setting the content-type to be application/json instead of Go's defult http.DefaultClient
	if connectionTypeSlected == http.MethodPost || connectionTypeSlected == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, "", nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, "", nil, err
	}

	reply := &pairingtypes.RelayReply{
		Data: body,
	}
	return reply, "", nil, nil
}
