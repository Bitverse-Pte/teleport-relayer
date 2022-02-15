package tendermint

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/teleport-network/teleport-relayer/app/interfaces"
	teleportsdk "github.com/teleport-network/teleport-sdk-go/client"

	abci "github.com/tendermint/tendermint/abci/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/light/provider"
	tmprototypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/tx"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/teleport-network/teleport-sdk-go/client"
	xibcbsc "github.com/teleport-network/teleport/x/xibc/clients/light-clients/bsc/types"
	xibceth "github.com/teleport-network/teleport/x/xibc/clients/light-clients/eth/types"
	xibctendermint "github.com/teleport-network/teleport/x/xibc/clients/light-clients/tendermint/types"
	clienttypes "github.com/teleport-network/teleport/x/xibc/core/client/types"
	commitmenttypes "github.com/teleport-network/teleport/x/xibc/core/commitment/types"
	"github.com/teleport-network/teleport/x/xibc/core/host"
	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"
	"github.com/teleport-network/teleport/x/xibc/exported"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/types"
	"github.com/teleport-network/teleport-relayer/app/types/errors"
)

var _ interfaces.IChain = new(Tendermint)

var (
	maxRetryAttempts    = 5
	regexpTooHigh       = regexp.MustCompile(`height \d+ must be less than or equal to`)
	regexpMissingHeight = regexp.MustCompile(`height \d+ is not available`)
	regexpTimedOut      = regexp.MustCompile(`Timeout exceeded`)
)

type Tendermint struct {
	Codec                 *codec.ProtoCodec
	TeleportSDK           *client.TeleportClient
	address               string
	chainName             string
	chainType             string
	updateClientFrequency uint64
	queryFilter           string
	fees                  string
}

func NewTendermintClient(
	chainType string,
	chainName string,
	updateClientFrequency uint64,
	config *config.Tendermint,
) (
	*Tendermint, error,
) {
	cdc := MakeCodec()
	cli, err := client.NewClient(config.GrpcAddr, config.ChainID)
	if err != nil {
		panic(err)
	}
	if err := cli.WithKeyring(keyring.NewInMemory(cli.GetCtx().KeyringOptions...)).ImportMnemonic(config.Key.Name, config.Key.Mnemonic); err != nil {
		panic(err)
	}
	address, err := cli.Key(config.Key.Name)
	if err != nil {
		panic(err)
	}
	return &Tendermint{
		chainType:             chainType,
		chainName:             chainName,
		Codec:                 cdc,
		TeleportSDK:           cli,
		updateClientFrequency: updateClientFrequency,
		address:               address,
		queryFilter:           config.QueryFilter,
		fees:                  strconv.FormatInt(config.Fee.Amount, 10) + config.Fee.Denom,
	}, err
}

func (c *Tendermint) ClientUpdateValidate(revisionHeight, delayHeight, updateHeight uint64) (uint64, error) {
	latestHeight, err := c.GetLatestHeight()
	if err != nil {
		return 0, fmt.Errorf("GetLatestHeight error:%+v", err)
	}
	if revisionHeight >= updateHeight {
		return 0, fmt.Errorf("no need update client")
	}
	updateHeight = revisionHeight + 1
	if updateHeight < latestHeight-delayHeight {
		updateHeight = latestHeight
	}
	return updateHeight, nil
}

func (c *Tendermint) GetPackets(fromBlock, toBlock uint64, destChainType string) (*types.Packets, error) {
	times := toBlock - fromBlock + 1
	pktss := make([][]packettypes.Packet, times)
	ackss := make([][]types.AckPacket, times)
	var l sync.Mutex
	var wg sync.WaitGroup
	var anyErr error
	wg.Add(int(times))
	for i := fromBlock; i <= toBlock; i++ {
		go func(height uint64) {
			defer wg.Done()
			pkt, err := c.getBlockPackets(height, destChainType)
			if err != nil {
				anyErr = err
				return
			}
			l.Lock()
			pktss[height-fromBlock] = pkt.BizPackets
			ackss[height-fromBlock] = pkt.AckPackets
			l.Unlock()
		}(i)
	}
	wg.Wait()
	if anyErr != nil {
		return nil, anyErr
	}
	var packets types.Packets
	for _, pkts := range pktss {
		packets.BizPackets = append(packets.BizPackets, pkts...)
	}
	for _, acks := range ackss {
		packets.AckPackets = append(packets.AckPackets, acks...)
	}
	return &packets, nil
}

