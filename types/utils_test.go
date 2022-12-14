package types

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/reapchain/cosmos-sdk/crypto/keys/ed25519"
	"github.com/reapchain/cosmos-sdk/crypto/keys/multisig"
	"github.com/reapchain/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/reapchain/cosmos-sdk/crypto/types"
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/reapchain/ethermint/crypto/ethsecp256k1"
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("reapchain", "reappub")
}

func TestIsSupportedKeys(t *testing.T) {
	testCases := []struct {
		name        string
		pk          cryptotypes.PubKey
		isSupported bool
	}{
		{
			"nil key",
			nil,
			false,
		},
		{
			"ethsecp256k1 key",
			&ethsecp256k1.PubKey{},
			true,
		},
		{
			"ed25519 key",
			&ed25519.PubKey{},
			true,
		},
		{
			"multisig key - no pubkeys",
			&multisig.LegacyAminoPubKey{},
			false,
		},
		{
			"multisig key - valid pubkeys",
			multisig.NewLegacyAminoPubKey(2, []cryptotypes.PubKey{&ed25519.PubKey{}, &ed25519.PubKey{}, &ed25519.PubKey{}}),
			true,
		},
		{
			"multisig key - nested multisig",
			multisig.NewLegacyAminoPubKey(2, []cryptotypes.PubKey{&ed25519.PubKey{}, &ed25519.PubKey{}, &multisig.LegacyAminoPubKey{}}),
			false,
		},
		{
			"multisig key - invalid pubkey",
			multisig.NewLegacyAminoPubKey(2, []cryptotypes.PubKey{&ed25519.PubKey{}, &ed25519.PubKey{}, &secp256k1.PubKey{}}),
			false,
		},
		{
			"cosmos secp256k1",
			&secp256k1.PubKey{},
			false,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, tc.isSupported, IsSupportedKey(tc.pk), tc.name)
	}
}

func TestGetReapchainAddressFromBech32(t *testing.T) {
	testCases := []struct {
		name       string
		address    string
		expAddress string
		expError   bool
	}{
		{
			"blank bech32 address",
			" ",
			"",
			true,
		},
		{
			"invalid bech32 address",
			"reapchain",
			"",
			true,
		},
		{
			"invalid address bytes",
			"reap123456",
			"",
			true,
		},
		{
			"reapchain address",
			"reap10hz0qunf49j9rvn3twrtdu0cm8hnav47vawaxm",
			"reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
			false,
		},
	}

	for _, tc := range testCases {
		addr, err := GetReapchainAddressFromBech32(tc.address)
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
			require.Equal(t, tc.expAddress, addr.String(), tc.name)
		}
	}
}
