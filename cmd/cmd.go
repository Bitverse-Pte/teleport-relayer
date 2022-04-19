package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/cmd/generate"
	"github.com/teleport-network/teleport-relayer/version"
)

var (
	rootCmd = &cobra.Command{
		Use:   "relayer",
		Short: "relayer for xibc",
		Run:   func(cmd *cobra.Command, args []string) { _ = cmd.Help() },
	}
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start XIBC relayer.",
		Run:   func(cmd *cobra.Command, args []string) { online() },
	}
	batchCmd = &cobra.Command{
		Use:   "evmClientSync",
		Short: "eth light client sync",
		Run:   func(cmd *cobra.Command, args []string) { evmClientSync() },
	}
	configInitCmd = &cobra.Command{
		Use:   "init",
		Short: "init configuration file",
		Run:   func(cmd *cobra.Command, args []string) { config.InitConfig() },
	}
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Generate the files needed for create client: clientStatus & consensusState",
		Run:     func(cmd *cobra.Command, args []string) { generate.GenerateClientFiles() },
	}
	genCmd = &cobra.Command{
		Use:     "genFiles [type] [chainName] [chainId] [url] [height] ",
		Aliases: []string{"genF"},
		Short:   "Generate the files needed for create client: clientStatus & consensusState",
		Args:    cobra.RangeArgs(4, 5),
		Run: func(cmd *cobra.Command, args []string) {
			generate.Type = args[0]
			generate.ChainName = args[1]
			generate.ChainID = args[2]
			generate.GrpcAddr = args[3]
			if len(args) == 5 {
				height, err := strconv.ParseInt(args[4], 10, 64)
				if err != nil {
					return
				}
				generate.Height = height

			}
			generate.GenClientFiles()
		},
	}
	versionCmd = version.NewVersionCommand()

	manualRelayCmd = &cobra.Command{
		Use:     "relay [chainName] [fromHeight] [toHeight] [srcChain] [destChain] [sequence] [relayChain]",
		Short:   "manual relay with the packet fromHeight and toHeight",
		Example: fmt.Sprintf("relay teleport 1 1 teleport rinkeby \nrelay bsctest 1 1 bsctest rinkeby teleport 1"),
		Args:    cobra.RangeArgs(1, 6),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 && len(args) != 3 && len(args) != 7 {
				return errors.New("incorrect quantity")
			}
			ChainName = args[0]
			// when hash flag not nil ,relay by hash
			if cmd.Flag("hash").Value.String() != "" {
				err := manualRelay()
				if err != nil {
					return err
				}
			} else {
				arg1, err := strconv.ParseUint(args[1], 10, 64)
				if err != nil {
					return err
				}
				arg2, err := strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}

				FromHeight = arg1
				ToHeight = arg2

				if ToHeight < FromHeight || FromHeight < 1 {
					return errors.New("invalid height")
				}
				if len(args) == 6 {
					SrcChain = args[3]
					DestChain = args[4]
					arg5, err := strconv.ParseUint(args[5], 10, 64)
					if err != nil {
						return err
					}
					Sequence = arg5
					RelayChain = args[6]
				}
				err = manualRelay()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
)

func init() {
	configInitCmd.Flags().StringVarP(&config.Home, "home", "", "", "config path: .relayer")
	generateCmd.Flags().StringVarP(&config.Home, "home", "", "", "config path: .teleport-relayer")
	generateCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	genCmd.Flags().StringVarP(&generate.OutPut, "outPut", "o", "", "-o .")
	genCmd.Flags().StringVarP(&generate.PacketAddr, "packet", "", "", "-packet [packet address]")
	batchCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	startCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	startCmd.Flags().StringVarP(&config.Home, "home", "", "", "default ~/.teleport-relayer")
	startCmd.Flags().StringVarP(&config.Port, "port", "", "", "default 8080")
	manualRelayCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	manualRelayCmd.Flags().StringVarP(&Hash, "hash", "", "", "hash <hash>")

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(configInitCmd)
	rootCmd.AddCommand(batchCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(manualRelayCmd)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
