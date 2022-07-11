package cmd

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	packettypes "github.com/teleport-network/teleport/x/xibc/core/packet/types"

	"github.com/teleport-network/teleport-relayer/app"
	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/app/types"
)

// NewQueryPacketCommand returns the packet by hash and chain name
func NewQueryPacketCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "packet [chainName] [hash]",
		Short:   "Print the packets",
		Example: fmt.Sprintf("relayer query packet teleport 63D51044E292812599EDDEE90FC54F967D10244295E77454B0BE9741D6645E0B -c config.toml"),
		Args:    cobra.RangeArgs(2, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			chainName := args[0]
			hash := args[1]
			packet, err := QueryPacketByHash(chainName, hash)
			if err != nil {
				return err
			}

			for _, v := range packet.BizPackets {
				cmd.Println("Type : Packet")
				cmd.Println("srcChain:", v.GetSrcChain())
				cmd.Println("destChain:", v.GetDstChain())
				cmd.Println("sequence:", v.GetSequence())

				if len(v.TransferData) != 0 {
					var transferData packettypes.TransferData
					if err = transferData.ABIDecode(v.TransferData); err != nil {
						return err
					}
					cmd.Println("TransferData: ", transferData.String())

					amount := big.NewInt(0)
					amount.SetBytes(transferData.Amount)
					cmd.Println("Amount: ", amount.String())
				}

				if len(v.CallData) != 0 {
					var callData packettypes.CallData
					if err = callData.ABIDecode(v.CallData); err != nil {
						return err
					}
					cmd.Println("CallData: ", callData.String())
				}
				cmd.Println()
			}

			for _, acknowledgement := range packet.AckPackets {
				cmd.Println("Type : ACK")
				v := acknowledgement.Packet
				cmd.Println("srcChain:", v.GetSrcChain())
				cmd.Println("destChain:", v.GetDstChain())
				cmd.Println("sequence:", v.GetSequence())

				a := acknowledgement.Acknowledgement
				var ack packettypes.Acknowledgement
				if err := ack.ABIDecode(a); err != nil {
					return err
				}
				cmd.Println("code:", ack.GetCode())
				cmd.Println("message:", ack.GetMessage())
				cmd.Println("result:", ack.GetResult())
				cmd.Println("relayer:", ack.GetRelayer())
				cmd.Println("feeOption:", ack.GetFeeOption(), "\n")

			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")

	return cmd
}

func QueryPacketByHash(chainName, hash string) (*types.Packets, error) {
	a := app.NewApp()
	return a.QueryPacketsByHash(chainName, hash)
}