func (c *Tendermint) getBlockPackets(height uint64, destChainType string) (*types.Packets, error) {
	var bizPackets []packettypes.Packet
	var ackPackets []types.AckPacket
	res, err := c.TeleportSDK.TMServiceQuery.GetBlockByHeight(context.Background(), &tmservice.GetBlockByHeightRequest{
		Height: int64(height),
	})
	if err != nil {
		return nil, err
	}
	packets := types.Packets{}
	for _, t := range res.Block.GetData().Txs {
		hash := hex.EncodeToString(tmhash.Sum(t))
		res, err := c.TeleportSDK.TxClient.GetTx(context.Background(), &tx.GetTxRequest{
			Hash: hash,
		})
		if err != nil {
			continue
		}
		if len(res.TxResponse.Logs) == 0 {
			continue
		}
		stringEvents := res.TxResponse.Logs[0].Events
		tmpPackets, err := c.getPackets(stringEvents, destChainType)
		if err != nil {
			return nil, err
		}
		bizPackets = append(bizPackets, tmpPackets...)

		tmpAckPacks, err := c.getAckPackets(stringEvents, destChainType)
		if err != nil {
			return nil, err
		}
		ackPackets = append(ackPackets, tmpAckPacks...)
	}
	packets.BizPackets = bizPackets
	packets.AckPackets = ackPackets
	return &packets, nil
}

func GetStatus(ack []byte) PacketStatus {
	// TODO
	return Success
}

func (c *Tendermint) GetPacketDataList(port uint64, data []byte) interface{} {
	return nil
}

// TODO
type CrossChainPacket struct {
	Commitment string
	Packet     string
	AckPacket  string
	Height     uint64
	TxHash     string
	Sender     string
	Status     string
}

type PacketStatus int8

const (
	InProcess PacketStatus = iota + 1
	Success
	Fail
)

func (c *Tendermint) GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error) {
	if height != 0 && height <= 2 {
		// TODO return nil, nil, 0, fmt.Errorf("proof queries at height <= 2 are not supported")
	}
	// Use the IAVL height if a valid tendermint height is passed in.
	// A height of 0 will query with the latest state.
	if height != 0 {
		height--
	}
	var key []byte
	switch typ {
	case types.CommitmentPoof:
		key = host.PacketCommitmentKey(sourChainName, destChainName, sequence)
	case types.AckProof:
		key = host.PacketAcknowledgementKey(sourChainName, destChainName, sequence)
	default:
		return nil, errors.ErrGetProof
	}
	// _, proofBz, _, err := c.TeleportSDK.QueryTendermintProof(int64(height), key)

	storeName := host.ModuleName
	path := fmt.Sprintf("/store/%s/%s", storeName, "key")
	res, err := c.TeleportSDK.ABCIQuery.Query(context.Background(), &abci.RequestQuery{
		Data:   key,
		Path:   path,
		Height: int64(height),
		Prove:  true,
	})
	if err != nil {
		return nil, err
	}
	merkleProof, err := commitmenttypes.ConvertProofs(res.ProofOps)
	if err != nil {
		//TODO return nil, nil, 0, err
	}
	proofBz, err := c.Codec.Marshal(&merkleProof)
	if err != nil {
		return nil, err
	}
	return proofBz, nil
}

// ConvertProofs converts crypto.ProofOps into MerkleProof

