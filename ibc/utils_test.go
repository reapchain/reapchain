package ibc

import (
	"testing"

	"github.com/reapchain/reapchain/v4/x/claims/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/reapchain/cosmos-sdk/types"

	transfertypes "github.com/reapchain/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/reapchain/ibc-go/v3/modules/core/04-channel/types"
	ibctesting "github.com/reapchain/ibc-go/v3/testing"
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("reapchain", "reappub")
}

func TestGetTransferSenderRecipient(t *testing.T) {
	testCases := []struct {
		name         string
		packet       channeltypes.Packet
		expSender    string
		expRecipient string
		expError     bool
	}{
		{
			"empty packet",
			channeltypes.Packet{},
			"", "",
			true,
		},
		{
			"invalid packet data",
			channeltypes.Packet{
				Data: ibctesting.MockFailPacketData,
			},
			"", "",
			true,
		},
		{
			"empty FungibleTokenPacketData",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{},
				),
			},
			"", "",
			true,
		},
		{
			"invalid sender",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos1",
						Receiver: "reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
						Amount:   "123456",
					},
				),
			},
			"", "",
			true,
		},
		{
			"invalid recipient",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos198wd74dtl5zmr7p2y60h5syqg0j025u72ja9fn",
						Receiver: "reap1123123",
						Amount:   "123456",
					},
				),
			},
			"", "",
			true,
		},
		{
			"valid - cosmos sender, reapchain recipient",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos198wd74dtl5zmr7p2y60h5syqg0j025u72ja9fn",
						Receiver: "reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
						Amount:   "123456",
					},
				),
			},
			"reap10hz0qunf49j9rvn3twrtdu0cm8hnav47vawaxm",
			"reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
			false,
		},
		{
			"valid - reapchain sender, cosmos recipient",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
						Receiver: "cosmos198wd74dtl5zmr7p2y60h5syqg0j025u72ja9fn",
						Amount:   "123456",
					},
				),
			},
			"reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
			"reap10hz0qunf49j9rvn3twrtdu0cm8hnav47vawaxm",
			false,
		},
		{
			"valid - osmosis sender, reapchain recipient",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "osmo1qql8ag4cluz6r4dz28p3w00dnc9w8ueuhnecd2",
						Receiver: "reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
						Amount:   "123456",
					},
				),
			},
			"reap10hz0qunf49j9rvn3twrtdu0cm8hnav47vawaxm",
			"reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
			false,
		},
	}

	for _, tc := range testCases {
		sender, recipient, _, _, err := GetTransferSenderRecipient(tc.packet)
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
			require.Equal(t, tc.expSender, sender.String())
			require.Equal(t, tc.expRecipient, recipient.String())
		}
	}
}

func TestGetTransferAmount(t *testing.T) {
	testCases := []struct {
		name      string
		packet    channeltypes.Packet
		expAmount string
		expError  bool
	}{
		{
			"empty packet",
			channeltypes.Packet{},
			"",
			true,
		},
		{
			"invalid packet data",
			channeltypes.Packet{
				Data: ibctesting.MockFailPacketData,
			},
			"",
			true,
		},
		{
			"invalid amount - empty",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos198wd74dtl5zmr7p2y60h5syqg0j025u72ja9fn",
						Receiver: "reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
						Amount:   "",
					},
				),
			},
			"",
			true,
		},
		{
			"invalid amount - non-int",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos198wd74dtl5zmr7p2y60h5syqg0j025u72ja9fn",
						Receiver: "reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
						Amount:   "test",
					},
				),
			},
			"test",
			true,
		},
		{
			"valid",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos198wd74dtl5zmr7p2y60h5syqg0j025u72ja9fn",
						Receiver: "reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
						Amount:   "10000",
					},
				),
			},
			"10000",
			false,
		},
		{
			"valid - IBCTriggerAmt",
			channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos198wd74dtl5zmr7p2y60h5syqg0j025u72ja9fn",
						Receiver: "reap1ktkajgqv25830rgve5qwhwfs6m4hjruaml33ff",
						Amount:   types.IBCTriggerAmt,
					},
				),
			},
			types.IBCTriggerAmt,
			false,
		},
	}

	for _, tc := range testCases {
		amt, err := GetTransferAmount(tc.packet)
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
			require.Equal(t, tc.expAmount, amt)
		}
	}
}
