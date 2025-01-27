package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/peggyjv/sommelier/x/allocation/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper, gs types.GenesisState) {
	k.setParams(ctx, gs.Params)
	// Set the vote period at initialization
	k.SetCommitPeriodStart(ctx, ctx.BlockHeight())

	for _, cellar := range gs.Cellars {
		k.SetCellar(ctx, *cellar)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	var cellars []*types.Cellar
	k.IterateCellars(ctx, func(cellar types.Cellar) (stop bool) {
		cellars = append(cellars, &cellar)
		return false
	})

	return types.GenesisState{
		Params:  k.GetParamSet(ctx),
		Cellars: cellars,
	}
}