func (c *Tendermint) RelayPackets(msgs []sdk.Msg) (string, error) {
	var err error
	var packetMsgs []sdk.Msg
	var packetDetail []string
	for _, val := range msgs {
		switch pkt := val.(type) {
		case *packettypes.MsgRecvPacket:
			pkt.Signer = c.address
			packetMsgs = append(packetMsgs, pkt)
			packetMsg := fmt.Sprintf("srcChain:%v,dest%v,sequence:%v chainType:packet", pkt.Packet.SourceChain, pkt.Packet.DestinationChain, pkt.Packet.Sequence)
			packetDetail = append(packetDetail, packetMsg)
		case *packettypes.MsgAcknowledgement:
			pkt.Signer = c.address
			packetMsgs = append(packetMsgs, pkt)
			packetMsg := fmt.Sprintf("srcChain:%v,dest%v,sequence:%v chainType:ack", pkt.Packet.SourceChain, pkt.Packet.DestinationChain, pkt.Packet.Sequence)
			packetDetail = append(packetDetail, packetMsg)
		default:
			return strings.Join(packetDetail,","), fmt.Errorf("invalid packet type")
		}
	}
	if len(packetMsgs) == 0 {
		return strings.Join(packetDetail,","), fmt.Errorf("invalid msgs or empty")
	}
	txf, err := teleportsdk.Prepare(c.TeleportSDK, packetMsgs[0].GetSigners()[0], packetMsgs[0])
	if err != nil {
		return strings.Join(packetDetail,","), err
	}
	txf = txf.WithFees(c.fees)
	res, err := c.TeleportSDK.Broadcast(txf, packetMsgs...)
	if err != nil {
		return strings.Join(packetDetail,","), fmt.Errorf("broadcast tx error:%+v", err)
	}
	if res.TxResponse.Code != 0 {
		return strings.Join(packetDetail,","), fmt.Errorf(res.TxResponse.RawLog)
	}
	return fmt.Sprintf("Relay tx Hash:%v\nPacketDetail:%v\n", res.TxResponse.TxHash,strings.Join(packetDetail,",")), nil
}

func (c *Tendermint) GetBlockTimestamp(height uint64) (uint64, error) {
	res, err := c.TeleportSDK.TMServiceQuery.GetBlockByHeight(
		context.Background(),
		&tmservice.GetBlockByHeightRequest{Height: int64(height)},
	)
	if err != nil {
		return 0, err
	}
	return uint64(res.Block.Header.Time.Unix()), nil
}

func (c *Tendermint) GetBlockHeader(req *types.GetBlockHeaderReq) (exported.Header, error) {
	res, err := c.TeleportSDK.TMServiceQuery.GetBlockByHeight(
		context.Background(),
		&tmservice.GetBlockByHeightRequest{Height: int64(req.LatestHeight)},
	)
	if err != nil {
		return nil, err
	}
	nextRes, err := c.TeleportSDK.TMServiceQuery.GetBlockByHeight(
		context.Background(),
		&tmservice.GetBlockByHeightRequest{Height: int64(req.LatestHeight) + 1},
	)
	if err != nil {
		time.Sleep(5 * time.Second) // TODO
		nextRes, err = c.TeleportSDK.TMServiceQuery.GetBlockByHeight(
			context.Background(),
			&tmservice.GetBlockByHeightRequest{Height: int64(req.LatestHeight) + 1},
		)
		if err != nil {
			return nil, err
		}
	}
	signedHeader := &tmprototypes.SignedHeader{
		Header: &res.Block.Header,        // TODO ToProto
		Commit: nextRes.Block.LastCommit, // TODO ToProto
	}

	validatorSet, err := c.GetValidator(int64(req.LatestHeight))
	if err != nil {
		return nil, err
	}
	validator, err := validatorSet.ToProto()
	if err != nil {
		return nil, err
	}
	validator.TotalVotingPower = validatorSet.TotalVotingPower()
	trustedValidators, err := c.GetValidator(int64(req.TrustedHeight))
	if err != nil {
		return nil, err
	}
	trustedValidator, err := trustedValidators.ToProto()
	if err != nil {
		return nil, err
	}
	trustedValidator.TotalVotingPower = trustedValidators.TotalVotingPower()

	// The trusted fields may be nil. They may be filled before relaying messages to a client.
	// The relayer is responsible for querying client and injecting appropriate trusted fields.
	return &xibctendermint.Header{
		SignedHeader: signedHeader,
		ValidatorSet: validator,
		TrustedHeight: clienttypes.Height{
			RevisionNumber: req.RevisionNumber,
			RevisionHeight: req.TrustedHeight,
		},
		TrustedValidators: trustedValidator,
	}, nil
}

func (c *Tendermint) GetLightClientState(chainName string) (exported.ClientState, error) {
	ctx := context.Background()

	res, err := c.TeleportSDK.XIBCClientQuery.ClientState(
		ctx,
		&clienttypes.QueryClientStateRequest{ChainName: chainName},
	)
	if err != nil {
		return nil, err
	}

	var clientState exported.ClientState
	if err := c.Codec.UnpackAny(res.ClientState, &clientState); err != nil {
		return nil, err
	}

	return clientState, nil
}

