package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cosmos/go-bip39"
	"github.com/reapchain/reapchain-core/privval"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	cfg "github.com/reapchain/reapchain-core/config"
	"github.com/reapchain/reapchain-core/libs/cli"
	tmos "github.com/reapchain/reapchain-core/libs/os"
	tmrand "github.com/reapchain/reapchain-core/libs/rand"
	"github.com/reapchain/reapchain-core/types"

	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/client/flags"
	"github.com/reapchain/cosmos-sdk/client/input"
	"github.com/reapchain/cosmos-sdk/server"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/module"
	"github.com/reapchain/cosmos-sdk/x/genutil"
	genutilcli "github.com/reapchain/cosmos-sdk/x/genutil/client/cli"
)

type printInfo struct {
	Moniker    string          `json:"moniker" yaml:"moniker"`
	ChainID    string          `json:"chain_id" yaml:"chain_id"`
	NodeID     string          `json:"node_id" yaml:"node_id"`
	GenTxsDir  string          `json:"gentxs_dir" yaml:"gentxs_dir"`
	AppMessage json.RawMessage `json:"app_message" yaml:"app_message"`
}

func newPrintInfo(moniker, chainID, nodeID, genTxsDir string, appMessage json.RawMessage) printInfo {
	return printInfo{
		Moniker:    moniker,
		ChainID:    chainID,
		NodeID:     nodeID,
		GenTxsDir:  genTxsDir,
		AppMessage: appMessage,
	}
}

func displayInfo(info printInfo) error {
	out, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(os.Stderr, "%s\n", string(sdk.MustSortJSON(out))); err != nil {
		return err
	}

	return nil
}

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func InitCmd(mbm module.BasicManager, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [moniker]",
		Short: "Initialize private validator, p2p, genesis, and application configuration files",
		Long:  `Initialize validators's and node's configuration files.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config
			config.SetRoot(clientCtx.HomeDir)

			// Set peers in and out to an 8:1 ratio to prevent choking
			config.P2P.MaxNumInboundPeers = 240
			config.P2P.MaxNumOutboundPeers = 30

			// Set default seeds
			seeds := []string{
				"40f4fac63da8b1ce8f850b0fa0f79b2699d2ce72@seed.reapchain.jerrychong.com:26656",                 // jerrychong
				"e3e11fca4ecf4035a751f3fea90e3a821e274487@bd-reapchain-mainnet-seed-node-01.bdnodes.net:26656", // blockdaemon
				"fc86e7e75c5d2e4699535e1b1bec98ae55b16826@bd-reapchain-mainnet-seed-node-02.bdnodes.net:26656", // blockdaemon
			}
			config.P2P.Seeds = strings.Join(seeds, ",")

			config.Mempool.Size = 10000
			config.StateSync.TrustPeriod = 112 * time.Hour

			config.SetRoot(clientCtx.HomeDir)

			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("evmos_9000-%v", tmrand.Str(6))
			}

			// Get bip39 mnemonic
			var mnemonic string
			recover, _ := cmd.Flags().GetBool(genutilcli.FlagRecover)
			if recover {
				inBuf := bufio.NewReader(cmd.InOrStdin())
				value, err := input.GetString("Enter your bip39 mnemonic", inBuf)
				if err != nil {
					return err
				}

				mnemonic = value
				if !bip39.IsMnemonicValid(mnemonic) {
					return errors.New("invalid mnemonic")
				}
			}

			nodeID, _, err := genutil.InitializeNodeValidatorFilesFromMnemonic(config, mnemonic)
			privValKeyFile := config.PrivValidatorKeyFile()
			privValStateFile := config.PrivValidatorStateFile()
			privValidator := privval.LoadOrGenFilePV(privValKeyFile, privValStateFile)
			corePubKey, _ := privValidator.GetPubKey()
			if err != nil {
				return err
			}

			config.Moniker = args[0]

			genFile := config.GenesisFile()
			overwrite, _ := cmd.Flags().GetBool(genutilcli.FlagOverwrite)

			if !overwrite && tmos.FileExists(genFile) {
				return fmt.Errorf("genesis.json file already exists: %v", genFile)
			}

			appState, err := json.MarshalIndent(mbm.DefaultGenesis(cdc), "", " ")
			if err != nil {
				return errors.Wrap(err, "Failed to marshall default genesis state")
			}

			genDoc := &types.GenesisDoc{}
			if _, err := os.Stat(genFile); err != nil {
				if !os.IsNotExist(err) {
					return err
				}
			} else {
				genDoc, err = types.GenesisDocFromFile(genFile)
				if err != nil {
					return errors.Wrap(err, "Failed to read genesis doc from file")
				}
			}

			genDoc.ChainID = chainID
			genDoc.Validators = nil
			genDoc.AppState = appState

			genDoc.Validators = []types.GenesisValidator{{
				Name:    args[0],
				Address: corePubKey.Address(),
				PubKey:  corePubKey,
				Power:   10,
				Type:    "standing",
			}}

			genDoc.StandingMembers = []types.GenesisMember{{
				Address: corePubKey.Address(),
				PubKey:  corePubKey,
				Name:    args[0],
				Power:   100,
			}}

			qrnValue := tmrand.Uint64()
			qrn := types.NewQrn(1, corePubKey, qrnValue)
			qrn.Timestamp = genDoc.GenesisTime

			err = privValidator.SignQrn(qrn)
			if err != nil {
				fmt.Println("Can't sign qrn", "err", err)
			}

			if qrn.VerifySign() == false {
				fmt.Println("Is invalid sign of qrn")
			}

			genDoc.Qrns = []types.Qrn{*qrn}

			if err := genDoc.SaveAs(genFile); err != nil {
				return err
			}

			genDoc.SteeringMemberCandidates = []types.GenesisMember{}

			genDoc.Vrfs = []types.Vrf{}

			genDoc.ConsensusRound = types.NewConsensusRound(1, 4, 4, 4)

			if err := genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return errors.Wrap(err, "Failed to export gensis file")
			}

			toPrint := newPrintInfo(config.Moniker, chainID, nodeID, "", appState)

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
			return displayInfo(toPrint)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().BoolP(genutilcli.FlagOverwrite, "o", false, "overwrite the genesis.json file")
	cmd.Flags().Bool(genutilcli.FlagRecover, false, "provide seed phrase to recover existing key instead of creating")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")

	return cmd
}
