package baseapp

import "context"

type CallbackFn = func(keeper interface{}, ctx context.Context) error

var endBlockerCallback = map[string]CallbackFn{}

// register callback for endblocker, runs at the end of any modules endblock
// only applies for cosmos-sdk native modules
// NOTE: This function is not thread-safe, call on newApp only
func RegisterEndBlockerCallback(moduleName string, cb CallbackFn) {
	endBlockerCallback[moduleName] = cb
}

func GetCallback(moduleName string) CallbackFn {
	return endBlockerCallback[moduleName]
}