func (c *Tendermint) GetLightClientConsensusState(chainName string, height uint64) (exported.ConsensusState, error) {
	//return c.TeleportSDK.GetConsensusState(chainName, height)
	res, err := c.TeleportSDK.XIBCClientQuery.ConsensusState(context.Background(), &clienttypes.QueryConsensusStateRequest{
		ChainName:      chainName,
		RevisionHeight: height, //TODO
	})
	if err != nil {
		return nil, err
	}
	var consensusState exported.ConsensusState
	if err := c.Codec.UnpackAny(res.ConsensusState, &consensusState); err != nil {
		return nil, err
	}
	return consensusState, nil
}

func (c *Tendermint) GetLatestHeight() (uint64, error) {
	block, err := c.TeleportSDK.TMServiceQuery.GetLatestBlock(context.Background(), new(tmservice.GetLatestBlockRequest))
	if err != nil {
		return 0, err
	}
	var height = block.Block.Header.Height
	return uint64(height), err
}

func (c *Tendermint) GetResult(hash string) (uint64, error) {
	res, err := c.TeleportSDK.TxClient.GetTx(context.Background(), &tx.GetTxRequest{
		Hash: hash,
	})
	if err != nil {
		return 0, err
	}
	code := uint64(res.TxResponse.Code)
	return code, nil
}

func (c *Tendermint) GetLightClientDelayHeight(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	if err != nil {
		return 0, err
	}
	//return res.GetDelayBlock(), nil
	return res.GetDelayBlock(), nil
}

func (c *Tendermint) GetLightClientDelayTime(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	if err != nil {
		return 0, err
	}
	return res.GetDelayTime(), nil
}

func (c *Tendermint) UpdateClient(header exported.Header, chainName string) error {
	h := codectypes.UnsafePackAny(header)
	msg := clienttypes.MsgUpdateClient{
		ChainName: chainName,
		Header:    h,
		Signer:    c.address,
	}
	txf, err := teleportsdk.Prepare(c.TeleportSDK, msg.GetSigners()[0], &msg)
	if err != nil {
		return err
	}
	txf = txf.WithFees(c.fees)
	res, err := c.TeleportSDK.Broadcast(txf, &msg)
	if err != nil {
		return err
	}
	if res.TxResponse.Code != 0 {
		return fmt.Errorf(res.TxResponse.RawLog)
	}
	return nil
}

func (c *Tendermint) BatchUpdateClient(headers []exported.Header, chainName string) error {
	var msgs []sdk.Msg
	for _, header := range headers {
		h := codectypes.UnsafePackAny(header)
		msg := &clienttypes.MsgUpdateClient{
			ChainName: chainName,
			Header:    h,
			Signer:    c.address,
		}
		msgs = append(msgs, msg)
	}
	if len(msgs) == 0 {
		return fmt.Errorf("msgs is empty")
	}
	txf, err := teleportsdk.Prepare(c.TeleportSDK, msgs[0].GetSigners()[0], msgs[0])
	if err != nil {
		return err
	}
	txf = txf.WithFees(c.fees)
	res, err := c.TeleportSDK.Broadcast(txf, msgs...)
	if err != nil {
		return err
	}
	if res.TxResponse.Code != 0 {
		return fmt.Errorf(res.TxResponse.RawLog)
	}
	return nil
}

