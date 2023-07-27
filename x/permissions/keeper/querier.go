package keeper

import (
	"context"

	"github.com/reapchain/cosmos-sdk/codec"
	"github.com/reapchain/cosmos-sdk/store/prefix"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/types/query"
	"github.com/reapchain/reapchain/v8/x/permissions/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetWhitelistedValidatorList(goCtx context.Context) (list []types.WhitelistedValidator, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var allWhiteListedValidators []types.WhitelistedValidator
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.WhiteListedValidatorKey))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		validator := MustUnmarshalValidator(k.cdc, iterator.Value())
		allWhiteListedValidators = append(allWhiteListedValidators, validator)
	}

	return allWhiteListedValidators, nil
}

func MustUnmarshalValidator(cdc codec.BinaryCodec, value []byte) types.WhitelistedValidator {
	validator, err := UnmarshalValidator(cdc, value)
	if err != nil {
		panic(err)
	}

	return validator
}

func UnmarshalValidator(cdc codec.BinaryCodec, value []byte) (v types.WhitelistedValidator, err error) {
	err = cdc.Unmarshal(value, &v)
	return v, err
}

func (k Keeper) GetAllWhiteListedValidators(goCtx context.Context, req *types.QueryAllWhitelistedValidatorsRequest) (*types.QueryAllWhitelistedValidatorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var allWhiteListedValidators []types.WhitelistedValidator

	store := ctx.KVStore(k.storeKey)
	whiteListedValidatorStore := prefix.NewStore(store, types.KeyPrefix(types.WhiteListedValidatorKey))

	pageRes, err := query.Paginate(whiteListedValidatorStore, req.Pagination, func(key []byte, value []byte) error {
		var whiteListedValidator types.WhitelistedValidator
		if err := k.cdc.Unmarshal(value, &whiteListedValidator); err != nil {
			return err
		}

		allWhiteListedValidators = append(allWhiteListedValidators, whiteListedValidator)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWhitelistedValidatorsResponse{WhitelistedValidators: allWhiteListedValidators, Pagination: pageRes}, nil
}

func (k Keeper) GetWhiteListedValidatorCount(goCtx context.Context) uint32 {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var counter uint32
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.WhiteListedValidatorKey))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		counter++
	}
	return counter
}

func (k Keeper) FindValidator(ctx sdk.Context, addr sdk.ValAddress) (found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteListedValidatorKey))
	value := store.Get(types.GetValidatorKey(addr))

	if value == nil {
		return false
	}

	return true
}

func (k Keeper) GetValidatorWithAddress(ctx sdk.Context, valAddr string) (validator types.WhitelistedValidator, err error) {

	validatorAddress, err := sdk.ValAddressFromBech32(valAddr)

	if err != nil {
		return validator, types.ErrInvalidValidatorAddress
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhiteListedValidatorKey))

	value := store.Get(types.GetValidatorKey(validatorAddress))
	if value == nil {
		return validator, types.ErrValidatorNotFound
	}
	var whiteListedValidator types.WhitelistedValidator

	if err := k.cdc.Unmarshal(value, &whiteListedValidator); err != nil {
		return validator, err
	}
	return whiteListedValidator, nil
}
