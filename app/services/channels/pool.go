package channels

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/teleport-network/teleport-relayer/tools"
)

type CrossChainPacketPool interface {
	Write(key uint64, value []sdk.Msg)
	Delete(key uint64)
}

type PacketPool struct {
	BP         map[uint64][]sdk.Msg
	syncHeight uint64
	lock       sync.Mutex
}

func (p *PacketPool) Write(key uint64, value []sdk.Msg) {
	p.lock.Lock()
	p.BP[key] = value
	p.lock.Unlock()
}

func (p *PacketPool) Delete(key uint64) {
	p.lock.Lock()
	delete(p.BP, key)
	p.lock.Unlock()
}

func (p *PacketPool) GetPacket(height uint64) []sdk.Msg {
	// TODO
	return nil
}

type PacketDBPool struct {
	DB *tools.GormDB
}

func (dbPool *PacketDBPool) Write(value interface{}) {
	dbPool.DB.Client.Create(value)
}

func (dbPool *PacketDBPool) GetPacket(height uint64) []sdk.Msg {
	return nil
}
