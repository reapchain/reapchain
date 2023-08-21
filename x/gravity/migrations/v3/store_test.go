package v3

import (
	"testing"

	"github.com/reapchain/cosmos-sdk/codec"
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	"github.com/reapchain/cosmos-sdk/std"
	"github.com/reapchain/cosmos-sdk/testutil"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/module"
	"github.com/reapchain/cosmos-sdk/x/auth"
	"github.com/reapchain/cosmos-sdk/x/auth/vesting"
	"github.com/reapchain/cosmos-sdk/x/bank"
	"github.com/reapchain/cosmos-sdk/x/capability"
	"github.com/reapchain/cosmos-sdk/x/crisis"
	"github.com/reapchain/cosmos-sdk/x/distribution"
	distrclient "github.com/reapchain/cosmos-sdk/x/distribution/client"
	"github.com/reapchain/cosmos-sdk/x/evidence"
	"github.com/reapchain/cosmos-sdk/x/genutil"
	"github.com/reapchain/cosmos-sdk/x/gov"
	"github.com/reapchain/cosmos-sdk/x/mint"
	"github.com/reapchain/cosmos-sdk/x/params"
	paramsclient "github.com/reapchain/cosmos-sdk/x/params/client"
	"github.com/reapchain/cosmos-sdk/x/slashing"
	"github.com/reapchain/cosmos-sdk/x/staking"
	"github.com/reapchain/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/reapchain/cosmos-sdk/x/upgrade/client"
	"github.com/stretchr/testify/require"

	v2 "github.com/reapchain/reapchain/v8/x/gravity/migrations/v2"
	"github.com/reapchain/reapchain/v8/x/gravity/types"
)

func TestMigrateAttestation(t *testing.T) {
	// create old prefixes KV store
	gravityKey := sdk.NewKVStoreKey("gravity")
	ctx := testutil.DefaultContext(gravityKey, sdk.NewTransientStoreKey("transient-test"))
	store := ctx.KVStore(gravityKey)
	marshaler := MakeTestMarshaler()

	nonce := uint64(1)

	msg := types.MsgBatchSendToEthClaim{
		EventNonce:     nonce,
		EthBlockHeight: 1,
		BatchNonce:     nonce,
		TokenContract:  "0x00000000000000000001",
		Orchestrator:   "0x00000000000000000004",
	}
	msgAny, _ := codectypes.NewAnyWithValue(&msg)

	_, err := msg.ClaimHash()
	require.NoError(t, err)

	dummyAttestation := &types.Attestation{
		Observed: false,
		Height:   uint64(1),
		Claim:    msgAny,
	}
	oldClaimHash, err := v2.MsgBatchSendToEthClaimHash(msg)
	require.NoError(t, err)
	newClaimHash, err := msg.ClaimHash()
	require.NoError(t, err)
	attestationOldKey := v2.GetAttestationKey(nonce, oldClaimHash)

	store.Set(attestationOldKey, marshaler.MustMarshal(dummyAttestation))

	// Run migrations
	err = MigrateStore(ctx, gravityKey, marshaler)
	require.NoError(t, err)

	oldKeyEntry := store.Get(attestationOldKey)
	newKeyEntry := store.Get(types.GetAttestationKey(nonce, newClaimHash))
	// Check migration results:
	require.Empty(t, oldKeyEntry)
	require.NotEqual(t, oldKeyEntry, newKeyEntry)
	require.NotEqual(t, newKeyEntry, []byte(""))
	require.NotEmpty(t, newKeyEntry)
}

// Need to duplicate these because of cyclical imports
// ModuleBasics is a mock module basic manager for testing
var (
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		vesting.AppModuleBasic{},
	)
)

// MakeTestMarshaler creates a proto codec for use in testing
func MakeTestMarshaler() codec.Codec {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	ModuleBasics.RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)
	return codec.NewProtoCodec(interfaceRegistry)
}
