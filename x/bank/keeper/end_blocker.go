package keeper

import (
	"context"
)

func (k BaseKeeper) EndBlocker(ctx context.Context) error {
	if k.endBlockerCb != nil {
		return k.endBlockerCb(&k, ctx)
	}

	return nil
}
