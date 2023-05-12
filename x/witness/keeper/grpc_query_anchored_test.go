package keeper_test

import (
	"testing"

	sdk "github.com/reapchain/cosmos-sdk/types"
	sdkerrors "github.com/reapchain/cosmos-sdk/types/errors"
	"github.com/reapchain/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "reapchain/testutil/keeper"
	"reapchain/testutil/nullify"
	"github.com/reapchain/reapchain/v4/x/witness/types"
)

func TestAnchoredQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.WitnessKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAnchored(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAnchoredRequest
		response *types.QueryGetAnchoredResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetAnchoredRequest{Id: msgs[0].Id},
			response: &types.QueryGetAnchoredResponse{Anchored: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetAnchoredRequest{Id: msgs[1].Id},
			response: &types.QueryGetAnchoredResponse{Anchored: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetAnchoredRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Anchored(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAnchoredQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.WitnessKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAnchored(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAnchoredRequest {
		return &types.QueryAllAnchoredRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AnchoredAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Anchored), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Anchored),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AnchoredAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Anchored), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Anchored),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AnchoredAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Anchored),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AnchoredAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
