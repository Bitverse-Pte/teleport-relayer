package cmd

import (
	"github.com/teleport-network/teleport-relayer/app"
)

func online() {
	a := app.NewApp()
	a.Start()

}

func ethClientSync() {
	a := app.NewApp()
	a.EthClientSync()
}
