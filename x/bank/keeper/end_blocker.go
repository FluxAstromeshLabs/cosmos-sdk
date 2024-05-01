package keeper

import (
	"context"
	"fmt"
	"slices"

	"golang.org/x/exp/maps"
)

func (k BaseKeeper) EndBlocker(ctx context.Context) error {
	cbMap := k.GetEndBlockerCallbackMap()
	names := maps.Keys(cbMap)
	slices.Sort(names)
	for _, name := range names {
		cb := cbMap[name]
		if cb != nil {
			err := cb(&k, ctx)
			if err != nil {
				return fmt.Errorf("callback '%s' err: %w", name, err)
			}
		}
	}

	return nil
}
