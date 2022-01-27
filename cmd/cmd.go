package cmd

import (
	"github.com/teleport-network/teleport-relayer/version"
	"os"

	"github.com/spf13/cobra"

	"github.com/teleport-network/teleport-relayer/app/config"
	"github.com/teleport-network/teleport-relayer/cmd/generate"
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
	versionCmd = version.NewVersionCommand()
)

func init() {
	configInitCmd.Flags().StringVarP(&config.Home, "home", "", "", "config path: .relayer")
	generateCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	generateCmd.Flags().StringVarP(&config.Home, "home", "", "", "config path: .teleport-relayer")
	batchCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	startCmd.Flags().StringVarP(&config.LocalConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	startCmd.Flags().StringVarP(&config.Home, "home", "", "", "config path: /opt/local.toml")
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(configInitCmd)
	rootCmd.AddCommand(batchCmd)
	rootCmd.AddCommand(versionCmd)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
