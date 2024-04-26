package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k BaseKeeper) EndBlocker(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.TransientStore(k.tKey)
	itr := store.Iterator(nil, nil)
	defer itr.Close()

	denomToUpdates := map[string][]*types.AccountBalance{}
	keyCdc, valueCdc := k.Balances.KeyCodec(), k.Balances.ValueCodec()
	for itr.Valid() {
		key := itr.Key()
		value := itr.Value()
		_, p, err := keyCdc.Decode(key)
		if err != nil {
			panic(err)
		}

		v, err := valueCdc.Decode(value)
		if err != nil {
			panic(err)
		}

		denom := p.K2()
		denomToUpdates[denom] = append(denomToUpdates[denom], &types.AccountBalance{
			Account: p.K1(),
			Balance: v,
		})
	}

	// events doesn't break consensus so we don't need to sort the map
	updateEvent := &types.BalanceUpdate{}
	for denom, updates := range denomToUpdates {
		updateEvent.Updates = append(updateEvent.Updates, &types.DenomBalanceUpdate{
			Denom:    denom,
			Balances: updates,
		})
	}
	sdkCtx.EventManager().EmitTypedEvent(updateEvent)
	return nil
}
