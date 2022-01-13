package cross_data

import (
	"github.com/teleport-network/teleport-relayer/app/services/channels"
	"github.com/teleport-network/teleport-relayer/app/services/interfaces"
)

type DataService struct {
	Channel channels.Channel
	ChainA  interfaces.IChain
	ChainB  interfaces.IChain
}

func (ds *DataService) SyncCrossChainTxs() {
	// TODO
}
