package main

import (
	"fmt"
	"os"

	clienthelpers "cosmossdk.io/client/v2/helpers"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/AliErcanOzgokce/babylon-relayer/app"
	"github.com/AliErcanOzgokce/babylon-relayer/cmd/babylon-relayerd/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, clienthelpers.EnvPrefix, app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
