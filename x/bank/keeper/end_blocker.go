package keeper

import (
	"fmt"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k BaseKeeper) EndBlocker(ctx sdk.Context) {
	fmt.Println("bank end blocker runs")
	store := ctx.TransientStore(k.tKey)
	pref := prefix.NewStore(store, types.BalancesPrefix)
	itr := pref.Iterator(nil, nil)
	for itr.Valid() {
		k := itr.Key()
		v := itr.Value()
		fmt.Println("k:", k, "v:", v)
		itr.Next()
	}
}
