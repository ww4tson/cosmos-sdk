package v4

import (
	storetypes "cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
)

var ParamsKey = []byte{0x00}

// Migrate migrates the x/auth module state from the consensus version 3 to
// version 4. Specifically, it takes the parameters that are currently stored
// and managed by the x/params modules and stores them directly into the x/auth
// module state.
func Migrate(ctx sdk.Context, storeService storetypes.KVStoreService, legacySubspace exported.Subspace, cdc codec.BinaryCodec) error {
	// This migration has been removed from v0.48.0+. If you are migrating from v0.45/v0.46 to v0.48, please use the params.MigrateModuleFromParams function and migrate all core modules manually.

	// Example params.MigrateModuleFromParams(
	// 	ctx,
	// 	runtime.NewKVStoreService(keys[authtypes.StoreKey]),
	// 	app.GetSubspace(authtypes.ModuleName),
	// 	app.appCodec,
	// 	authtypes.Params{},
	// 	authtypes.ParamsKey,
	// )

	return nil
}
