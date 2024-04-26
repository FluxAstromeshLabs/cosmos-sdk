package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k BaseKeeper) EndBlocker(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.TransientStore(k.tKey)
	itr := store.Iterator(nil, nil)
	defer itr.Close()

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

		fmt.Println("address:", p.K1().String(), "denom:", p.K2(), "=> balance:", v.String())
		itr.Next()
	}
	return nil
}