func (c *Tendermint) GetCommitmentsPacket(sourceChainName, destChainName string, sequence uint64) error {
	//_, err := c.TeleportSDK.PacketCommitment(destChainName, sourceChainName, sequence)
	_, err := c.TeleportSDK.XIBCPacketQuery.PacketCommitment(context.Background(), &packettypes.QueryPacketCommitmentRequest{
		SourceChain: sourceChainName,
		DestChain:   destChainName,
		Sequence:    sequence,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Tendermint) GetReceiptPacket(sourChainName, destChianName string, sequence uint64) (bool, error) {
	//result, err := c.TeleportSDK.PacketReceipt(destChianName, sourChainName, sequence)
	result, err := c.TeleportSDK.XIBCPacketQuery.PacketReceipt(context.Background(), &packettypes.QueryPacketReceiptRequest{
		SourceChain: sourChainName,
		DestChain:   destChianName,
		Sequence:    sequence,
	})
	if err != nil {
		return false, err
	}
	return result.Received, nil
}

func (c *Tendermint) ChainName() string {

	return c.chainName
}

func (c *Tendermint) ChainType() string {
	return c.chainType
}

func (c *Tendermint) UpdateClientFrequency() uint64 {
	return c.updateClientFrequency
}

func (c *Tendermint) GetValidator(height int64) (*tmtypes.ValidatorSet, error) {
	const maxPages = 100

	var (
		perPage = 100
		page    = 1
		total   = -1
	)
	var vals []*tmtypes.Validator
OUTER_LOOP:
	for len(vals) != total && page <= maxPages {
		for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
			//res, err := c.TeleportSDK.Validators(ctx, &height, &page, &perPage)
			res, err := c.TeleportSDK.TMServiceQuery.GetValidatorSetByHeight(context.Background(), &tmservice.GetValidatorSetByHeightRequest{
				Height:     height,
				Pagination: &query.PageRequest{}, //TODO
			})

			switch {
			case err == nil:
				// Validate response.
				if len(res.Validators) == 0 {
					return nil, provider.ErrBadLightBlock{
						Reason: fmt.Errorf(
							"validator set is empty (height: %d, page: %d, per_page: %d)",
							height, page, perPage,
						),
					}
				}
				if res.Pagination.Total <= 0 {
					return nil, provider.ErrBadLightBlock{
						Reason: fmt.Errorf(
							"total number of vals is <= 0: %d (height: %d, page: %d, per_page: %d)",
							res.Pagination.Total, height, page, perPage,
						),
					}
				}
				total = int(res.Pagination.Total)
				for _, v := range res.Validators {
					var pubKey tmcrypto.PubKey
					var pk cryptotypes.PubKey
					if err := c.Codec.UnpackAny(v.PubKey, &pk); err != nil {
						return nil, fmt.Errorf("UnpackAny err:%v", err)
					}
					pubKey, err := cryptocodec.ToTmPubKeyInterface(pk)
					if err != nil {
						//TODO
					}

					tmVal := tmtypes.Validator{
						Address:          pubKey.Address(),
						PubKey:           pubKey,
						VotingPower:      v.VotingPower,
						ProposerPriority: v.ProposerPriority,
					}
					vals = append(vals, &tmVal)
				}
				page++
				continue OUTER_LOOP

			case regexpTooHigh.MatchString(err.Error()):
				return nil, fmt.Errorf("height requested is too high")

			case regexpMissingHeight.MatchString(err.Error()):
				return nil, provider.ErrLightBlockNotFound

			// if we have exceeded retry attempts then return no response error
			case attempt == maxRetryAttempts:
				return nil, provider.ErrNoResponse

			case regexpTimedOut.MatchString(err.Error()):
				// we wait and try again with exponential backoff
				time.Sleep(backoffTimeout(uint16(attempt)))
				continue

			// context canceled or connection refused we return the error
			default:
				return nil, err
			}
		}
	}

	validatorSet := tmtypes.NewValidatorSet(vals)

	return validatorSet, nil
}

// exponential backoff (with jitter)
// 0.5s -> 2s -> 4.5s -> 8s -> 12.5 with 1s variation
func backoffTimeout(attempt uint16) time.Duration {
	// nolint:gosec // G404: Use of weak random number generator
	return time.Duration(500*attempt*attempt)*time.Millisecond + time.Duration(rand.Intn(1000))*time.Millisecond
}

func (c *Tendermint) getCrossChainPackets(stringEvents sdk.StringEvents, destChainType string) ([]packettypes.Packet, error) {
	protoEvents := getEventsVals(types.EventTypeSendPacket, stringEvents)
	var packets []packettypes.Packet
	for _, protoEvent := range protoEvents {
		event, ok := protoEvent.(*packettypes.EventSendPacket)
		if !ok {
			// TODO
		}
		// if event.SrcChain == c.ChainName() || (event.DstChain != c.ChainName() && event.RelayChain != c.ChainName()) {
		//	continue // TODO
		// }
		sequence, err := strconv.Atoi(event.GetSequence())
		if err != nil {
			return nil, err
		}
		//data, err := hex.DecodeString(event.GetDataList())
		tmpPack := packettypes.Packet{
			Sequence:         uint64(sequence),
			SourceChain:      event.GetSrcChain(),
			DestinationChain: event.GetDstChain(),
			Ports:            event.GetPorts(),
			RelayChain:       event.RelayChain, // skip validate empty
			DataList:         event.GetDataList(),
		}
		packets = append(packets, tmpPack)
	}
	return packets, nil
}

func (c *Tendermint) getPackets(stringEvents sdk.StringEvents, destChainType string) ([]packettypes.Packet, error) {
	if strings.Contains(c.queryFilter, types.Packet) {
		return nil, nil
	}
	protoEvents := getEventsVals(types.EventTypeSendPacket, stringEvents)
	var packets []packettypes.Packet
	for _, protoEvent := range protoEvents {
		event, ok := protoEvent.(*packettypes.EventSendPacket)
		if !ok {
			return nil, fmt.Errorf("proto parse failed")
		}
		sequence, err := strconv.Atoi(event.GetSequence())
		if err != nil {
			return nil, err
		}
		tmpPack := packettypes.Packet{
			Sequence:         uint64(sequence),
			SourceChain:      event.GetSrcChain(),
			DestinationChain: event.GetDstChain(),
			Ports:            event.GetPorts(),
			RelayChain:       event.RelayChain,
			DataList:         event.GetDataList(),
		}

		packets = append(packets, tmpPack)
	}
	return packets, nil
}

func (c *Tendermint) getAckPackets(stringEvents sdk.StringEvents, destChainType string) ([]types.AckPacket, error) {
	if strings.Contains(c.queryFilter, types.Ack) {
		return nil, nil
	}
	protoEvents := getEventsVals(types.EventTypeWriteAck, stringEvents)
	var ackPackets []types.AckPacket
	for _, protoEvent := range protoEvents {
		event, ok := protoEvent.(*packettypes.EventWriteAck)
		if !ok {
			return nil, fmt.Errorf("proto parse failed")
		}
		if event.GetSrcChain() != c.ChainName() && event.GetRelayChain() != c.ChainName() {
			continue
		}
		sequence, err := strconv.Atoi(event.GetSequence())
		if err != nil {
			return nil, err
		}
		tmpPack := packettypes.Packet{
			Sequence:         uint64(sequence),
			SourceChain:      event.GetSrcChain(),
			DestinationChain: event.GetDstChain(),
			Ports:            event.GetPorts(),
			RelayChain:       event.RelayChain,
			DataList:         event.GetDataList(),
		}
		var ackPacket types.AckPacket
		ackPacket.Packet = tmpPack
		ackPacket.Acknowledgement = event.Ack
		ackPackets = append(ackPackets, ackPacket)
	}
	return ackPackets, nil
}

func getEventsVals(typ string, stringEvents sdk.StringEvents) []proto.Message {
	var events []proto.Message
	for _, e := range stringEvents {
		abciEvent := abci.Event{}
		if e.Type == typ {
			abciEvent.Type = e.Type
			for _, attr := range e.Attributes {
				abciEvent.Attributes = append(abciEvent.Attributes, abci.EventAttribute{
					Key:   []byte(attr.Key),
					Value: []byte(attr.Value),
				})
			}
			protoEvent, err := sdk.ParseTypedEvent(abciEvent)
			if err != nil {
				//TODO
			}
			events = append(events, protoEvent)
		}
	}
	return events
}

func (c *Tendermint) isExistPacket(typ string, stringEvents sdk.StringEvents) bool {
	for _, e := range stringEvents {
		if e.Type == typ {
			for _, attr := range e.Attributes {
				if attr.Key == "sequence" {
					val := e.Attributes
					fmt.Println(val)
					return true
				}
			}
		}
	}
	return false
}

func (c *Tendermint) isExitsFromStringList(sources []string, target string) bool {
	for _, source := range sources {
		if source == target {
			return true
		}
	}
	return false
}

func MakeCodec() *codec.ProtoCodec {
	ir := codectypes.NewInterfaceRegistry()
	clienttypes.RegisterInterfaces(ir)
	govtypes.RegisterInterfaces(ir)
	xibcbsc.RegisterInterfaces(ir)
	xibctendermint.RegisterInterfaces(ir)
	xibceth.RegisterInterfaces(ir)
	packettypes.RegisterInterfaces(ir)
	ir.RegisterInterface("cosmos.v1beta1.Msg", (*sdk.Msg)(nil))
	tx.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)
	return codec.NewProtoCodec(ir)
}
