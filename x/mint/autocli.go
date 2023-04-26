package mint

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	mintv1beta1 "cosmossdk.io/api/cosmos/mint/v1beta1"
)

func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: mintv1beta1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					Use:   "params",
					Short: "query minting parameters",
				},
				{
					Use:   "inflation",
					Short: "Query the current minting inflation value",
				},
				{
					Use:   "annual-provisions",
					Short: "Query the current minting annual provisions value",
				},
			},
		},
	}
}