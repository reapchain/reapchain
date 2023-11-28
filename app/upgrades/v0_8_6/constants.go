package v0_8_6

import sdk "github.com/reapchain/cosmos-sdk/types"

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v0.8.6"
	// MainnetUpgradeHeight defines the Reapchain mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 305000
)

var (
	// MainnetMinGasPrices defines 20B aevmos (or atevmos) as the minimum gas price value on the fee market module.
	// See https://commonwealth.im/evmos/discussion/5073-global-min-gas-price-value-for-cosmos-sdk-and-evm-transaction-choosing-a-value for reference
	MainnetMinGasPrices = sdk.NewDec(20_000_000_000)
	// MainnetMinGasMultiplier defines the min gas multiplier value on the fee market module.
	// 50% of the leftover gas will be refunded
	MainnetMinGasMultiplier = sdk.NewDecWithPrec(5, 1)
)
