package consensusplugin

import "github.com/cometbft/cometbft/p2p"

type ReactorI interface {
	GetReactors() map[string]p2p.Reactor
}

var ReactorPluginSingleton ReactorI
