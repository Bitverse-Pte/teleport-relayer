package cmd

import (
	"github.com/teleport-network/teleport-relayer/app"
	"github.com/teleport-network/teleport-relayer/app/types"
)

var (
	ChainName  string
	FromHeight uint64
	ToHeight   uint64

	Hash       string
	SrcChain   string
	DestChain  string
	RelayChain string
	Sequence   uint64
)

func manualRelay() error {
	a := app.NewApp()
	var detail *types.PacketDetail
	if Hash == "" {
		detail = types.NewPacketDetail(ChainName, Sequence, SrcChain, DestChain, FromHeight, ToHeight, "", "")
	} else {
		detail = types.NewPacketDetail(ChainName, 0, "", "", 0, 0, Hash, "")
	}
	err := a.ManualRelay(detail)
	if err != nil {
		return err
	}
	return nil
}
