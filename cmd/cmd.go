package cmd

import (
	"github.com/teleport-network/teleport-relayer/version"
	"os"
	"strconv"

	"github.com/teleport-network/teleport-relayer/cmd/generate"

	"github.com/spf13/cobra"

	"github.com/teleport-network/teleport-relayer/app/config"
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
)

func init() {
	configInitCmd.Flags().StringVarP(&config.Home, "home", "", "", "config path: .relayer")
	generateCmd.Flags().StringVarP(&config.Home, "home", "", "", "config path: .teleport-relayer")
	generateCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	genCmd.Flags().StringVarP(&generate.OutPut, "outPut", "o", "", "-o .")
	genCmd.Flags().StringVarP(&generate.PacketAddr, "packet", "", "", "-packet [packet address]")
	batchCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	startCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	startCmd.Flags().StringVarP(&config.Home, "home", "", "", "config path: /opt/local.toml")
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(configInitCmd)
	rootCmd.AddCommand(batchCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(versionCmd)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
