package main_test

import (
	"fmt"
	"testing"

	"github.com/reapchain/cosmos-sdk/client/flags"
	svrcmd "github.com/reapchain/cosmos-sdk/server/cmd"
	"github.com/reapchain/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/tharsis/evmos/v4/app"
	evmosd "github.com/tharsis/evmos/v4/cmd/evmosd"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := evmosd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",       // Test the init cmd
		"evmos-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, "evmos_9000-1"),
	})

	err := svrcmd.Execute(rootCmd, app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := evmosd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"mykey",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, app.DefaultNodeHome)
	require.Error(t, err)
}
